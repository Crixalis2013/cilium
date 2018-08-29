/*
 *  Copyright (C) 2018 Authors of Cilium
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program; if not, write to the Free Software
 *  Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 */
#include <node_config.h>
#include <bpf/api.h>

#include <stdint.h>
#include <stdio.h>

#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <sys/socket.h>

#define SOCKMAP 1

#include "../lib/utils.h"
#include "../lib/common.h"
#include "../lib/maps.h"
#include "../lib/lb.h"
#include "../lib/eps.h"
#include "../lib/events.h"
#include "../lib/policy.h"

#include "sockops_config.h"
#include "bpf_sockops.h"

static inline void bpf_sock_ops_ipv4(struct bpf_sock_ops *skops)
{
	struct endpoint_info *exists;
	struct lb4_key lb4_key = {};
	struct sock_key key = {};
	struct lb4_service *svc;
	int err, verdict;
	__u32 dstID = 0;

	sk_extract4_key(skops, &key);
	sk_lb4_key(&lb4_key, &key);

	bpf_printk("skops ipv4: (IP)  sip4 %u dip4 %u family %u\n", key.sip4, key.dip4, key.family);
	bpf_printk("skops ipv4: (TCP) dport %u sport %u\n", key.dport, key.sport);

	/* If endpoint a service use L4/L3 stack for now. These can be
	 * pulled in as needed.
	 */
	svc = __lb4_lookup_service(&lb4_key);
	if (svc) {
		bpf_printk("ipv4 endpoint: services required use L4/L3 stack\n");
		return;
	}

	/* Policy lookup required to learn proxy port */
	if (1) {
		struct remote_endpoint_info *info;

		info = lookup_ip4_remote_endpoint(key.dip4);
		if (info != NULL && info->sec_label) {
			dstID = info->sec_label;
		} else if ((key.dip4 & IPV4_CLUSTER_MASK) == IPV4_CLUSTER_RANGE) {
			dstID = CLUSTER_ID;
		} else {
			dstID = WORLD_ID;
		}
	}

	verdict = policy_sk_egress(dstID, key.sip4, key.dport);
	if (redirect_to_proxy(verdict)) {
		err = sock_hash_update(skops, &SOCK_OPS_MAP, &key, BPF_ANY);
		if (err) {
			bpf_printk("ipv4 endpoint: map proxy update failed\n");
		} else {
			bpf_printk("ipv4 endpoint: map proxy update\n");
		}
		return;
	}

	/* Lookup IPv4 address, this will return a match if:
	 * - The destination IP address belongs to the local endpoint manage
	 *   by Cilium.
	 * - The destination IP address is an IP address associated with the
	 *   host itself.
	 * Then because these are local IPs that have passed LB/Policy/NAT
	 * blocks redirect directly to socket.
	 */
	exists = __lookup_ip4_endpoint(key.dip4);
	if (!exists)
		return;

	bpf_printk("skops ipv4: local endpoint %d\n", exists);
	err = sock_hash_update(skops, &SOCK_OPS_MAP, &key, BPF_ANY);
	if (err)
		bpf_printk("ipv4 endpoint: redirect local hash update failed\n");
	else
		bpf_printk("ipv4 endpoint: redirect local hash update\n");
}

static inline void bpf_sock_ops_ipv6(struct bpf_sock_ops *skops)
{

}

__section("sockops")
int bpf_sockmap(struct bpf_sock_ops *skops)
{
	__u32 family, op;

	if (!skops)
		return 0;

	family = skops->family;
	op = skops->op;

	switch (op) {
	case BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB:
	case BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB:
		if (family == AF_INET6)
			bpf_sock_ops_ipv6(skops);
		else if (family == AF_INET)
			bpf_sock_ops_ipv4(skops);
		break;
	default:
		break;
	}

	return 0;
}

BPF_LICENSE("GPL");
int _version __section("version") = 1;
