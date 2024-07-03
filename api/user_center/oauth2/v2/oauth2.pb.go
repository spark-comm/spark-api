// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: user_center/oauth2/v2/oauth2.proto

package v2

import (
	v2 "github.com/spark-comm/spark-api/api/user_center/auth/v2"
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

type AccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	NodeId string `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *AccessTokenRequest) Reset() {
	*x = AccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_oauth2_v2_oauth2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokenRequest) ProtoMessage() {}

func (x *AccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_oauth2_v2_oauth2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokenRequest.ProtoReflect.Descriptor instead.
func (*AccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_center_oauth2_v2_oauth2_proto_rawDescGZIP(), []int{0}
}

func (x *AccessTokenRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AccessTokenRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *AccessTokenRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type AccessTokenReplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken *v2.AuthToken `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *AccessTokenReplay) Reset() {
	*x = AccessTokenReplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_oauth2_v2_oauth2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessTokenReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokenReplay) ProtoMessage() {}

func (x *AccessTokenReplay) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_oauth2_v2_oauth2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokenReplay.ProtoReflect.Descriptor instead.
func (*AccessTokenReplay) Descriptor() ([]byte, []int) {
	return file_user_center_oauth2_v2_oauth2_proto_rawDescGZIP(), []int{1}
}

func (x *AccessTokenReplay) GetAuthToken() *v2.AuthToken {
	if x != nil {
		return x.AuthToken
	}
	return nil
}

type QrConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId       string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	RedirectUri  string `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	ResponseType string `protobuf:"bytes,3,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	Scope        string `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	State        string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *QrConnectRequest) Reset() {
	*x = QrConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_oauth2_v2_oauth2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QrConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QrConnectRequest) ProtoMessage() {}

func (x *QrConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_oauth2_v2_oauth2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QrConnectRequest.ProtoReflect.Descriptor instead.
func (*QrConnectRequest) Descriptor() ([]byte, []int) {
	return file_user_center_oauth2_v2_oauth2_proto_rawDescGZIP(), []int{2}
}

func (x *QrConnectRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *QrConnectRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *QrConnectRequest) GetResponseType() string {
	if x != nil {
		return x.ResponseType
	}
	return ""
}

func (x *QrConnectRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *QrConnectRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type QrConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QrConnectReply) Reset() {
	*x = QrConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_oauth2_v2_oauth2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QrConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QrConnectReply) ProtoMessage() {}

func (x *QrConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_oauth2_v2_oauth2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QrConnectReply.ProtoReflect.Descriptor instead.
func (*QrConnectReply) Descriptor() ([]byte, []int) {
	return file_user_center_oauth2_v2_oauth2_proto_rawDescGZIP(), []int{3}
}

var File_user_center_oauth2_v2_oauth2_proto protoreflect.FileDescriptor

var file_user_center_oauth2_v2_oauth2_proto_rawDesc = []byte{
	0x0a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x32, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x32, 0x1a,
	0x1e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x59, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x11, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12,
	0x41, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x51, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x69, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x51, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x74, 0x0a, 0x06, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x12, 0x6a, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x42, 0x5b, 0x0a, 0x19,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x32, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_center_oauth2_v2_oauth2_proto_rawDescOnce sync.Once
	file_user_center_oauth2_v2_oauth2_proto_rawDescData = file_user_center_oauth2_v2_oauth2_proto_rawDesc
)

func file_user_center_oauth2_v2_oauth2_proto_rawDescGZIP() []byte {
	file_user_center_oauth2_v2_oauth2_proto_rawDescOnce.Do(func() {
		file_user_center_oauth2_v2_oauth2_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_center_oauth2_v2_oauth2_proto_rawDescData)
	})
	return file_user_center_oauth2_v2_oauth2_proto_rawDescData
}

var file_user_center_oauth2_v2_oauth2_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_center_oauth2_v2_oauth2_proto_goTypes = []interface{}{
	(*AccessTokenRequest)(nil), // 0: api.user_center.oauth2.v2.AccessTokenRequest
	(*AccessTokenReplay)(nil),  // 1: api.user_center.oauth2.v2.AccessTokenReplay
	(*QrConnectRequest)(nil),   // 2: api.user_center.oauth2.v2.QrConnectRequest
	(*QrConnectReply)(nil),     // 3: api.user_center.oauth2.v2.QrConnectReply
	(*v2.AuthToken)(nil),       // 4: api.user_center.auth.v2.AuthToken
}
var file_user_center_oauth2_v2_oauth2_proto_depIdxs = []int32{
	4, // 0: api.user_center.oauth2.v2.AccessTokenReplay.auth_token:type_name -> api.user_center.auth.v2.AuthToken
	0, // 1: api.user_center.oauth2.v2.Oauth2.AccessToken:input_type -> api.user_center.oauth2.v2.AccessTokenRequest
	1, // 2: api.user_center.oauth2.v2.Oauth2.AccessToken:output_type -> api.user_center.oauth2.v2.AccessTokenReplay
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_center_oauth2_v2_oauth2_proto_init() }
func file_user_center_oauth2_v2_oauth2_proto_init() {
	if File_user_center_oauth2_v2_oauth2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_center_oauth2_v2_oauth2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessTokenRequest); i {
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
		file_user_center_oauth2_v2_oauth2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessTokenReplay); i {
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
		file_user_center_oauth2_v2_oauth2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QrConnectRequest); i {
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
		file_user_center_oauth2_v2_oauth2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QrConnectReply); i {
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
			RawDescriptor: file_user_center_oauth2_v2_oauth2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_center_oauth2_v2_oauth2_proto_goTypes,
		DependencyIndexes: file_user_center_oauth2_v2_oauth2_proto_depIdxs,
		MessageInfos:      file_user_center_oauth2_v2_oauth2_proto_msgTypes,
	}.Build()
	File_user_center_oauth2_v2_oauth2_proto = out.File
	file_user_center_oauth2_v2_oauth2_proto_rawDesc = nil
	file_user_center_oauth2_v2_oauth2_proto_goTypes = nil
	file_user_center_oauth2_v2_oauth2_proto_depIdxs = nil
}
