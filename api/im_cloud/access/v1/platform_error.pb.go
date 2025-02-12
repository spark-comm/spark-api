// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/im_cloud/access/v1/platform_error.proto

// 定义包名

package v2

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorPlatformReason int32

const (
	// 获取密钥失败
	ErrorPlatformReason_GET_NODE_KEY_ERR ErrorPlatformReason = 0
	// 下发主节点信息失败
	ErrorPlatformReason_DISTRIBUTE_MASTER_NODE_CONFIG_ERR ErrorPlatformReason = 1
	// 下发代理节点信息失败
	ErrorPlatformReason_DISTRIBUTE_PROXY_NODE_CONFIG_ERR ErrorPlatformReason = 2
	// 获取节点状态失败
	ErrorPlatformReason_NODE_STATUS_ERR ErrorPlatformReason = 3
	// 下发用户失败
	ErrorPlatformReason_DISTRIBUTE_USER_ERR ErrorPlatformReason = 4
	// 上报用户失败
	ErrorPlatformReason_ESCALATION_USER_ERR ErrorPlatformReason = 5
)

// Enum value maps for ErrorPlatformReason.
var (
	ErrorPlatformReason_name = map[int32]string{
		0: "GET_NODE_KEY_ERR",
		1: "DISTRIBUTE_MASTER_NODE_CONFIG_ERR",
		2: "DISTRIBUTE_PROXY_NODE_CONFIG_ERR",
		3: "NODE_STATUS_ERR",
		4: "DISTRIBUTE_USER_ERR",
		5: "ESCALATION_USER_ERR",
	}
	ErrorPlatformReason_value = map[string]int32{
		"GET_NODE_KEY_ERR":                  0,
		"DISTRIBUTE_MASTER_NODE_CONFIG_ERR": 1,
		"DISTRIBUTE_PROXY_NODE_CONFIG_ERR":  2,
		"NODE_STATUS_ERR":                   3,
		"DISTRIBUTE_USER_ERR":               4,
		"ESCALATION_USER_ERR":               5,
	}
)

func (x ErrorPlatformReason) Enum() *ErrorPlatformReason {
	p := new(ErrorPlatformReason)
	*p = x
	return p
}

func (x ErrorPlatformReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorPlatformReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_im_cloud_access_v1_platform_error_proto_enumTypes[0].Descriptor()
}

func (ErrorPlatformReason) Type() protoreflect.EnumType {
	return &file_api_im_cloud_access_v1_platform_error_proto_enumTypes[0]
}

func (x ErrorPlatformReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorPlatformReason.Descriptor instead.
func (ErrorPlatformReason) EnumDescriptor() ([]byte, []int) {
	return file_api_im_cloud_access_v1_platform_error_proto_rawDescGZIP(), []int{0}
}

var File_api_im_cloud_access_v1_platform_error_proto protoreflect.FileDescriptor

var file_api_im_cloud_access_v1_platform_error_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xe9, 0x01, 0x0a, 0x13, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x10, 0x47, 0x45, 0x54, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x4b,
	0x45, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x2b,
	0x0a, 0x21, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x4d, 0x41, 0x53,
	0x54, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f,
	0x45, 0x52, 0x52, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x2a, 0x0a, 0x20, 0x44,
	0x49, 0x53, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f,
	0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x10,
	0x02, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x4e, 0x4f, 0x44, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45,
	0xf4, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf4,
	0x03, 0x12, 0x1d, 0x0a, 0x13, 0x45, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03,
	0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x55, 0x0a, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6d,
	0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x32,
	0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_im_cloud_access_v1_platform_error_proto_rawDescOnce sync.Once
	file_api_im_cloud_access_v1_platform_error_proto_rawDescData = file_api_im_cloud_access_v1_platform_error_proto_rawDesc
)

func file_api_im_cloud_access_v1_platform_error_proto_rawDescGZIP() []byte {
	file_api_im_cloud_access_v1_platform_error_proto_rawDescOnce.Do(func() {
		file_api_im_cloud_access_v1_platform_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_im_cloud_access_v1_platform_error_proto_rawDescData)
	})
	return file_api_im_cloud_access_v1_platform_error_proto_rawDescData
}

var file_api_im_cloud_access_v1_platform_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_im_cloud_access_v1_platform_error_proto_goTypes = []any{
	(ErrorPlatformReason)(0), // 0: api.im_cloud.access.v2.ErrorPlatformReason
}
var file_api_im_cloud_access_v1_platform_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_im_cloud_access_v1_platform_error_proto_init() }
func file_api_im_cloud_access_v1_platform_error_proto_init() {
	if File_api_im_cloud_access_v1_platform_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_im_cloud_access_v1_platform_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_im_cloud_access_v1_platform_error_proto_goTypes,
		DependencyIndexes: file_api_im_cloud_access_v1_platform_error_proto_depIdxs,
		EnumInfos:         file_api_im_cloud_access_v1_platform_error_proto_enumTypes,
	}.Build()
	File_api_im_cloud_access_v1_platform_error_proto = out.File
	file_api_im_cloud_access_v1_platform_error_proto_rawDesc = nil
	file_api_im_cloud_access_v1_platform_error_proto_goTypes = nil
	file_api_im_cloud_access_v1_platform_error_proto_depIdxs = nil
}
