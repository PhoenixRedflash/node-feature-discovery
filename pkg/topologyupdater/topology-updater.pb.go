//
//Copyright 2021 The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.17.3
// source: pkg/topologyupdater/topology-updater.proto

package topologyupdater

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NodeTopologyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NfdVersion       string           `protobuf:"bytes,1,opt,name=nfd_version,json=nfdVersion,proto3" json:"nfd_version,omitempty"`
	NodeName         string           `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	TopologyPolicies []string         `protobuf:"bytes,3,rep,name=topology_policies,json=topologyPolicies,proto3" json:"topology_policies,omitempty"`
	Zones            []*v1alpha1.Zone `protobuf:"bytes,4,rep,name=zones,proto3" json:"zones,omitempty"`
}

func (x *NodeTopologyRequest) Reset() {
	*x = NodeTopologyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_topologyupdater_topology_updater_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeTopologyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeTopologyRequest) ProtoMessage() {}

func (x *NodeTopologyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_topologyupdater_topology_updater_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeTopologyRequest.ProtoReflect.Descriptor instead.
func (*NodeTopologyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_topologyupdater_topology_updater_proto_rawDescGZIP(), []int{0}
}

func (x *NodeTopologyRequest) GetNfdVersion() string {
	if x != nil {
		return x.NfdVersion
	}
	return ""
}

func (x *NodeTopologyRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *NodeTopologyRequest) GetTopologyPolicies() []string {
	if x != nil {
		return x.TopologyPolicies
	}
	return nil
}

func (x *NodeTopologyRequest) GetZones() []*v1alpha1.Zone {
	if x != nil {
		return x.Zones
	}
	return nil
}

type NodeTopologyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NodeTopologyResponse) Reset() {
	*x = NodeTopologyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_topologyupdater_topology_updater_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeTopologyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeTopologyResponse) ProtoMessage() {}

func (x *NodeTopologyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_topologyupdater_topology_updater_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeTopologyResponse.ProtoReflect.Descriptor instead.
func (*NodeTopologyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_topologyupdater_topology_updater_proto_rawDescGZIP(), []int{1}
}

var File_pkg_topologyupdater_topology_updater_proto protoreflect.FileDescriptor

var file_pkg_topologyupdater_topology_updater_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2d, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x1a, 0x66, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x38, 0x73, 0x74, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x61, 0x77, 0x61, 0x72, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x77,
	0x67, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x6e, 0x66, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6e, 0x66, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x74,
	0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x7a, 0x6f, 0x6e, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0x16,
	0x0a, 0x14, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x71, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x61, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x24, 0x2e, 0x74,
	0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x73, 0x69, 0x67,
	0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_topologyupdater_topology_updater_proto_rawDescOnce sync.Once
	file_pkg_topologyupdater_topology_updater_proto_rawDescData = file_pkg_topologyupdater_topology_updater_proto_rawDesc
)

func file_pkg_topologyupdater_topology_updater_proto_rawDescGZIP() []byte {
	file_pkg_topologyupdater_topology_updater_proto_rawDescOnce.Do(func() {
		file_pkg_topologyupdater_topology_updater_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_topologyupdater_topology_updater_proto_rawDescData)
	})
	return file_pkg_topologyupdater_topology_updater_proto_rawDescData
}

var file_pkg_topologyupdater_topology_updater_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_topologyupdater_topology_updater_proto_goTypes = []interface{}{
	(*NodeTopologyRequest)(nil),  // 0: topologyupdater.NodeTopologyRequest
	(*NodeTopologyResponse)(nil), // 1: topologyupdater.NodeTopologyResponse
	(*v1alpha1.Zone)(nil),        // 2: v1alpha1.Zone
}
var file_pkg_topologyupdater_topology_updater_proto_depIdxs = []int32{
	2, // 0: topologyupdater.NodeTopologyRequest.zones:type_name -> v1alpha1.Zone
	0, // 1: topologyupdater.NodeTopology.UpdateNodeTopology:input_type -> topologyupdater.NodeTopologyRequest
	1, // 2: topologyupdater.NodeTopology.UpdateNodeTopology:output_type -> topologyupdater.NodeTopologyResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_topologyupdater_topology_updater_proto_init() }
func file_pkg_topologyupdater_topology_updater_proto_init() {
	if File_pkg_topologyupdater_topology_updater_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_topologyupdater_topology_updater_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeTopologyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_topologyupdater_topology_updater_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeTopologyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_topologyupdater_topology_updater_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_topologyupdater_topology_updater_proto_goTypes,
		DependencyIndexes: file_pkg_topologyupdater_topology_updater_proto_depIdxs,
		MessageInfos:      file_pkg_topologyupdater_topology_updater_proto_msgTypes,
	}.Build()
	File_pkg_topologyupdater_topology_updater_proto = out.File
	file_pkg_topologyupdater_topology_updater_proto_rawDesc = nil
	file_pkg_topologyupdater_topology_updater_proto_goTypes = nil
	file_pkg_topologyupdater_topology_updater_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeTopologyClient is the client API for NodeTopology service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeTopologyClient interface {
	UpdateNodeTopology(ctx context.Context, in *NodeTopologyRequest, opts ...grpc.CallOption) (*NodeTopologyResponse, error)
}

type nodeTopologyClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeTopologyClient(cc grpc.ClientConnInterface) NodeTopologyClient {
	return &nodeTopologyClient{cc}
}

func (c *nodeTopologyClient) UpdateNodeTopology(ctx context.Context, in *NodeTopologyRequest, opts ...grpc.CallOption) (*NodeTopologyResponse, error) {
	out := new(NodeTopologyResponse)
	err := c.cc.Invoke(ctx, "/topologyupdater.NodeTopology/UpdateNodeTopology", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeTopologyServer is the server API for NodeTopology service.
type NodeTopologyServer interface {
	UpdateNodeTopology(context.Context, *NodeTopologyRequest) (*NodeTopologyResponse, error)
}

// UnimplementedNodeTopologyServer can be embedded to have forward compatible implementations.
type UnimplementedNodeTopologyServer struct {
}

func (*UnimplementedNodeTopologyServer) UpdateNodeTopology(context.Context, *NodeTopologyRequest) (*NodeTopologyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeTopology not implemented")
}

func RegisterNodeTopologyServer(s *grpc.Server, srv NodeTopologyServer) {
	s.RegisterService(&_NodeTopology_serviceDesc, srv)
}

func _NodeTopology_UpdateNodeTopology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeTopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTopologyServer).UpdateNodeTopology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topologyupdater.NodeTopology/UpdateNodeTopology",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTopologyServer).UpdateNodeTopology(ctx, req.(*NodeTopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeTopology_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topologyupdater.NodeTopology",
	HandlerType: (*NodeTopologyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateNodeTopology",
			Handler:    _NodeTopology_UpdateNodeTopology_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/topologyupdater/topology-updater.proto",
}
