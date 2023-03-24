// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: test.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string       `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Info     *InfoRequest `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetInfo() *InfoRequest {
	if x != nil {
		return x.Info
	}
	return nil
}

type Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *Login) Reset() {
	*x = Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Login) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x60, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0x19, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0x6d,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b,
	0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x12, 0x5a,
	0x10, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*User)(nil),        // 0: service.User
	(*Login)(nil),       // 1: service.Login
	(*InfoRequest)(nil), // 2: service.InfoRequest
}
var file_test_proto_depIdxs = []int32{
	2, // 0: service.User.info:type_name -> service.InfoRequest
	0, // 1: service.LoginService.CheckLogin:input_type -> service.User
	2, // 2: service.LoginService.CheckAge:input_type -> service.InfoRequest
	1, // 3: service.LoginService.CheckLogin:output_type -> service.Login
	1, // 4: service.LoginService.CheckAge:output_type -> service.Login
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	file_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginServiceClient interface {
	CheckLogin(ctx context.Context, in *User, opts ...grpc.CallOption) (*Login, error)
	CheckAge(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*Login, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) CheckLogin(ctx context.Context, in *User, opts ...grpc.CallOption) (*Login, error) {
	out := new(Login)
	err := c.cc.Invoke(ctx, "/service.LoginService/CheckLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) CheckAge(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*Login, error) {
	out := new(Login)
	err := c.cc.Invoke(ctx, "/service.LoginService/CheckAge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
type LoginServiceServer interface {
	CheckLogin(context.Context, *User) (*Login, error)
	CheckAge(context.Context, *InfoRequest) (*Login, error)
}

// UnimplementedLoginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (*UnimplementedLoginServiceServer) CheckLogin(context.Context, *User) (*Login, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckLogin not implemented")
}
func (*UnimplementedLoginServiceServer) CheckAge(context.Context, *InfoRequest) (*Login, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAge not implemented")
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_CheckLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).CheckLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.LoginService/CheckLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).CheckLogin(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_CheckAge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).CheckAge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.LoginService/CheckAge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).CheckAge(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckLogin",
			Handler:    _LoginService_CheckLogin_Handler,
		},
		{
			MethodName: "CheckAge",
			Handler:    _LoginService_CheckAge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}