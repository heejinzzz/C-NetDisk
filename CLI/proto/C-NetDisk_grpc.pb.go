// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: proto/C-NetDisk.proto

package proto

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

// CNetDiskClient is the client API for CNetDisk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CNetDiskClient interface {
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	GetItemInfo(ctx context.Context, in *GetItemInfoRequest, opts ...grpc.CallOption) (*GetItemInfoResponse, error)
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (CNetDisk_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (CNetDisk_DownloadFileClient, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error)
	RenameItem(ctx context.Context, in *RenameItemRequest, opts ...grpc.CallOption) (*RenameItemResponse, error)
}

type cNetDiskClient struct {
	cc grpc.ClientConnInterface
}

func NewCNetDiskClient(cc grpc.ClientConnInterface) CNetDiskClient {
	return &cNetDiskClient{cc}
}

func (c *cNetDiskClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, "/CNetDisk/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cNetDiskClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/CNetDisk/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cNetDiskClient) GetItemInfo(ctx context.Context, in *GetItemInfoRequest, opts ...grpc.CallOption) (*GetItemInfoResponse, error) {
	out := new(GetItemInfoResponse)
	err := c.cc.Invoke(ctx, "/CNetDisk/GetItemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cNetDiskClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, "/CNetDisk/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cNetDiskClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (CNetDisk_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &CNetDisk_ServiceDesc.Streams[0], "/CNetDisk/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &cNetDiskUploadFileClient{stream}
	return x, nil
}

type CNetDisk_UploadFileClient interface {
	Send(*UploadFileRequest) error
	Recv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type cNetDiskUploadFileClient struct {
	grpc.ClientStream
}

func (x *cNetDiskUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cNetDiskUploadFileClient) Recv() (*UploadFileResponse, error) {
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cNetDiskClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (CNetDisk_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &CNetDisk_ServiceDesc.Streams[1], "/CNetDisk/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &cNetDiskDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CNetDisk_DownloadFileClient interface {
	Recv() (*DownloadFileResponse, error)
	grpc.ClientStream
}

type cNetDiskDownloadFileClient struct {
	grpc.ClientStream
}

func (x *cNetDiskDownloadFileClient) Recv() (*DownloadFileResponse, error) {
	m := new(DownloadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cNetDiskClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error) {
	out := new(DeleteItemResponse)
	err := c.cc.Invoke(ctx, "/CNetDisk/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cNetDiskClient) RenameItem(ctx context.Context, in *RenameItemRequest, opts ...grpc.CallOption) (*RenameItemResponse, error) {
	out := new(RenameItemResponse)
	err := c.cc.Invoke(ctx, "/CNetDisk/RenameItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CNetDiskServer is the server API for CNetDisk service.
// All implementations should embed UnimplementedCNetDiskServer
// for forward compatibility
type CNetDiskServer interface {
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	GetItemInfo(context.Context, *GetItemInfoRequest) (*GetItemInfoResponse, error)
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	UploadFile(CNetDisk_UploadFileServer) error
	DownloadFile(*DownloadFileRequest, CNetDisk_DownloadFileServer) error
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error)
	RenameItem(context.Context, *RenameItemRequest) (*RenameItemResponse, error)
}

// UnimplementedCNetDiskServer should be embedded to have forward compatible implementations.
type UnimplementedCNetDiskServer struct {
}

func (UnimplementedCNetDiskServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedCNetDiskServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedCNetDiskServer) GetItemInfo(context.Context, *GetItemInfoRequest) (*GetItemInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemInfo not implemented")
}
func (UnimplementedCNetDiskServer) CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedCNetDiskServer) UploadFile(CNetDisk_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedCNetDiskServer) DownloadFile(*DownloadFileRequest, CNetDisk_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedCNetDiskServer) DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedCNetDiskServer) RenameItem(context.Context, *RenameItemRequest) (*RenameItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameItem not implemented")
}

// UnsafeCNetDiskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CNetDiskServer will
// result in compilation errors.
type UnsafeCNetDiskServer interface {
	mustEmbedUnimplementedCNetDiskServer()
}

func RegisterCNetDiskServer(s grpc.ServiceRegistrar, srv CNetDiskServer) {
	s.RegisterService(&CNetDisk_ServiceDesc, srv)
}

func _CNetDisk_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNetDiskServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CNetDisk/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNetDiskServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CNetDisk_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNetDiskServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CNetDisk/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNetDiskServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CNetDisk_GetItemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNetDiskServer).GetItemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CNetDisk/GetItemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNetDiskServer).GetItemInfo(ctx, req.(*GetItemInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CNetDisk_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNetDiskServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CNetDisk/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNetDiskServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CNetDisk_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CNetDiskServer).UploadFile(&cNetDiskUploadFileServer{stream})
}

type CNetDisk_UploadFileServer interface {
	Send(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type cNetDiskUploadFileServer struct {
	grpc.ServerStream
}

func (x *cNetDiskUploadFileServer) Send(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cNetDiskUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CNetDisk_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CNetDiskServer).DownloadFile(m, &cNetDiskDownloadFileServer{stream})
}

type CNetDisk_DownloadFileServer interface {
	Send(*DownloadFileResponse) error
	grpc.ServerStream
}

type cNetDiskDownloadFileServer struct {
	grpc.ServerStream
}

func (x *cNetDiskDownloadFileServer) Send(m *DownloadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CNetDisk_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNetDiskServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CNetDisk/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNetDiskServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CNetDisk_RenameItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNetDiskServer).RenameItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CNetDisk/RenameItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNetDiskServer).RenameItem(ctx, req.(*RenameItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CNetDisk_ServiceDesc is the grpc.ServiceDesc for CNetDisk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CNetDisk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CNetDisk",
	HandlerType: (*CNetDiskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _CNetDisk_UserRegister_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _CNetDisk_UserLogin_Handler,
		},
		{
			MethodName: "GetItemInfo",
			Handler:    _CNetDisk_GetItemInfo_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _CNetDisk_CreateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _CNetDisk_DeleteItem_Handler,
		},
		{
			MethodName: "RenameItem",
			Handler:    _CNetDisk_RenameItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _CNetDisk_UploadFile_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _CNetDisk_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/C-NetDisk.proto",
}
