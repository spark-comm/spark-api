// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: api/im_cloud/user/v2/profile_error.proto

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

type ErrorProfileReason int32

const (
	// 用户信息未找到
	ErrorProfileReason_PROFILE_NOT_FOUND ErrorProfileReason = 0
	// 获取数据失败
	ErrorProfileReason_GET_DATE ErrorProfileReason = 2
	// 编辑用户异常
	ErrorProfileReason_FAILED_TO_EDIT_PROFILE ErrorProfileReason = 3
	// 手机号不能为空
	ErrorProfileReason_PHONE_IS_NOT_EMPTY ErrorProfileReason = 4
	// 搜索用户错误
	ErrorProfileReason_SEARCH_PROFILE_ERR ErrorProfileReason = 5
	// 缓存数据未找到
	ErrorProfileReason_CACHE_DATA_Not_Found ErrorProfileReason = 6
	// 删除用户信息失败
	ErrorProfileReason_FAILED_DELETE_PROFILE ErrorProfileReason = 7
	// 设置私密聊天密码失败
	ErrorProfileReason_SET_PRIVACY_PASSWORD_FAIL ErrorProfileReason = 8
	// 未设置私密聊天密码
	ErrorProfileReason_NOT_SET_PRIVACY_PASSWORD ErrorProfileReason = 9
	// 历史密码验证失败
	ErrorProfileReason_HISTORY_PRIVACY_PASSWORD_VERIFICATION_FAIL ErrorProfileReason = 10
	// 输入邮箱与用户绑定邮箱不一致
	ErrorProfileReason_PARAMS_EMAIL_NOT_IS_USER_BIND ErrorProfileReason = 11
	// 输入手机与用户绑定手机不一致
	ErrorProfileReason_PARAMS_PHONE_NOT_IS_USER_BIND ErrorProfileReason = 12
	// 用户非客户
	ErrorProfileReason_USER_IS_NOT_CUSTOMER ErrorProfileReason = 13
	// 该推广码无效
	ErrorProfileReason_THIS_SHARE_CODE_IS_INVALID ErrorProfileReason = 14
)

// Enum value maps for ErrorProfileReason.
var (
	ErrorProfileReason_name = map[int32]string{
		0:  "PROFILE_NOT_FOUND",
		2:  "GET_DATE",
		3:  "FAILED_TO_EDIT_PROFILE",
		4:  "PHONE_IS_NOT_EMPTY",
		5:  "SEARCH_PROFILE_ERR",
		6:  "CACHE_DATA_Not_Found",
		7:  "FAILED_DELETE_PROFILE",
		8:  "SET_PRIVACY_PASSWORD_FAIL",
		9:  "NOT_SET_PRIVACY_PASSWORD",
		10: "HISTORY_PRIVACY_PASSWORD_VERIFICATION_FAIL",
		11: "PARAMS_EMAIL_NOT_IS_USER_BIND",
		12: "PARAMS_PHONE_NOT_IS_USER_BIND",
		13: "USER_IS_NOT_CUSTOMER",
		14: "THIS_SHARE_CODE_IS_INVALID",
	}
	ErrorProfileReason_value = map[string]int32{
		"PROFILE_NOT_FOUND":                          0,
		"GET_DATE":                                   2,
		"FAILED_TO_EDIT_PROFILE":                     3,
		"PHONE_IS_NOT_EMPTY":                         4,
		"SEARCH_PROFILE_ERR":                         5,
		"CACHE_DATA_Not_Found":                       6,
		"FAILED_DELETE_PROFILE":                      7,
		"SET_PRIVACY_PASSWORD_FAIL":                  8,
		"NOT_SET_PRIVACY_PASSWORD":                   9,
		"HISTORY_PRIVACY_PASSWORD_VERIFICATION_FAIL": 10,
		"PARAMS_EMAIL_NOT_IS_USER_BIND":              11,
		"PARAMS_PHONE_NOT_IS_USER_BIND":              12,
		"USER_IS_NOT_CUSTOMER":                       13,
		"THIS_SHARE_CODE_IS_INVALID":                 14,
	}
)

func (x ErrorProfileReason) Enum() *ErrorProfileReason {
	p := new(ErrorProfileReason)
	*p = x
	return p
}

func (x ErrorProfileReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorProfileReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_im_cloud_user_v2_profile_error_proto_enumTypes[0].Descriptor()
}

