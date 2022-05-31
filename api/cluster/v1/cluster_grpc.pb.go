// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.4
// source: cluster/v1/cluster.proto

package v1

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

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterClient interface {
	// get cluster info
	GetCluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterReply, error)
	// list cluster metadata
	ListCluster(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListReply, error)
	// create cluster
	CreateCluster(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*ClusterCreateReply, error)
	// update cluster
	UpdateCluster(ctx context.Context, in *ClusterUpdateRequest, opts ...grpc.CallOption) (*ClusterUpdateReply, error)
	// delete cluster with clusterID
	DeleteCluster(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*ClusterDeleteReply, error)
	// managed external cluster
	ManagedCluster(ctx context.Context, in *ClusterManagedRequest, opts ...grpc.CallOption) (*ClusterManagedReply, error)
	// log cluster with clusterID
	LogCluster(ctx context.Context, in *ClusterLogRequest, opts ...grpc.CallOption) (*ClusterLogReply, error)
	// watch cluster with clusterID
	WatchCluster(ctx context.Context, in *ClusterWatchRequest, opts ...grpc.CallOption) (Cluster_WatchClusterClient, error)
	// compare difference between current version and last version of clusters
	CompareCluster(ctx context.Context, in *ClusterCompareRequest, opts ...grpc.CallOption) (*ClusterCompareReply, error)
}

type clusterClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterClient(cc grpc.ClientConnInterface) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) GetCluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterReply, error) {
	out := new(ClusterReply)
	err := c.cc.Invoke(ctx, "/cluster.v1.Cluster/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ListCluster(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListReply, error) {
	out := new(ClusterListReply)
	err := c.cc.Invoke(ctx, "/cluster.v1.Cluster/ListCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) CreateCluster(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*ClusterCreateReply, error) {
	out := new(ClusterCreateReply)
	err := c.cc.Invoke(ctx, "/cluster.v1.Cluster/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) UpdateCluster(ctx context.Context, in *ClusterUpdateRequest, opts ...grpc.CallOption) (*ClusterUpdateReply, error) {
	out := new(ClusterUpdateReply)
	err := c.cc.Invoke(ctx, "/cluster.v1.Cluster/UpdateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) DeleteCluster(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*ClusterDeleteReply, error) {
	out := new(ClusterDeleteReply)
	err := c.cc.Invoke(ctx, "/cluster.v1.Cluster/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ManagedCluster(ctx context.Context, in *ClusterManagedRequest, opts ...grpc.CallOption) (*ClusterManagedReply, error) {
	out := new(ClusterManagedReply)
	err := c.cc.Invoke(ctx, "/cluster.v1.Cluster/ManagedCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) LogCluster(ctx context.Context, in *ClusterLogRequest, opts ...grpc.CallOption) (*ClusterLogReply, error) {
	out := new(ClusterLogReply)
	err := c.cc.Invoke(ctx, "/cluster.v1.Cluster/LogCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) WatchCluster(ctx context.Context, in *ClusterWatchRequest, opts ...grpc.CallOption) (Cluster_WatchClusterClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cluster_ServiceDesc.Streams[0], "/cluster.v1.Cluster/WatchCluster", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterWatchClusterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cluster_WatchClusterClient interface {
	Recv() (*ClusterWatchReply, error)
	grpc.ClientStream
}

type clusterWatchClusterClient struct {
	grpc.ClientStream
}

func (x *clusterWatchClusterClient) Recv() (*ClusterWatchReply, error) {
	m := new(ClusterWatchReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterClient) CompareCluster(ctx context.Context, in *ClusterCompareRequest, opts ...grpc.CallOption) (*ClusterCompareReply, error) {
	out := new(ClusterCompareReply)
	err := c.cc.Invoke(ctx, "/cluster.v1.Cluster/CompareCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServer is the server API for Cluster service.
// All implementations must embed UnimplementedClusterServer
// for forward compatibility
type ClusterServer interface {
	// get cluster info
	GetCluster(context.Context, *ClusterRequest) (*ClusterReply, error)
	// list cluster metadata
	ListCluster(context.Context, *ClusterListRequest) (*ClusterListReply, error)
	// create cluster
	CreateCluster(context.Context, *ClusterCreateRequest) (*ClusterCreateReply, error)
	// update cluster
	UpdateCluster(context.Context, *ClusterUpdateRequest) (*ClusterUpdateReply, error)
	// delete cluster with clusterID
	DeleteCluster(context.Context, *ClusterDeleteRequest) (*ClusterDeleteReply, error)
	// managed external cluster
	ManagedCluster(context.Context, *ClusterManagedRequest) (*ClusterManagedReply, error)
	// log cluster with clusterID
	LogCluster(context.Context, *ClusterLogRequest) (*ClusterLogReply, error)
	// watch cluster with clusterID
	WatchCluster(*ClusterWatchRequest, Cluster_WatchClusterServer) error
	// compare difference between current version and last version of clusters
	CompareCluster(context.Context, *ClusterCompareRequest) (*ClusterCompareReply, error)
	mustEmbedUnimplementedClusterServer()
}

// UnimplementedClusterServer must be embedded to have forward compatible implementations.
type UnimplementedClusterServer struct {
}

func (UnimplementedClusterServer) GetCluster(context.Context, *ClusterRequest) (*ClusterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedClusterServer) ListCluster(context.Context, *ClusterListRequest) (*ClusterListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCluster not implemented")
}
func (UnimplementedClusterServer) CreateCluster(context.Context, *ClusterCreateRequest) (*ClusterCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedClusterServer) UpdateCluster(context.Context, *ClusterUpdateRequest) (*ClusterUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCluster not implemented")
}
func (UnimplementedClusterServer) DeleteCluster(context.Context, *ClusterDeleteRequest) (*ClusterDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (UnimplementedClusterServer) ManagedCluster(context.Context, *ClusterManagedRequest) (*ClusterManagedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManagedCluster not implemented")
}
func (UnimplementedClusterServer) LogCluster(context.Context, *ClusterLogRequest) (*ClusterLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogCluster not implemented")
}
func (UnimplementedClusterServer) WatchCluster(*ClusterWatchRequest, Cluster_WatchClusterServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchCluster not implemented")
}
func (UnimplementedClusterServer) CompareCluster(context.Context, *ClusterCompareRequest) (*ClusterCompareReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompareCluster not implemented")
}
func (UnimplementedClusterServer) mustEmbedUnimplementedClusterServer() {}

// UnsafeClusterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServer will
// result in compilation errors.
type UnsafeClusterServer interface {
	mustEmbedUnimplementedClusterServer()
}

func RegisterClusterServer(s grpc.ServiceRegistrar, srv ClusterServer) {
	s.RegisterService(&Cluster_ServiceDesc, srv)
}

func _Cluster_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.Cluster/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).GetCluster(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_ListCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ListCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.Cluster/ListCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ListCluster(ctx, req.(*ClusterListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.Cluster/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).CreateCluster(ctx, req.(*ClusterCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.Cluster/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).UpdateCluster(ctx, req.(*ClusterUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.Cluster/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).DeleteCluster(ctx, req.(*ClusterDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_ManagedCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterManagedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ManagedCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.Cluster/ManagedCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ManagedCluster(ctx, req.(*ClusterManagedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_LogCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).LogCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.Cluster/LogCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).LogCluster(ctx, req.(*ClusterLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_WatchCluster_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClusterWatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClusterServer).WatchCluster(m, &clusterWatchClusterServer{stream})
}

type Cluster_WatchClusterServer interface {
	Send(*ClusterWatchReply) error
	grpc.ServerStream
}

type clusterWatchClusterServer struct {
	grpc.ServerStream
}

func (x *clusterWatchClusterServer) Send(m *ClusterWatchReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Cluster_CompareCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterCompareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).CompareCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.v1.Cluster/CompareCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).CompareCluster(ctx, req.(*ClusterCompareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cluster_ServiceDesc is the grpc.ServiceDesc for Cluster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cluster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cluster.v1.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCluster",
			Handler:    _Cluster_GetCluster_Handler,
		},
		{
			MethodName: "ListCluster",
			Handler:    _Cluster_ListCluster_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _Cluster_CreateCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _Cluster_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _Cluster_DeleteCluster_Handler,
		},
		{
			MethodName: "ManagedCluster",
			Handler:    _Cluster_ManagedCluster_Handler,
		},
		{
			MethodName: "LogCluster",
			Handler:    _Cluster_LogCluster_Handler,
		},
		{
			MethodName: "CompareCluster",
			Handler:    _Cluster_CompareCluster_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchCluster",
			Handler:       _Cluster_WatchCluster_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cluster/v1/cluster.proto",
}
