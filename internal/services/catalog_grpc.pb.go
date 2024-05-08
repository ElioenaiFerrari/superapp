// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: catalog.proto

package services

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

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogServiceClient interface {
	GetCatalog(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*Catalog, error)
	ListCatalogs(ctx context.Context, in *ListCatalogsRequest, opts ...grpc.CallOption) (*ListCatalogsResponse, error)
	GetCatalogQRCode(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*GetCatalogQRCodeResponse, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) GetCatalog(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*Catalog, error) {
	out := new(Catalog)
	err := c.cc.Invoke(ctx, "/services.CatalogService/GetCatalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListCatalogs(ctx context.Context, in *ListCatalogsRequest, opts ...grpc.CallOption) (*ListCatalogsResponse, error) {
	out := new(ListCatalogsResponse)
	err := c.cc.Invoke(ctx, "/services.CatalogService/ListCatalogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCatalogQRCode(ctx context.Context, in *GetCatalogRequest, opts ...grpc.CallOption) (*GetCatalogQRCodeResponse, error) {
	out := new(GetCatalogQRCodeResponse)
	err := c.cc.Invoke(ctx, "/services.CatalogService/GetCatalogQRCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
// All implementations should embed UnimplementedCatalogServiceServer
// for forward compatibility
type CatalogServiceServer interface {
	GetCatalog(context.Context, *GetCatalogRequest) (*Catalog, error)
	ListCatalogs(context.Context, *ListCatalogsRequest) (*ListCatalogsResponse, error)
	GetCatalogQRCode(context.Context, *GetCatalogRequest) (*GetCatalogQRCodeResponse, error)
}

// UnimplementedCatalogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (UnimplementedCatalogServiceServer) GetCatalog(context.Context, *GetCatalogRequest) (*Catalog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCatalog not implemented")
}
func (UnimplementedCatalogServiceServer) ListCatalogs(context.Context, *ListCatalogsRequest) (*ListCatalogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCatalogs not implemented")
}
func (UnimplementedCatalogServiceServer) GetCatalogQRCode(context.Context, *GetCatalogRequest) (*GetCatalogQRCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCatalogQRCode not implemented")
}

// UnsafeCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServiceServer will
// result in compilation errors.
type UnsafeCatalogServiceServer interface {
	mustEmbedUnimplementedCatalogServiceServer()
}

func RegisterCatalogServiceServer(s grpc.ServiceRegistrar, srv CatalogServiceServer) {
	s.RegisterService(&CatalogService_ServiceDesc, srv)
}

func _CatalogService_GetCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CatalogService/GetCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCatalog(ctx, req.(*GetCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListCatalogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatalogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListCatalogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CatalogService/ListCatalogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListCatalogs(ctx, req.(*ListCatalogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCatalogQRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCatalogQRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CatalogService/GetCatalogQRCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCatalogQRCode(ctx, req.(*GetCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogService_ServiceDesc is the grpc.ServiceDesc for CatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCatalog",
			Handler:    _CatalogService_GetCatalog_Handler,
		},
		{
			MethodName: "ListCatalogs",
			Handler:    _CatalogService_ListCatalogs_Handler,
		},
		{
			MethodName: "GetCatalogQRCode",
			Handler:    _CatalogService_GetCatalogQRCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog.proto",
}
