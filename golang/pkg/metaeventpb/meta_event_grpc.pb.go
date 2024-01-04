// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: meta_event.proto

package metaeventpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetaEventServiceClient is the client API for MetaEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaEventServiceClient interface {
	OpenShard(ctx context.Context, in *OpenShardRequest, opts ...grpc.CallOption) (*OpenShardResponse, error)
	CloseShard(ctx context.Context, in *CloseShardRequest, opts ...grpc.CallOption) (*CloseShardResponse, error)
	CreateTableOnShard(ctx context.Context, in *CreateTableOnShardRequest, opts ...grpc.CallOption) (*CreateTableOnShardResponse, error)
	DropTableOnShard(ctx context.Context, in *DropTableOnShardRequest, opts ...grpc.CallOption) (*DropTableOnShardResponse, error)
	OpenTableOnShard(ctx context.Context, in *OpenTableOnShardRequest, opts ...grpc.CallOption) (*OpenTableOnShardResponse, error)
	CloseTableOnShard(ctx context.Context, in *CloseTableOnShardRequest, opts ...grpc.CallOption) (*CloseTableOnShardResponse, error)
	SplitShard(ctx context.Context, in *SplitShardRequest, opts ...grpc.CallOption) (*SplitShardResponse, error)
	MergeShards(ctx context.Context, in *MergeShardsRequest, opts ...grpc.CallOption) (*MergeShardsResponse, error)
	ChangeShardRole(ctx context.Context, in *ChangeShardRoleRequest, opts ...grpc.CallOption) (*ChangeShardRoleResponse, error)
}

type metaEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaEventServiceClient(cc grpc.ClientConnInterface) MetaEventServiceClient {
	return &metaEventServiceClient{cc}
}

