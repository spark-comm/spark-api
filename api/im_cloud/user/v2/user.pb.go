// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: im_cloud/user/v2/user.proto

package v2

import (
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

type GetUserLoginStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetUserLoginStatusReq) Reset() {
	*x = GetUserLoginStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_cloud_user_v2_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLoginStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLoginStatusReq) ProtoMessage() {}

func (x *GetUserLoginStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_cloud_user_v2_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLoginStatusReq.ProtoReflect.Descriptor instead.
func (*GetUserLoginStatusReq) Descriptor() ([]byte, []int) {
	return file_im_cloud_user_v2_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserLoginStatusReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserLoginStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetUserLoginStatusReply) Reset() {
	*x = GetUserLoginStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_cloud_user_v2_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLoginStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLoginStatusReply) ProtoMessage() {}

func (x *GetUserLoginStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_im_cloud_user_v2_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLoginStatusReply.ProtoReflect.Descriptor instead.
func (*GetUserLoginStatusReply) Descriptor() ([]byte, []int) {
	return file_im_cloud_user_v2_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserLoginStatusReply) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetUserLoginStatusReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_im_cloud_user_v2_user_proto protoreflect.FileDescriptor

var file_im_cloud_user_v2_user_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x49, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0x78, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x70, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x51, 0x0a, 0x14, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_cloud_user_v2_user_proto_rawDescOnce sync.Once
	file_im_cloud_user_v2_user_proto_rawDescData = file_im_cloud_user_v2_user_proto_rawDesc
)

func file_im_cloud_user_v2_user_proto_rawDescGZIP() []byte {
	file_im_cloud_user_v2_user_proto_rawDescOnce.Do(func() {
		file_im_cloud_user_v2_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_cloud_user_v2_user_proto_rawDescData)
	})
	return file_im_cloud_user_v2_user_proto_rawDescData
}

var file_im_cloud_user_v2_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_im_cloud_user_v2_user_proto_goTypes = []interface{}{
	(*GetUserLoginStatusReq)(nil),   // 0: api.im_cloud.user.v2.GetUserLoginStatusReq
	(*GetUserLoginStatusReply)(nil), // 1: api.im_cloud.user.v2.GetUserLoginStatusReply
}
var file_im_cloud_user_v2_user_proto_depIdxs = []int32{
	0, // 0: api.im_cloud.user.v2.User.GetUserLoginStatus:input_type -> api.im_cloud.user.v2.GetUserLoginStatusReq
	1, // 1: api.im_cloud.user.v2.User.GetUserLoginStatus:output_type -> api.im_cloud.user.v2.GetUserLoginStatusReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_im_cloud_user_v2_user_proto_init() }
func file_im_cloud_user_v2_user_proto_init() {
	if File_im_cloud_user_v2_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_im_cloud_user_v2_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLoginStatusReq); i {
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
		file_im_cloud_user_v2_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLoginStatusReply); i {
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
			RawDescriptor: file_im_cloud_user_v2_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_cloud_user_v2_user_proto_goTypes,
		DependencyIndexes: file_im_cloud_user_v2_user_proto_depIdxs,
		MessageInfos:      file_im_cloud_user_v2_user_proto_msgTypes,
	}.Build()
	File_im_cloud_user_v2_user_proto = out.File
	file_im_cloud_user_v2_user_proto_rawDesc = nil
	file_im_cloud_user_v2_user_proto_goTypes = nil
	file_im_cloud_user_v2_user_proto_depIdxs = nil
}
