// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/managed_placement_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [ManagedPlacementViewService.GetManagedPlacementView][google.ads.googleads.v3.services.ManagedPlacementViewService.GetManagedPlacementView].
type GetManagedPlacementViewRequest struct {
	// Required. The resource name of the Managed Placement View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetManagedPlacementViewRequest) Reset()         { *m = GetManagedPlacementViewRequest{} }
func (m *GetManagedPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetManagedPlacementViewRequest) ProtoMessage()    {}
func (*GetManagedPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ee6f7f223a331f, []int{0}
}

func (m *GetManagedPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetManagedPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetManagedPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManagedPlacementViewRequest.Merge(m, src)
}
func (m *GetManagedPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Size(m)
}
func (m *GetManagedPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManagedPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetManagedPlacementViewRequest proto.InternalMessageInfo

func (m *GetManagedPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetManagedPlacementViewRequest)(nil), "google.ads.googleads.v3.services.GetManagedPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/managed_placement_view_service.proto", fileDescriptor_d6ee6f7f223a331f)
}

var fileDescriptor_d6ee6f7f223a331f = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x25, 0x39, 0x10, 0x0c, 0xda, 0xa4, 0xb9, 0x23, 0x27, 0x1a, 0x8e, 0x2b, 0x8e, 0x2b, 0x66,
	0xc0, 0x14, 0x87, 0x23, 0x8a, 0xb3, 0x20, 0x2b, 0x82, 0xba, 0x9c, 0x90, 0x42, 0x02, 0x61, 0x2e,
	0xf9, 0x8c, 0x03, 0xc9, 0x4c, 0xcc, 0xcc, 0xe6, 0x0a, 0xb1, 0xb1, 0xf1, 0x07, 0xf8, 0x0f, 0xae,
	0xf4, 0xa7, 0x6c, 0x6b, 0x67, 0x65, 0xa1, 0x8d, 0xbf, 0xe2, 0xc8, 0x4e, 0x26, 0xbb, 0x0b, 0x9b,
	0xdd, 0xee, 0x91, 0xf7, 0xbe, 0xf7, 0xbe, 0xf9, 0x5e, 0xbc, 0x97, 0x85, 0x94, 0x45, 0x09, 0x98,
	0xe5, 0x0a, 0x1b, 0xd8, 0xa1, 0x36, 0xc2, 0x0a, 0x9a, 0x96, 0x67, 0xa0, 0x70, 0xc5, 0x04, 0x2b,
	0x20, 0x4f, 0xeb, 0x92, 0x65, 0x50, 0x81, 0xd0, 0x69, 0xcb, 0xe1, 0x3a, 0xed, 0x79, 0x54, 0x37,
	0x52, 0x4b, 0x3f, 0x34, 0xb3, 0x88, 0xe5, 0x0a, 0x0d, 0x36, 0xa8, 0x8d, 0x90, 0xb5, 0x09, 0x9e,
	0x8f, 0x05, 0x35, 0xa0, 0xe4, 0xbc, 0x19, 0x4f, 0x32, 0x09, 0xc1, 0x03, 0x3b, 0x5f, 0x73, 0xcc,
	0x84, 0x90, 0x9a, 0x69, 0x2e, 0x85, 0xea, 0xd9, 0xc3, 0x35, 0x36, 0x2b, 0x39, 0x08, 0xdd, 0x13,
	0x8f, 0xd6, 0x88, 0x8f, 0x1c, 0xca, 0x3c, 0xbd, 0x82, 0x4f, 0xac, 0xe5, 0xb2, 0x31, 0x82, 0x93,
	0xd7, 0xde, 0xc3, 0x29, 0xe8, 0x37, 0x26, 0x7a, 0x66, 0x93, 0x63, 0x0e, 0xd7, 0x97, 0xf0, 0x79,
	0x0e, 0x4a, 0xfb, 0x67, 0xde, 0x7d, 0xbb, 0x63, 0x2a, 0x58, 0x05, 0x47, 0x4e, 0xe8, 0x9c, 0xdd,
	0x9d, 0x1c, 0xfc, 0xa1, 0xee, 0xe5, 0x3d, 0xcb, 0xbc, 0x65, 0x15, 0x3c, 0xbe, 0x71, 0xbd, 0xe3,
	0x6d, 0x4e, 0xef, 0xcd, 0x11, 0xfc, 0x7f, 0x8e, 0x77, 0x38, 0x12, 0xe6, 0xbf, 0x40, 0xfb, 0x4e,
	0x88, 0x76, 0xef, 0x19, 0x5c, 0x8c, 0x3a, 0x0c, 0x27, 0x46, 0xdb, 0xe6, 0x4f, 0xde, 0xfd, 0xa6,
	0x9b, 0x2f, 0xfc, 0xf6, 0xeb, 0xef, 0x0f, 0xf7, 0x89, 0x7f, 0xd1, 0xd5, 0xf3, 0x65, 0x83, 0x79,
	0x96, 0xcd, 0x95, 0x96, 0x15, 0x34, 0x0a, 0x9f, 0xdb, 0xbe, 0x36, 0xcc, 0x14, 0x3e, 0xff, 0x1a,
	0x1c, 0x2f, 0xe8, 0xd1, 0x6a, 0x81, 0x1e, 0xd5, 0x5c, 0xa1, 0x4c, 0x56, 0x93, 0xef, 0xae, 0x77,
	0x9a, 0xc9, 0x6a, 0xef, 0x73, 0x27, 0xe1, 0x8e, 0x53, 0xce, 0xba, 0xee, 0x66, 0xce, 0x87, 0x57,
	0xbd, 0x4b, 0x21, 0x4b, 0x26, 0x0a, 0x24, 0x9b, 0x02, 0x17, 0x20, 0x96, 0xcd, 0xe2, 0x55, 0xee,
	0xf8, 0xdf, 0xfd, 0xd4, 0x82, 0x1b, 0xf7, 0x60, 0x4a, 0xe9, 0x4f, 0x37, 0x9c, 0x1a, 0x43, 0x9a,
	0x2b, 0x64, 0x60, 0x87, 0xe2, 0x08, 0xf5, 0xc1, 0x6a, 0x61, 0x25, 0x09, 0xcd, 0x55, 0x32, 0x48,
	0x92, 0x38, 0x4a, 0xac, 0xe4, 0xbf, 0x7b, 0x6a, 0xbe, 0x13, 0x42, 0x73, 0x45, 0xc8, 0x20, 0x22,
	0x24, 0x8e, 0x08, 0xb1, 0xb2, 0xab, 0x3b, 0xcb, 0x3d, 0xa3, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0xdd, 0xdf, 0xe2, 0x84, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ManagedPlacementViewServiceClient is the client API for ManagedPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagedPlacementViewServiceClient interface {
	// Returns the requested Managed Placement view in full detail.
	GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error)
}

type managedPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagedPlacementViewServiceClient(cc grpc.ClientConnInterface) ManagedPlacementViewServiceClient {
	return &managedPlacementViewServiceClient{cc}
}

func (c *managedPlacementViewServiceClient) GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error) {
	out := new(resources.ManagedPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.ManagedPlacementViewService/GetManagedPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagedPlacementViewServiceServer is the server API for ManagedPlacementViewService service.
type ManagedPlacementViewServiceServer interface {
	// Returns the requested Managed Placement view in full detail.
	GetManagedPlacementView(context.Context, *GetManagedPlacementViewRequest) (*resources.ManagedPlacementView, error)
}

// UnimplementedManagedPlacementViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedManagedPlacementViewServiceServer struct {
}

func (*UnimplementedManagedPlacementViewServiceServer) GetManagedPlacementView(ctx context.Context, req *GetManagedPlacementViewRequest) (*resources.ManagedPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagedPlacementView not implemented")
}

func RegisterManagedPlacementViewServiceServer(s *grpc.Server, srv ManagedPlacementViewServiceServer) {
	s.RegisterService(&_ManagedPlacementViewService_serviceDesc, srv)
}

func _ManagedPlacementViewService_GetManagedPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManagedPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.ManagedPlacementViewService/GetManagedPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, req.(*GetManagedPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagedPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.ManagedPlacementViewService",
	HandlerType: (*ManagedPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManagedPlacementView",
			Handler:    _ManagedPlacementViewService_GetManagedPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/managed_placement_view_service.proto",
}