func (c *metaEventServiceClient) OpenShard(ctx context.Context, in *OpenShardRequest, opts ...grpc.CallOption) (*OpenShardResponse, error) {
	out := new(OpenShardResponse)
	err := c.cc.Invoke(ctx, "/meta_event.MetaEventService/OpenShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEventServiceClient) CloseShard(ctx context.Context, in *CloseShardRequest, opts ...grpc.CallOption) (*CloseShardResponse, error) {
	out := new(CloseShardResponse)
	err := c.cc.Invoke(ctx, "/meta_event.MetaEventService/CloseShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEventServiceClient) CreateTableOnShard(ctx context.Context, in *CreateTableOnShardRequest, opts ...grpc.CallOption) (*CreateTableOnShardResponse, error) {
	out := new(CreateTableOnShardResponse)
	err := c.cc.Invoke(ctx, "/meta_event.MetaEventService/CreateTableOnShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEventServiceClient) DropTableOnShard(ctx context.Context, in *DropTableOnShardRequest, opts ...grpc.CallOption) (*DropTableOnShardResponse, error) {
	out := new(DropTableOnShardResponse)
	err := c.cc.Invoke(ctx, "/meta_event.MetaEventService/DropTableOnShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEventServiceClient) OpenTableOnShard(ctx context.Context, in *OpenTableOnShardRequest, opts ...grpc.CallOption) (*OpenTableOnShardResponse, error) {
	out := new(OpenTableOnShardResponse)
	err := c.cc.Invoke(ctx, "/meta_event.MetaEventService/OpenTableOnShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEventServiceClient) CloseTableOnShard(ctx context.Context, in *CloseTableOnShardRequest, opts ...grpc.CallOption) (*CloseTableOnShardResponse, error) {
	out := new(CloseTableOnShardResponse)
	err := c.cc.Invoke(ctx, "/meta_event.MetaEventService/CloseTableOnShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEventServiceClient) SplitShard(ctx context.Context, in *SplitShardRequest, opts ...grpc.CallOption) (*SplitShardResponse, error) {
	out := new(SplitShardResponse)
	err := c.cc.Invoke(ctx, "/meta_event.MetaEventService/SplitShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEventServiceClient) MergeShards(ctx context.Context, in *MergeShardsRequest, opts ...grpc.CallOption) (*MergeShardsResponse, error) {
	out := new(MergeShardsResponse)
	err := c.cc.Invoke(ctx, "/meta_event.MetaEventService/MergeShards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEventServiceClient) ChangeShardRole(ctx context.Context, in *ChangeShardRoleRequest, opts ...grpc.CallOption) (*ChangeShardRoleResponse, error) {
	out := new(ChangeShardRoleResponse)
	err := c.cc.Invoke(ctx, "/meta_event.MetaEventService/ChangeShardRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaEventServiceServer is the server API for MetaEventService service.
// All implementations must embed UnimplementedMetaEventServiceServer
// for forward compatibility
type MetaEventServiceServer interface {
	OpenShard(context.Context, *OpenShardRequest) (*OpenShardResponse, error)
	CloseShard(context.Context, *CloseShardRequest) (*CloseShardResponse, error)
	CreateTableOnShard(context.Context, *CreateTableOnShardRequest) (*CreateTableOnShardResponse, error)
	DropTableOnShard(context.Context, *DropTableOnShardRequest) (*DropTableOnShardResponse, error)
	OpenTableOnShard(context.Context, *OpenTableOnShardRequest) (*OpenTableOnShardResponse, error)
	CloseTableOnShard(context.Context, *CloseTableOnShardRequest) (*CloseTableOnShardResponse, error)
	SplitShard(context.Context, *SplitShardRequest) (*SplitShardResponse, error)
	MergeShards(context.Context, *MergeShardsRequest) (*MergeShardsResponse, error)
	ChangeShardRole(context.Context, *ChangeShardRoleRequest) (*ChangeShardRoleResponse, error)
	mustEmbedUnimplementedMetaEventServiceServer()
}

// UnimplementedMetaEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetaEventServiceServer struct {
}

func (UnimplementedMetaEventServiceServer) OpenShard(context.Context, *OpenShardRequest) (*OpenShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenShard not implemented")
}
func (UnimplementedMetaEventServiceServer) CloseShard(context.Context, *CloseShardRequest) (*CloseShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseShard not implemented")
}
func (UnimplementedMetaEventServiceServer) CreateTableOnShard(context.Context, *CreateTableOnShardRequest) (*CreateTableOnShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTableOnShard not implemented")
}
func (UnimplementedMetaEventServiceServer) DropTableOnShard(context.Context, *DropTableOnShardRequest) (*DropTableOnShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropTableOnShard not implemented")
}
func (UnimplementedMetaEventServiceServer) OpenTableOnShard(context.Context, *OpenTableOnShardRequest) (*OpenTableOnShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenTableOnShard not implemented")
}
func (UnimplementedMetaEventServiceServer) CloseTableOnShard(context.Context, *CloseTableOnShardRequest) (*CloseTableOnShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTableOnShard not implemented")
}
func (UnimplementedMetaEventServiceServer) SplitShard(context.Context, *SplitShardRequest) (*SplitShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SplitShard not implemented")
}
func (UnimplementedMetaEventServiceServer) MergeShards(context.Context, *MergeShardsRequest) (*MergeShardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeShards not implemented")
}
func (UnimplementedMetaEventServiceServer) ChangeShardRole(context.Context, *ChangeShardRoleRequest) (*ChangeShardRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeShardRole not implemented")
}
func (UnimplementedMetaEventServiceServer) mustEmbedUnimplementedMetaEventServiceServer() {}

// UnsafeMetaEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetaEventServiceServer will
// result in compilation errors.
type UnsafeMetaEventServiceServer interface {
	mustEmbedUnimplementedMetaEventServiceServer()
}

func RegisterMetaEventServiceServer(s grpc.ServiceRegistrar, srv MetaEventServiceServer) {
	s.RegisterService(&MetaEventService_ServiceDesc, srv)
}

func _MetaEventService_OpenShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEventServiceServer).OpenShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_event.MetaEventService/OpenShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEventServiceServer).OpenShard(ctx, req.(*OpenShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEventService_CloseShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEventServiceServer).CloseShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_event.MetaEventService/CloseShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEventServiceServer).CloseShard(ctx, req.(*CloseShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEventService_CreateTableOnShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableOnShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEventServiceServer).CreateTableOnShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_event.MetaEventService/CreateTableOnShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEventServiceServer).CreateTableOnShard(ctx, req.(*CreateTableOnShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEventService_DropTableOnShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropTableOnShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEventServiceServer).DropTableOnShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_event.MetaEventService/DropTableOnShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEventServiceServer).DropTableOnShard(ctx, req.(*DropTableOnShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEventService_OpenTableOnShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenTableOnShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEventServiceServer).OpenTableOnShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_event.MetaEventService/OpenTableOnShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEventServiceServer).OpenTableOnShard(ctx, req.(*OpenTableOnShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEventService_CloseTableOnShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseTableOnShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEventServiceServer).CloseTableOnShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_event.MetaEventService/CloseTableOnShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEventServiceServer).CloseTableOnShard(ctx, req.(*CloseTableOnShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEventService_SplitShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SplitShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEventServiceServer).SplitShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_event.MetaEventService/SplitShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEventServiceServer).SplitShard(ctx, req.(*SplitShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEventService_MergeShards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeShardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEventServiceServer).MergeShards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_event.MetaEventService/MergeShards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEventServiceServer).MergeShards(ctx, req.(*MergeShardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEventService_ChangeShardRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeShardRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEventServiceServer).ChangeShardRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_event.MetaEventService/ChangeShardRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEventServiceServer).ChangeShardRole(ctx, req.(*ChangeShardRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetaEventService_ServiceDesc is the grpc.ServiceDesc for MetaEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetaEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meta_event.MetaEventService",
	HandlerType: (*MetaEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenShard",
			Handler:    _MetaEventService_OpenShard_Handler,
		},
		{
			MethodName: "CloseShard",
			Handler:    _MetaEventService_CloseShard_Handler,
		},
		{
			MethodName: "CreateTableOnShard",
			Handler:    _MetaEventService_CreateTableOnShard_Handler,
		},
		{
			MethodName: "DropTableOnShard",
			Handler:    _MetaEventService_DropTableOnShard_Handler,
		},
		{
			MethodName: "OpenTableOnShard",
			Handler:    _MetaEventService_OpenTableOnShard_Handler,
		},
		{
			MethodName: "CloseTableOnShard",
			Handler:    _MetaEventService_CloseTableOnShard_Handler,
		},
		{
			MethodName: "SplitShard",
			Handler:    _MetaEventService_SplitShard_Handler,
		},
		{
			MethodName: "MergeShards",
			Handler:    _MetaEventService_MergeShards_Handler,
		},
		{
			MethodName: "ChangeShardRole",
			Handler:    _MetaEventService_ChangeShardRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meta_event.proto",
}
