// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/config_dump.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	v21 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2"
	v2 "github.com/cilium/cilium/pkg/envoy/envoy/config/bootstrap/v2"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The :ref:`/config_dump <operations_admin_interface_config_dump>` admin endpoint uses this wrapper
// message to maintain and serve arbitrary configuration information from any component in Envoy.
type ConfigDump struct {
	// This list is serialized and dumped in its entirety at the
	// :ref:`/config_dump <operations_admin_interface_config_dump>` endpoint.
	//
	// The following configurations are currently supported and will be dumped in the order given
	// below:
	//
	// * *bootstrap*: :ref:`BootstrapConfigDump <envoy_api_msg_admin.v2alpha.BootstrapConfigDump>`
	// * *clusters*: :ref:`ClustersConfigDump <envoy_api_msg_admin.v2alpha.ClustersConfigDump>`
	// * *listeners*: :ref:`ListenersConfigDump <envoy_api_msg_admin.v2alpha.ListenersConfigDump>`
	// * *routes*:  :ref:`RoutesConfigDump <envoy_api_msg_admin.v2alpha.RoutesConfigDump>`
	Configs              []*any.Any `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConfigDump) Reset()         { *m = ConfigDump{} }
func (m *ConfigDump) String() string { return proto.CompactTextString(m) }
func (*ConfigDump) ProtoMessage()    {}
func (*ConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{0}
}

func (m *ConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigDump.Unmarshal(m, b)
}
func (m *ConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigDump.Marshal(b, m, deterministic)
}
func (m *ConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigDump.Merge(m, src)
}
func (m *ConfigDump) XXX_Size() int {
	return xxx_messageInfo_ConfigDump.Size(m)
}
func (m *ConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigDump proto.InternalMessageInfo

func (m *ConfigDump) GetConfigs() []*any.Any {
	if m != nil {
		return m.Configs
	}
	return nil
}

// This message describes the bootstrap configuration that Envoy was started with. This includes
// any CLI overrides that were merged. Bootstrap configuration information can be used to recreate
// the static portions of an Envoy configuration by reusing the output as the bootstrap
// configuration for another Envoy.
type BootstrapConfigDump struct {
	Bootstrap *v2.Bootstrap `protobuf:"bytes,1,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	// The timestamp when the BootstrapConfig was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BootstrapConfigDump) Reset()         { *m = BootstrapConfigDump{} }
func (m *BootstrapConfigDump) String() string { return proto.CompactTextString(m) }
func (*BootstrapConfigDump) ProtoMessage()    {}
func (*BootstrapConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{1}
}

func (m *BootstrapConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootstrapConfigDump.Unmarshal(m, b)
}
func (m *BootstrapConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootstrapConfigDump.Marshal(b, m, deterministic)
}
func (m *BootstrapConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapConfigDump.Merge(m, src)
}
func (m *BootstrapConfigDump) XXX_Size() int {
	return xxx_messageInfo_BootstrapConfigDump.Size(m)
}
func (m *BootstrapConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapConfigDump proto.InternalMessageInfo

func (m *BootstrapConfigDump) GetBootstrap() *v2.Bootstrap {
	if m != nil {
		return m.Bootstrap
	}
	return nil
}

func (m *BootstrapConfigDump) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Envoy's listener manager fills this message with all currently known listeners. Listener
// configuration information can be used to recreate an Envoy configuration by populating all
// listeners as static listeners or by returning them in a LDS response.
type ListenersConfigDump struct {
	// This is the :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` in the
	// last processed LDS discovery response. If there are only static bootstrap listeners, this field
	// will be "".
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The statically loaded listener configs.
	StaticListeners []*ListenersConfigDump_StaticListener `protobuf:"bytes,2,rep,name=static_listeners,json=staticListeners,proto3" json:"static_listeners,omitempty"`
	// The dynamically loaded active listeners. These are listeners that are available to service
	// data plane traffic.
	DynamicActiveListeners []*ListenersConfigDump_DynamicListener `protobuf:"bytes,3,rep,name=dynamic_active_listeners,json=dynamicActiveListeners,proto3" json:"dynamic_active_listeners,omitempty"`
	// The dynamically loaded warming listeners. These are listeners that are currently undergoing
	// warming in preparation to service data plane traffic. Note that if attempting to recreate an
	// Envoy configuration from a configuration dump, the warming listeners should generally be
	// discarded.
	DynamicWarmingListeners []*ListenersConfigDump_DynamicListener `protobuf:"bytes,4,rep,name=dynamic_warming_listeners,json=dynamicWarmingListeners,proto3" json:"dynamic_warming_listeners,omitempty"`
	// The dynamically loaded draining listeners. These are listeners that are currently undergoing
	// draining in preparation to stop servicing data plane traffic. Note that if attempting to
	// recreate an Envoy configuration from a configuration dump, the draining listeners should
	// generally be discarded.
	DynamicDrainingListeners []*ListenersConfigDump_DynamicListener `protobuf:"bytes,5,rep,name=dynamic_draining_listeners,json=dynamicDrainingListeners,proto3" json:"dynamic_draining_listeners,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                               `json:"-"`
	XXX_unrecognized         []byte                                 `json:"-"`
	XXX_sizecache            int32                                  `json:"-"`
}

func (m *ListenersConfigDump) Reset()         { *m = ListenersConfigDump{} }
func (m *ListenersConfigDump) String() string { return proto.CompactTextString(m) }
func (*ListenersConfigDump) ProtoMessage()    {}
func (*ListenersConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{2}
}

func (m *ListenersConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenersConfigDump.Unmarshal(m, b)
}
func (m *ListenersConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenersConfigDump.Marshal(b, m, deterministic)
}
func (m *ListenersConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenersConfigDump.Merge(m, src)
}
func (m *ListenersConfigDump) XXX_Size() int {
	return xxx_messageInfo_ListenersConfigDump.Size(m)
}
func (m *ListenersConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenersConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_ListenersConfigDump proto.InternalMessageInfo

func (m *ListenersConfigDump) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *ListenersConfigDump) GetStaticListeners() []*ListenersConfigDump_StaticListener {
	if m != nil {
		return m.StaticListeners
	}
	return nil
}

func (m *ListenersConfigDump) GetDynamicActiveListeners() []*ListenersConfigDump_DynamicListener {
	if m != nil {
		return m.DynamicActiveListeners
	}
	return nil
}

func (m *ListenersConfigDump) GetDynamicWarmingListeners() []*ListenersConfigDump_DynamicListener {
	if m != nil {
		return m.DynamicWarmingListeners
	}
	return nil
}

func (m *ListenersConfigDump) GetDynamicDrainingListeners() []*ListenersConfigDump_DynamicListener {
	if m != nil {
		return m.DynamicDrainingListeners
	}
	return nil
}

// Describes a statically loaded cluster.
type ListenersConfigDump_StaticListener struct {
	// The listener config.
	Listener *v21.Listener `protobuf:"bytes,1,opt,name=listener,proto3" json:"listener,omitempty"`
	// The timestamp when the Listener was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListenersConfigDump_StaticListener) Reset()         { *m = ListenersConfigDump_StaticListener{} }
func (m *ListenersConfigDump_StaticListener) String() string { return proto.CompactTextString(m) }
func (*ListenersConfigDump_StaticListener) ProtoMessage()    {}
func (*ListenersConfigDump_StaticListener) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{2, 0}
}

func (m *ListenersConfigDump_StaticListener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenersConfigDump_StaticListener.Unmarshal(m, b)
}
func (m *ListenersConfigDump_StaticListener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenersConfigDump_StaticListener.Marshal(b, m, deterministic)
}
func (m *ListenersConfigDump_StaticListener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenersConfigDump_StaticListener.Merge(m, src)
}
func (m *ListenersConfigDump_StaticListener) XXX_Size() int {
	return xxx_messageInfo_ListenersConfigDump_StaticListener.Size(m)
}
func (m *ListenersConfigDump_StaticListener) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenersConfigDump_StaticListener.DiscardUnknown(m)
}

var xxx_messageInfo_ListenersConfigDump_StaticListener proto.InternalMessageInfo

func (m *ListenersConfigDump_StaticListener) GetListener() *v21.Listener {
	if m != nil {
		return m.Listener
	}
	return nil
}

func (m *ListenersConfigDump_StaticListener) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Describes a dynamically loaded cluster via the LDS API.
type ListenersConfigDump_DynamicListener struct {
	// This is the per-resource version information. This version is currently taken from the
	// :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` field at the time
	// that the listener was loaded. In the future, discrete per-listener versions may be supported
	// by the API.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The listener config.
	Listener *v21.Listener `protobuf:"bytes,2,opt,name=listener,proto3" json:"listener,omitempty"`
	// The timestamp when the Listener was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListenersConfigDump_DynamicListener) Reset()         { *m = ListenersConfigDump_DynamicListener{} }
func (m *ListenersConfigDump_DynamicListener) String() string { return proto.CompactTextString(m) }
func (*ListenersConfigDump_DynamicListener) ProtoMessage()    {}
func (*ListenersConfigDump_DynamicListener) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{2, 1}
}

func (m *ListenersConfigDump_DynamicListener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenersConfigDump_DynamicListener.Unmarshal(m, b)
}
func (m *ListenersConfigDump_DynamicListener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenersConfigDump_DynamicListener.Marshal(b, m, deterministic)
}
func (m *ListenersConfigDump_DynamicListener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenersConfigDump_DynamicListener.Merge(m, src)
}
func (m *ListenersConfigDump_DynamicListener) XXX_Size() int {
	return xxx_messageInfo_ListenersConfigDump_DynamicListener.Size(m)
}
func (m *ListenersConfigDump_DynamicListener) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenersConfigDump_DynamicListener.DiscardUnknown(m)
}

var xxx_messageInfo_ListenersConfigDump_DynamicListener proto.InternalMessageInfo

func (m *ListenersConfigDump_DynamicListener) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *ListenersConfigDump_DynamicListener) GetListener() *v21.Listener {
	if m != nil {
		return m.Listener
	}
	return nil
}

