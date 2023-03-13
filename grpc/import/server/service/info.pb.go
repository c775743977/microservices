// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: info.proto

package service

import (
	context "context"
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

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age     int32  `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{0}
}

func (x *InfoRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *InfoRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InfoRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age     int32  `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *InfoResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InfoResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_info_proto protoreflect.FileDescriptor

var file_info_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4f, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x50, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0x46, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x77, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_info_proto_rawDescOnce sync.Once
	file_info_proto_rawDescData = file_info_proto_rawDesc
)

func file_info_proto_rawDescGZIP() []byte {
	file_info_proto_rawDescOnce.Do(func() {
		file_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_info_proto_rawDescData)
	})
	return file_info_proto_rawDescData
}

var file_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_info_proto_goTypes = []interface{}{
	(*InfoRequest)(nil),  // 0: service.InfoRequest
	(*InfoResponse)(nil), // 1: service.InfoResponse
}
var file_info_proto_depIdxs = []int32{
	0, // 0: service.InfoService.ShowInfo:input_type -> service.InfoRequest
	1, // 1: service.InfoService.ShowInfo:output_type -> service.InfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_info_proto_init() }
func file_info_proto_init() {
	if File_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
			RawDescriptor: file_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_info_proto_goTypes,
		DependencyIndexes: file_info_proto_depIdxs,
		MessageInfos:      file_info_proto_msgTypes,
	}.Build()
	File_info_proto = out.File
	file_info_proto_rawDesc = nil
	file_info_proto_goTypes = nil
	file_info_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfoServiceClient interface {
	ShowInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type infoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoServiceClient(cc grpc.ClientConnInterface) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) ShowInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/service.InfoService/ShowInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
type InfoServiceServer interface {
	ShowInfo(context.Context, *InfoRequest) (*InfoResponse, error)
}

// UnimplementedInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInfoServiceServer struct {
}

func (*UnimplementedInfoServiceServer) ShowInfo(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowInfo not implemented")
}

func RegisterInfoServiceServer(s *grpc.Server, srv InfoServiceServer) {
	s.RegisterService(&_InfoService_serviceDesc, srv)
}

func _InfoService_ShowInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).ShowInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.InfoService/ShowInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).ShowInfo(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowInfo",
			Handler:    _InfoService_ShowInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info.proto",
}