func (ErrorProfileReason) Type() protoreflect.EnumType {
	return &file_api_im_cloud_user_v2_profile_error_proto_enumTypes[0]
}

func (x ErrorProfileReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorProfileReason.Descriptor instead.
func (ErrorProfileReason) EnumDescriptor() ([]byte, []int) {
	return file_api_im_cloud_user_v2_profile_error_proto_rawDescGZIP(), []int{0}
}

var File_api_im_cloud_user_v2_profile_error_proto protoreflect.FileDescriptor

var file_api_im_cloud_user_v2_profile_error_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xfb, 0x03, 0x0a, 0x12, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x11,
	0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x12, 0x0a, 0x08, 0x47, 0x45, 0x54,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xf8, 0x03, 0x12, 0x20, 0x0a,
	0x16, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x5f,
	0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12,
	0x1c, 0x0a, 0x12, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf6, 0x03, 0x12, 0x1c, 0x0a,
	0x12, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f,
	0x45, 0x52, 0x52, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x43,
	0x41, 0x43, 0x48, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4e, 0x6f, 0x74, 0x5f, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f,
	0x46, 0x49, 0x4c, 0x45, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0xf9, 0x03, 0x12, 0x23, 0x0a, 0x19,
	0x53, 0x45, 0x54, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x43, 0x59, 0x5f, 0x50, 0x41, 0x53, 0x53,
	0x57, 0x4f, 0x52, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x08, 0x1a, 0x04, 0xa8, 0x45, 0xf4,
	0x03, 0x12, 0x22, 0x0a, 0x18, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x50, 0x52, 0x49,
	0x56, 0x41, 0x43, 0x59, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x09, 0x1a,
	0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x34, 0x0a, 0x2a, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59,
	0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x43, 0x59, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x10, 0x0a, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x27, 0x0a, 0x1d, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x53, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x10, 0x0b, 0x1a, 0x04,
	0xa8, 0x45, 0xf4, 0x03, 0x12, 0x27, 0x0a, 0x1d, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x53, 0x5f, 0x50,
	0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x42, 0x49, 0x4e, 0x44, 0x10, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x18, 0x0a,
	0x14, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x0d, 0x12, 0x24, 0x0a, 0x1a, 0x54, 0x48, 0x49, 0x53, 0x5f,
	0x53, 0x48, 0x41, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x0e, 0x1a, 0x04, 0xa8, 0x45, 0xf8, 0x03, 0x1a, 0x04, 0xa0,
	0x45, 0xf4, 0x03, 0x42, 0x51, 0x0a, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6d, 0x5f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x37, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d,
	0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x6d, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_im_cloud_user_v2_profile_error_proto_rawDescOnce sync.Once
	file_api_im_cloud_user_v2_profile_error_proto_rawDescData = file_api_im_cloud_user_v2_profile_error_proto_rawDesc
)

func file_api_im_cloud_user_v2_profile_error_proto_rawDescGZIP() []byte {
	file_api_im_cloud_user_v2_profile_error_proto_rawDescOnce.Do(func() {
		file_api_im_cloud_user_v2_profile_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_im_cloud_user_v2_profile_error_proto_rawDescData)
	})
	return file_api_im_cloud_user_v2_profile_error_proto_rawDescData
}

var file_api_im_cloud_user_v2_profile_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_im_cloud_user_v2_profile_error_proto_goTypes = []interface{}{
	(ErrorProfileReason)(0), // 0: api.im_cloud.user.v2.ErrorProfileReason
}
var file_api_im_cloud_user_v2_profile_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_im_cloud_user_v2_profile_error_proto_init() }
func file_api_im_cloud_user_v2_profile_error_proto_init() {
	if File_api_im_cloud_user_v2_profile_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_im_cloud_user_v2_profile_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_im_cloud_user_v2_profile_error_proto_goTypes,
		DependencyIndexes: file_api_im_cloud_user_v2_profile_error_proto_depIdxs,
		EnumInfos:         file_api_im_cloud_user_v2_profile_error_proto_enumTypes,
	}.Build()
	File_api_im_cloud_user_v2_profile_error_proto = out.File
	file_api_im_cloud_user_v2_profile_error_proto_rawDesc = nil
	file_api_im_cloud_user_v2_profile_error_proto_goTypes = nil
	file_api_im_cloud_user_v2_profile_error_proto_depIdxs = nil
}