func (m *ListenersConfigDump_DynamicListener) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Envoy's cluster manager fills this message with all currently known clusters. Cluster
// configuration information can be used to recreate an Envoy configuration by populating all
// clusters as static clusters or by returning them in a CDS response.
type ClustersConfigDump struct {
	// This is the :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` in the
	// last processed CDS discovery response. If there are only static bootstrap clusters, this field
	// will be "".
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The statically loaded cluster configs.
	StaticClusters []*ClustersConfigDump_StaticCluster `protobuf:"bytes,2,rep,name=static_clusters,json=staticClusters,proto3" json:"static_clusters,omitempty"`
	// The dynamically loaded active clusters. These are clusters that are available to service
	// data plane traffic.
	DynamicActiveClusters []*ClustersConfigDump_DynamicCluster `protobuf:"bytes,3,rep,name=dynamic_active_clusters,json=dynamicActiveClusters,proto3" json:"dynamic_active_clusters,omitempty"`
	// The dynamically loaded warming clusters. These are clusters that are currently undergoing
	// warming in preparation to service data plane traffic. Note that if attempting to recreate an
	// Envoy configuration from a configuration dump, the warming clusters should generally be
	// discarded.
	DynamicWarmingClusters []*ClustersConfigDump_DynamicCluster `protobuf:"bytes,4,rep,name=dynamic_warming_clusters,json=dynamicWarmingClusters,proto3" json:"dynamic_warming_clusters,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                             `json:"-"`
	XXX_unrecognized       []byte                               `json:"-"`
	XXX_sizecache          int32                                `json:"-"`
}

func (m *ClustersConfigDump) Reset()         { *m = ClustersConfigDump{} }
func (m *ClustersConfigDump) String() string { return proto.CompactTextString(m) }
func (*ClustersConfigDump) ProtoMessage()    {}
func (*ClustersConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{3}
}

func (m *ClustersConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClustersConfigDump.Unmarshal(m, b)
}
func (m *ClustersConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClustersConfigDump.Marshal(b, m, deterministic)
}
func (m *ClustersConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClustersConfigDump.Merge(m, src)
}
func (m *ClustersConfigDump) XXX_Size() int {
	return xxx_messageInfo_ClustersConfigDump.Size(m)
}
func (m *ClustersConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_ClustersConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_ClustersConfigDump proto.InternalMessageInfo

func (m *ClustersConfigDump) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *ClustersConfigDump) GetStaticClusters() []*ClustersConfigDump_StaticCluster {
	if m != nil {
		return m.StaticClusters
	}
	return nil
}

func (m *ClustersConfigDump) GetDynamicActiveClusters() []*ClustersConfigDump_DynamicCluster {
	if m != nil {
		return m.DynamicActiveClusters
	}
	return nil
}

func (m *ClustersConfigDump) GetDynamicWarmingClusters() []*ClustersConfigDump_DynamicCluster {
	if m != nil {
		return m.DynamicWarmingClusters
	}
	return nil
}

// Describes a statically loaded cluster.
type ClustersConfigDump_StaticCluster struct {
	// The cluster config.
	Cluster *v21.Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// The timestamp when the Cluster was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClustersConfigDump_StaticCluster) Reset()         { *m = ClustersConfigDump_StaticCluster{} }
func (m *ClustersConfigDump_StaticCluster) String() string { return proto.CompactTextString(m) }
func (*ClustersConfigDump_StaticCluster) ProtoMessage()    {}
func (*ClustersConfigDump_StaticCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{3, 0}
}

func (m *ClustersConfigDump_StaticCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClustersConfigDump_StaticCluster.Unmarshal(m, b)
}
func (m *ClustersConfigDump_StaticCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClustersConfigDump_StaticCluster.Marshal(b, m, deterministic)
}
func (m *ClustersConfigDump_StaticCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClustersConfigDump_StaticCluster.Merge(m, src)
}
func (m *ClustersConfigDump_StaticCluster) XXX_Size() int {
	return xxx_messageInfo_ClustersConfigDump_StaticCluster.Size(m)
}
func (m *ClustersConfigDump_StaticCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_ClustersConfigDump_StaticCluster.DiscardUnknown(m)
}

var xxx_messageInfo_ClustersConfigDump_StaticCluster proto.InternalMessageInfo

func (m *ClustersConfigDump_StaticCluster) GetCluster() *v21.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ClustersConfigDump_StaticCluster) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Describes a dynamically loaded cluster via the CDS API.
type ClustersConfigDump_DynamicCluster struct {
	// This is the per-resource version information. This version is currently taken from the
	// :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` field at the time
	// that the cluster was loaded. In the future, discrete per-cluster versions may be supported by
	// the API.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The cluster config.
	Cluster *v21.Cluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// The timestamp when the Cluster was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClustersConfigDump_DynamicCluster) Reset()         { *m = ClustersConfigDump_DynamicCluster{} }
func (m *ClustersConfigDump_DynamicCluster) String() string { return proto.CompactTextString(m) }
func (*ClustersConfigDump_DynamicCluster) ProtoMessage()    {}
func (*ClustersConfigDump_DynamicCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{3, 1}
}

func (m *ClustersConfigDump_DynamicCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClustersConfigDump_DynamicCluster.Unmarshal(m, b)
}
func (m *ClustersConfigDump_DynamicCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClustersConfigDump_DynamicCluster.Marshal(b, m, deterministic)
}
func (m *ClustersConfigDump_DynamicCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClustersConfigDump_DynamicCluster.Merge(m, src)
}
func (m *ClustersConfigDump_DynamicCluster) XXX_Size() int {
	return xxx_messageInfo_ClustersConfigDump_DynamicCluster.Size(m)
}
func (m *ClustersConfigDump_DynamicCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_ClustersConfigDump_DynamicCluster.DiscardUnknown(m)
}

var xxx_messageInfo_ClustersConfigDump_DynamicCluster proto.InternalMessageInfo

func (m *ClustersConfigDump_DynamicCluster) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *ClustersConfigDump_DynamicCluster) GetCluster() *v21.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ClustersConfigDump_DynamicCluster) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Envoy's RDS implementation fills this message with all currently loaded routes, as described by
// their RouteConfiguration objects. Static routes configured in the bootstrap configuration are
// separated from those configured dynamically via RDS. Route configuration information can be used
// to recreate an Envoy configuration by populating all routes as static routes or by returning them
// in RDS responses.
type RoutesConfigDump struct {
	// The statically loaded route configs.
	StaticRouteConfigs []*RoutesConfigDump_StaticRouteConfig `protobuf:"bytes,2,rep,name=static_route_configs,json=staticRouteConfigs,proto3" json:"static_route_configs,omitempty"`
	// The dynamically loaded route configs.
	DynamicRouteConfigs  []*RoutesConfigDump_DynamicRouteConfig `protobuf:"bytes,3,rep,name=dynamic_route_configs,json=dynamicRouteConfigs,proto3" json:"dynamic_route_configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *RoutesConfigDump) Reset()         { *m = RoutesConfigDump{} }
func (m *RoutesConfigDump) String() string { return proto.CompactTextString(m) }
func (*RoutesConfigDump) ProtoMessage()    {}
func (*RoutesConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{4}
}

func (m *RoutesConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesConfigDump.Unmarshal(m, b)
}
func (m *RoutesConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesConfigDump.Marshal(b, m, deterministic)
}
func (m *RoutesConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesConfigDump.Merge(m, src)
}
func (m *RoutesConfigDump) XXX_Size() int {
	return xxx_messageInfo_RoutesConfigDump.Size(m)
}
func (m *RoutesConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesConfigDump proto.InternalMessageInfo

func (m *RoutesConfigDump) GetStaticRouteConfigs() []*RoutesConfigDump_StaticRouteConfig {
	if m != nil {
		return m.StaticRouteConfigs
	}
	return nil
}

func (m *RoutesConfigDump) GetDynamicRouteConfigs() []*RoutesConfigDump_DynamicRouteConfig {
	if m != nil {
		return m.DynamicRouteConfigs
	}
	return nil
}

type RoutesConfigDump_StaticRouteConfig struct {
	// The route config.
	RouteConfig *v21.RouteConfiguration `protobuf:"bytes,1,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// The timestamp when the Route was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RoutesConfigDump_StaticRouteConfig) Reset()         { *m = RoutesConfigDump_StaticRouteConfig{} }
func (m *RoutesConfigDump_StaticRouteConfig) String() string { return proto.CompactTextString(m) }
func (*RoutesConfigDump_StaticRouteConfig) ProtoMessage()    {}
func (*RoutesConfigDump_StaticRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{4, 0}
}

func (m *RoutesConfigDump_StaticRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.Unmarshal(m, b)
}
func (m *RoutesConfigDump_StaticRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.Marshal(b, m, deterministic)
}
func (m *RoutesConfigDump_StaticRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.Merge(m, src)
}
func (m *RoutesConfigDump_StaticRouteConfig) XXX_Size() int {
	return xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.Size(m)
}
func (m *RoutesConfigDump_StaticRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesConfigDump_StaticRouteConfig proto.InternalMessageInfo

func (m *RoutesConfigDump_StaticRouteConfig) GetRouteConfig() *v21.RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *RoutesConfigDump_StaticRouteConfig) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type RoutesConfigDump_DynamicRouteConfig struct {
	// This is the per-resource version information. This version is currently taken from the
	// :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` field at the time that
	// the route configuration was loaded.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The route config.
	RouteConfig *v21.RouteConfiguration `protobuf:"bytes,2,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// The timestamp when the Route was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RoutesConfigDump_DynamicRouteConfig) Reset()         { *m = RoutesConfigDump_DynamicRouteConfig{} }
func (m *RoutesConfigDump_DynamicRouteConfig) String() string { return proto.CompactTextString(m) }
func (*RoutesConfigDump_DynamicRouteConfig) ProtoMessage()    {}
func (*RoutesConfigDump_DynamicRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e190b1a64d2aa, []int{4, 1}
}

func (m *RoutesConfigDump_DynamicRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.Unmarshal(m, b)
}
func (m *RoutesConfigDump_DynamicRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.Marshal(b, m, deterministic)
}
func (m *RoutesConfigDump_DynamicRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.Merge(m, src)
}
func (m *RoutesConfigDump_DynamicRouteConfig) XXX_Size() int {
	return xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.Size(m)
}
func (m *RoutesConfigDump_DynamicRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig proto.InternalMessageInfo

func (m *RoutesConfigDump_DynamicRouteConfig) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *RoutesConfigDump_DynamicRouteConfig) GetRouteConfig() *v21.RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *RoutesConfigDump_DynamicRouteConfig) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigDump)(nil), "envoy.admin.v2alpha.ConfigDump")
	proto.RegisterType((*BootstrapConfigDump)(nil), "envoy.admin.v2alpha.BootstrapConfigDump")
	proto.RegisterType((*ListenersConfigDump)(nil), "envoy.admin.v2alpha.ListenersConfigDump")
	proto.RegisterType((*ListenersConfigDump_StaticListener)(nil), "envoy.admin.v2alpha.ListenersConfigDump.StaticListener")
	proto.RegisterType((*ListenersConfigDump_DynamicListener)(nil), "envoy.admin.v2alpha.ListenersConfigDump.DynamicListener")
	proto.RegisterType((*ClustersConfigDump)(nil), "envoy.admin.v2alpha.ClustersConfigDump")
	proto.RegisterType((*ClustersConfigDump_StaticCluster)(nil), "envoy.admin.v2alpha.ClustersConfigDump.StaticCluster")
	proto.RegisterType((*ClustersConfigDump_DynamicCluster)(nil), "envoy.admin.v2alpha.ClustersConfigDump.DynamicCluster")
	proto.RegisterType((*RoutesConfigDump)(nil), "envoy.admin.v2alpha.RoutesConfigDump")
	proto.RegisterType((*RoutesConfigDump_StaticRouteConfig)(nil), "envoy.admin.v2alpha.RoutesConfigDump.StaticRouteConfig")
	proto.RegisterType((*RoutesConfigDump_DynamicRouteConfig)(nil), "envoy.admin.v2alpha.RoutesConfigDump.DynamicRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/config_dump.proto", fileDescriptor_bd4e190b1a64d2aa)
}

var fileDescriptor_bd4e190b1a64d2aa = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xc7, 0x71, 0x9a, 0xb6, 0xf4, 0x49, 0x69, 0x8b, 0xd3, 0xa6, 0xa9, 0x97, 0x96, 0x0a, 0xa4,
	0xb2, 0xd8, 0x52, 0x78, 0x5d, 0x18, 0xfa, 0x32, 0x80, 0xc4, 0x64, 0x40, 0x8c, 0xd6, 0x35, 0x76,
	0xd2, 0x93, 0x12, 0x9f, 0xe5, 0x3b, 0x1b, 0x82, 0x90, 0x18, 0xf8, 0x10, 0x6c, 0x2c, 0x0c, 0xec,
	0xac, 0xec, 0x88, 0x8d, 0x6f, 0xc0, 0x67, 0x41, 0xb9, 0x7b, 0xce, 0xf1, 0x4b, 0xa9, 0x92, 0x26,
	0x5b, 0xf2, 0xdc, 0x73, 0xff, 0xdf, 0xff, 0xee, 0xfe, 0x4f, 0x02, 0xf7, 0x82, 0x30, 0x65, 0x23,
	0x87, 0xf8, 0x43, 0x1a, 0x3a, 0x69, 0x87, 0x0c, 0xa2, 0x0b, 0xe2, 0x74, 0x59, 0xd8, 0xa3, 0x7d,
	0xcf, 0x4f, 0x86, 0x91, 0x1d, 0xc5, 0x4c, 0x30, 0xb3, 0x29, 0xdb, 0x6c, 0xd9, 0x66, 0x63, 0x9b,
	0xd5, 0xc2, 0xbd, 0x11, 0x75, 0xd2, 0x8e, 0xd3, 0xf5, 0xb9, 0x6a, 0x2e, 0xd5, 0x07, 0xff, 0xa9,
	0xc7, 0x59, 0xfd, 0xbe, 0xaa, 0x2b, 0xaa, 0x73, 0xce, 0x98, 0xe0, 0x22, 0x26, 0xd1, 0xb8, 0x29,
	0xfb, 0x82, 0xad, 0x7b, 0x7d, 0xc6, 0xfa, 0x83, 0xc0, 0x91, 0xdf, 0xce, 0x93, 0x9e, 0x43, 0xc2,
	0x11, 0x2e, 0xed, 0x97, 0x97, 0x04, 0x1d, 0x06, 0x5c, 0x10, 0x7d, 0x06, 0x6b, 0xbb, 0xcf, 0xfa,
	0x4c, 0x7e, 0x74, 0xc6, 0x9f, 0x54, 0xf5, 0xf0, 0x04, 0xe0, 0x54, 0x82, 0xcf, 0x92, 0x61, 0x64,
	0x3e, 0x84, 0x55, 0x65, 0x83, 0xb7, 0x8d, 0x83, 0xa5, 0xa3, 0x46, 0x67, 0xdb, 0x56, 0xb2, 0xb6,
	0x96, 0xb5, 0x8f, 0xc3, 0xd1, 0x49, 0xfd, 0xf7, 0xdf, 0xfd, 0x1b, 0xae, 0x6e, 0x3d, 0xfc, 0x6a,
	0x40, 0xf3, 0x44, 0x3b, 0xcd, 0xa9, 0x3d, 0x87, 0xb5, 0xec, 0x00, 0x6d, 0xe3, 0xc0, 0x38, 0x6a,
	0x74, 0xee, 0xda, 0xea, 0x26, 0xd5, 0x56, 0x7b, 0x72, 0xbe, 0xb4, 0x63, 0x67, 0x12, 0xa8, 0x3f,
	0xd9, 0x6c, 0x3e, 0x83, 0xf5, 0x01, 0xe1, 0xc2, 0x4b, 0x22, 0x9f, 0x88, 0xc0, 0x6f, 0xd7, 0xa4,
	0x98, 0x55, 0x31, 0xf7, 0x5a, 0x9f, 0xd9, 0x6d, 0x8c, 0xfb, 0xdf, 0xa8, 0xf6, 0xc3, 0x5f, 0x2b,
	0xd0, 0x7c, 0x49, 0xb9, 0x08, 0xc2, 0x20, 0xe6, 0x39, 0x83, 0x77, 0x60, 0x3d, 0x0d, 0x62, 0x4e,
	0x59, 0xe8, 0xd1, 0xb0, 0xc7, 0xa4, 0xc7, 0x35, 0xb7, 0x81, 0xb5, 0x17, 0x61, 0x8f, 0x99, 0x17,
	0xb0, 0xc5, 0x05, 0x11, 0xb4, 0xeb, 0x0d, 0xb4, 0x40, 0xbb, 0x26, 0xaf, 0xe6, 0x89, 0x7d, 0x49,
	0x28, 0xec, 0x4b, 0x30, 0xf6, 0x2b, 0x29, 0xa0, 0x57, 0xf0, 0x74, 0x9b, 0xbc, 0x50, 0xe5, 0xe6,
	0x7b, 0x68, 0xfb, 0xa3, 0x90, 0x0c, 0x69, 0xd7, 0x23, 0x5d, 0x41, 0xd3, 0x20, 0x47, 0x5c, 0x92,
	0xc4, 0xa7, 0x53, 0x13, 0xcf, 0x94, 0x50, 0x09, 0xd9, 0x42, 0xfd, 0x63, 0x29, 0x3f, 0x21, 0x7f,
	0x80, 0x3d, 0x4d, 0x7e, 0x47, 0xe2, 0x21, 0x0d, 0xfb, 0x39, 0x74, 0x7d, 0x21, 0xe8, 0x5d, 0x04,
	0xbc, 0x55, 0xfa, 0x13, 0xf6, 0x47, 0xb0, 0x34, 0xdb, 0x8f, 0x09, 0x0d, 0x8b, 0xf0, 0xe5, 0x85,
	0xc0, 0xf5, 0xbd, 0x9e, 0x21, 0x20, 0xdb, 0x69, 0x7d, 0x36, 0x60, 0xa3, 0xf8, 0x3a, 0x66, 0x07,
	0x6e, 0x6a, 0x3e, 0x66, 0xb6, 0xa5, 0xf1, 0x11, 0x1d, 0xc7, 0x54, 0x77, 0xba, 0x59, 0xdf, 0x9c,
	0xf1, 0xb4, 0xbe, 0x1b, 0xb0, 0x59, 0x72, 0x3e, 0x4d, 0x34, 0xf3, 0x4e, 0x6b, 0xd7, 0x74, 0xba,
	0x34, 0xdb, 0x20, 0xfd, 0x58, 0x06, 0xf3, 0x74, 0x90, 0x70, 0x31, 0xf3, 0x1c, 0xf9, 0x80, 0x81,
	0xf7, 0xba, 0xb8, 0x1f, 0xc7, 0xe8, 0xd1, 0xa5, 0x8f, 0x5b, 0x85, 0xe0, 0x14, 0xe1, 0x02, 0xbe,
	0xec, 0x06, 0xcf, 0x17, 0xb9, 0x29, 0x60, 0xb7, 0x34, 0x43, 0x19, 0x4d, 0x8d, 0xd0, 0xe3, 0x69,
	0x69, 0xf8, 0x1e, 0x45, 0xdc, 0x4e, 0x61, 0x80, 0x32, 0x6a, 0x3a, 0x99, 0x5c, 0x3d, 0x3f, 0x19,
	0xb6, 0xbe, 0x00, 0x6c, 0xab, 0x38, 0x3c, 0x7a, 0x9b, 0xf5, 0x09, 0x6e, 0x15, 0x2e, 0xc5, 0x74,
	0x60, 0x15, 0xc1, 0x18, 0xdd, 0x9d, 0x62, 0x20, 0xb0, 0xcf, 0xd5, 0x5d, 0xf3, 0x06, 0xf7, 0x9b,
	0x01, 0x1b, 0x45, 0xc7, 0xd3, 0x44, 0x21, 0xe7, 0xb2, 0x76, 0x2d, 0x97, 0x33, 0x86, 0xf6, 0x4f,
	0x1d, 0xb6, 0x5c, 0x96, 0x88, 0x20, 0x1f, 0x59, 0x06, 0xdb, 0x98, 0xc7, 0x78, 0xbc, 0xe4, 0xe9,
	0xbf, 0xbd, 0xab, 0x7e, 0xdb, 0xcb, 0x22, 0x18, 0x49, 0x59, 0x56, 0x55, 0x7c, 0x30, 0x93, 0x97,
	0x17, 0xb8, 0x19, 0x83, 0x4e, 0x4f, 0x89, 0x78, 0xd5, 0x6f, 0x7b, 0x85, 0x88, 0xb7, 0x5d, 0x45,
	0x36, 0xfd, 0xca, 0x0a, 0xb7, 0xbe, 0x18, 0x70, 0xbb, 0xe2, 0xd1, 0x3c, 0x85, 0xf5, 0xbc, 0x03,
	0x8c, 0xca, 0x41, 0xf1, 0x11, 0x72, 0x1b, 0x92, 0x98, 0x08, 0xca, 0x42, 0xb7, 0x11, 0xe7, 0x44,
	0xe6, 0x4c, 0xce, 0x4f, 0x03, 0xcc, 0xea, 0x59, 0xa6, 0x49, 0x4f, 0xd9, 0x7d, 0x6d, 0x11, 0xee,
	0x67, 0x4b, 0xd4, 0xf9, 0x8a, 0x6c, 0x78, 0xf0, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x5f, 0xdd,
	0xab, 0x3e, 0x0a, 0x00, 0x00,
}
