// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/user_center/auth/v2/auth_error.proto

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

// 错误
type ErrorReason int32

const (
	// 为某个枚举单独设置错误码
	ErrorReason_REASON_PARAMS           ErrorReason = 0 // 参数错误
	ErrorReason_REASON_SYSTEM           ErrorReason = 1 // 系统错误
	ErrorReason_REASON_NAME_OR_PASSWORD ErrorReason = 2 // 用户名或者密码错误
	ErrorReason_REASON_CODE_IS_INVALID  ErrorReason = 3 // 验证码错误
	ErrorReason_REASON_LOGIN_STATUS     ErrorReason = 4 // 登录失败,用户状态异常
	ErrorReason_REASON_TOKEN_VERIFY     ErrorReason = 5 // token 验证失败
	ErrorReason_REASON_TOKEN_REFRESH    ErrorReason = 6 // token 刷新失败
	ErrorReason_REASON_CODE_TYPE        ErrorReason = 7 // 验证码类型错误
	// 验证码错误
	ErrorReason_REASON_VERIFY_CODE_ERROR ErrorReason = 8 //REASON
	// 用户禁用
	ErrorReason_USER_DISABLE_LOGIN ErrorReason = 9
	// 验证码已过期
	ErrorReason_VERIFICATION_CODE_HAS_EXPIRED ErrorReason = 10
	// 注册账号失败
	ErrorReason_REGISTER_ACCOUNT_ERROR ErrorReason = 11
	// 更新密码失败
	ErrorReason_UPDATE_ACCOUNT_PASSWORD_ERROR ErrorReason = 12
	// 用户不存在
	ErrorReason_USER_NOT_FOND ErrorReason = 13
	// 密码为空
	ErrorReason_PASSWORD_IS_EMPTY ErrorReason = 14
	// 用户已经实名认证过
	ErrorReason_USER_HAD_REAL_NAMED ErrorReason = 15
	// 身份证号已被其他账号绑定
	ErrorReason_ID_CARD_HAS_BIND_TO_ANOTHER ErrorReason = 16
	// 实名认证失败
	ErrorReason_REAL_NAME_FAILED ErrorReason = 17
	// 身份认证接口调用失败
	ErrorReason_CALL_IDENTIFY_INTERFACE_FAILED ErrorReason = 18
	// 输入手机号与当前绑定手机号码不一致
	ErrorReason_INPUT_PHONE_NOT_IS_USER_BINDING_PHONE ErrorReason = 19
	// 邮箱格式不正确
	ErrorReason_EMAIL_FORMAT_IS_CORRECT ErrorReason = 20
	// 邮箱绑定失败
	ErrorReason_MAIL_BINDING_ERROR ErrorReason = 21
	// 手机号绑定失败
	ErrorReason_PHONE_BINDING_ERROR           ErrorReason = 22
	ErrorReason_REASON_TOKEN_DESTRUCTION      ErrorReason = 23 // token 作废失败
	ErrorReason_REASON_TOKEN_VERIFY_ERROR     ErrorReason = 24 //token验证失败
	ErrorReason_REASON_TOKEN_HAS_EXPIRED      ErrorReason = 25 //token过期
	ErrorReason_REASON_GENERATE_TOKEN_ERROR   ErrorReason = 26 //token生成失败
	ErrorReason_UNSUPPORTED_VERIFICATION_TYPE ErrorReason = 27 //不支持的验证方式
	// 手机号错误
	ErrorReason_PHONE_ERROR ErrorReason = 28
	// 手机号错误
	ErrorReason_EMAIL_ERROR ErrorReason = 29
	// 发送验证码失败
	ErrorReason_SEND_CODE_ERROR ErrorReason = 30
	// 推广码无效
	ErrorReason_THIS_SHARE_CODE_IS_INVALID ErrorReason = 31
	// 身份核验不一致
	ErrorReason_INCONSISTENT_IDENTITY_VERIFICATION ErrorReason = 32
	// 邮箱已被绑定
	ErrorReason_THE_EMAIL_HAS_BEEN_BOUND ErrorReason = 33
	// 无可用的节点信息
	ErrorReason_NO_AVAILABLE_SERVERS ErrorReason = 34
	// 不支持的平台
	ErrorReason_UNSUPPORTED_PLATFORMS           ErrorReason = 35
	ErrorReason_REASON_AUTHORIZATION_CODE_ERROR ErrorReason = 36 //生产授权码失败
	// 获取服务节点错误
	ErrorReason_GET_SERVERS_NODE_ERR ErrorReason = 37
	// 获取平台节点错误
	ErrorReason_GET_PLATFORM_NODE_ERR ErrorReason = 38
	// 同步用户信息失败
	ErrorReason_SYNC_USER_INFO_ERR ErrorReason = 39
	// 用户未加入节点
	ErrorReason_USER_NOT_JOIN_NODE ErrorReason = 40
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:  "REASON_PARAMS",
		1:  "REASON_SYSTEM",
		2:  "REASON_NAME_OR_PASSWORD",
		3:  "REASON_CODE_IS_INVALID",
		4:  "REASON_LOGIN_STATUS",
		5:  "REASON_TOKEN_VERIFY",
		6:  "REASON_TOKEN_REFRESH",
		7:  "REASON_CODE_TYPE",
		8:  "REASON_VERIFY_CODE_ERROR",
		9:  "USER_DISABLE_LOGIN",
		10: "VERIFICATION_CODE_HAS_EXPIRED",
		11: "REGISTER_ACCOUNT_ERROR",
		12: "UPDATE_ACCOUNT_PASSWORD_ERROR",
		13: "USER_NOT_FOND",
		14: "PASSWORD_IS_EMPTY",
		15: "USER_HAD_REAL_NAMED",
		16: "ID_CARD_HAS_BIND_TO_ANOTHER",
		17: "REAL_NAME_FAILED",
		18: "CALL_IDENTIFY_INTERFACE_FAILED",
		19: "INPUT_PHONE_NOT_IS_USER_BINDING_PHONE",
		20: "EMAIL_FORMAT_IS_CORRECT",
		21: "MAIL_BINDING_ERROR",
		22: "PHONE_BINDING_ERROR",
		23: "REASON_TOKEN_DESTRUCTION",
		24: "REASON_TOKEN_VERIFY_ERROR",
		25: "REASON_TOKEN_HAS_EXPIRED",
		26: "REASON_GENERATE_TOKEN_ERROR",
		27: "UNSUPPORTED_VERIFICATION_TYPE",
		28: "PHONE_ERROR",
		29: "EMAIL_ERROR",
		30: "SEND_CODE_ERROR",
		31: "THIS_SHARE_CODE_IS_INVALID",
		32: "INCONSISTENT_IDENTITY_VERIFICATION",
		33: "THE_EMAIL_HAS_BEEN_BOUND",
		34: "NO_AVAILABLE_SERVERS",
		35: "UNSUPPORTED_PLATFORMS",
		36: "REASON_AUTHORIZATION_CODE_ERROR",
		37: "GET_SERVERS_NODE_ERR",
		38: "GET_PLATFORM_NODE_ERR",
		39: "SYNC_USER_INFO_ERR",
		40: "USER_NOT_JOIN_NODE",
	}
	ErrorReason_value = map[string]int32{
		"REASON_PARAMS":                         0,
		"REASON_SYSTEM":                         1,
		"REASON_NAME_OR_PASSWORD":               2,
		"REASON_CODE_IS_INVALID":                3,
		"REASON_LOGIN_STATUS":                   4,
		"REASON_TOKEN_VERIFY":                   5,
		"REASON_TOKEN_REFRESH":                  6,
		"REASON_CODE_TYPE":                      7,
		"REASON_VERIFY_CODE_ERROR":              8,
		"USER_DISABLE_LOGIN":                    9,
		"VERIFICATION_CODE_HAS_EXPIRED":         10,
		"REGISTER_ACCOUNT_ERROR":                11,
		"UPDATE_ACCOUNT_PASSWORD_ERROR":         12,
		"USER_NOT_FOND":                         13,
		"PASSWORD_IS_EMPTY":                     14,
		"USER_HAD_REAL_NAMED":                   15,
		"ID_CARD_HAS_BIND_TO_ANOTHER":           16,
		"REAL_NAME_FAILED":                      17,
		"CALL_IDENTIFY_INTERFACE_FAILED":        18,
		"INPUT_PHONE_NOT_IS_USER_BINDING_PHONE": 19,
		"EMAIL_FORMAT_IS_CORRECT":               20,
		"MAIL_BINDING_ERROR":                    21,
		"PHONE_BINDING_ERROR":                   22,
		"REASON_TOKEN_DESTRUCTION":              23,
		"REASON_TOKEN_VERIFY_ERROR":             24,
		"REASON_TOKEN_HAS_EXPIRED":              25,
		"REASON_GENERATE_TOKEN_ERROR":           26,
		"UNSUPPORTED_VERIFICATION_TYPE":         27,
		"PHONE_ERROR":                           28,
		"EMAIL_ERROR":                           29,
		"SEND_CODE_ERROR":                       30,
		"THIS_SHARE_CODE_IS_INVALID":            31,
		"INCONSISTENT_IDENTITY_VERIFICATION":    32,
		"THE_EMAIL_HAS_BEEN_BOUND":              33,
		"NO_AVAILABLE_SERVERS":                  34,
		"UNSUPPORTED_PLATFORMS":                 35,
		"REASON_AUTHORIZATION_CODE_ERROR":       36,
		"GET_SERVERS_NODE_ERR":                  37,
		"GET_PLATFORM_NODE_ERR":                 38,
		"SYNC_USER_INFO_ERR":                    39,
		"USER_NOT_JOIN_NODE":                    40,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_center_auth_v2_auth_error_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_user_center_auth_v2_auth_error_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_user_center_auth_v2_auth_error_proto_rawDescGZIP(), []int{0}
}

var File_api_user_center_auth_v2_auth_error_proto protoreflect.FileDescriptor

var file_api_user_center_auth_v2_auth_error_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x32, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf4, 0x0a, 0x0a, 0x0b, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x0d, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x53, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x90,
	0x03, 0x12, 0x17, 0x0a, 0x0d, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x53, 0x54,
	0x45, 0x4d, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x21, 0x0a, 0x17, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x50, 0x41, 0x53,
	0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x20, 0x0a,
	0x16, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x53, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x1d, 0x0a, 0x13, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1d,
	0x0a, 0x13, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x59, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1e, 0x0a,
	0x14, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x52, 0x45,
	0x46, 0x52, 0x45, 0x53, 0x48, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1a, 0x0a,
	0x10, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x22, 0x0a, 0x18, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x08, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1c, 0x0a,
	0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x10, 0x09, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x27, 0x0a, 0x1d, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x48, 0x41, 0x53, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0a, 0x1a, 0x04,
	0xa8, 0x45, 0xf4, 0x03, 0x12, 0x20, 0x0a, 0x16, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0b,
	0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x27, 0x0a, 0x1d, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12,
	0x17, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x4e, 0x44,
	0x10, 0x0d, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x50, 0x41, 0x53, 0x53,
	0x57, 0x4f, 0x52, 0x44, 0x5f, 0x49, 0x53, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x0e, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x48, 0x41,
	0x44, 0x5f, 0x52, 0x45, 0x41, 0x4c, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x44, 0x10, 0x0f, 0x1a, 0x04,
	0xa8, 0x45, 0x90, 0x03, 0x12, 0x25, 0x0a, 0x1b, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f,
	0x48, 0x41, 0x53, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x41, 0x4e, 0x4f, 0x54,
	0x48, 0x45, 0x52, 0x10, 0x10, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1a, 0x0a, 0x10, 0x52,
	0x45, 0x41, 0x4c, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x11, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x28, 0x0a, 0x1e, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x46, 0x41,
	0x43, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x12, 0x1a, 0x04, 0xa8, 0x45, 0xf4,
	0x03, 0x12, 0x2f, 0x0a, 0x25, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x42, 0x49, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x13, 0x1a, 0x04, 0xa8, 0x45,
	0x90, 0x03, 0x12, 0x21, 0x0a, 0x17, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x43, 0x54, 0x10, 0x14, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x42, 0x49,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x15, 0x1a, 0x04, 0xa8,
	0x45, 0xf4, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x42, 0x49, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x16, 0x1a, 0x04, 0xa8, 0x45,
	0xf4, 0x03, 0x12, 0x22, 0x0a, 0x18, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x17,
	0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x23, 0x0a, 0x19, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x18, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x22, 0x0a, 0x18, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x48, 0x41, 0x53, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x19, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x25, 0x0a, 0x1b, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41,
	0x54, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x1a,
	0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x27, 0x0a, 0x1d, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50,
	0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x1b, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x15, 0x0a, 0x0b, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x1c,
	0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x15, 0x0a, 0x0b, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x1d, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x19, 0x0a,
	0x0f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x1e, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x24, 0x0a, 0x1a, 0x54, 0x48, 0x49, 0x53,
	0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x1f, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x2c,
	0x0a, 0x22, 0x49, 0x4e, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x49,
	0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x20, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x22, 0x0a, 0x18,
	0x54, 0x48, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x42, 0x45,
	0x45, 0x4e, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x21, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03,
	0x12, 0x1e, 0x0a, 0x14, 0x4e, 0x4f, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x53, 0x10, 0x22, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03,
	0x12, 0x1f, 0x0a, 0x15, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f,
	0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x53, 0x10, 0x23, 0x1a, 0x04, 0xa8, 0x45, 0x90,
	0x03, 0x12, 0x29, 0x0a, 0x1f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x41, 0x55, 0x54, 0x48,
	0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x24, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1e, 0x0a, 0x14,
	0x47, 0x45, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x53, 0x5f, 0x4e, 0x4f, 0x44, 0x45,
	0x5f, 0x45, 0x52, 0x52, 0x10, 0x25, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1f, 0x0a, 0x15,
	0x47, 0x45, 0x54, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4e, 0x4f, 0x44,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x26, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1c, 0x0a,
	0x12, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f,
	0x45, 0x52, 0x52, 0x10, 0x27, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x4e, 0x4f, 0x44,
	0x45, 0x10, 0x28, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42,
	0x57, 0x0a, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_center_auth_v2_auth_error_proto_rawDescOnce sync.Once
	file_api_user_center_auth_v2_auth_error_proto_rawDescData = file_api_user_center_auth_v2_auth_error_proto_rawDesc
)

func file_api_user_center_auth_v2_auth_error_proto_rawDescGZIP() []byte {
	file_api_user_center_auth_v2_auth_error_proto_rawDescOnce.Do(func() {
		file_api_user_center_auth_v2_auth_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_center_auth_v2_auth_error_proto_rawDescData)
	})
	return file_api_user_center_auth_v2_auth_error_proto_rawDescData
}

var file_api_user_center_auth_v2_auth_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_user_center_auth_v2_auth_error_proto_goTypes = []any{
	(ErrorReason)(0), // 0: api.user_center.auth.v2.ErrorReason
}
var file_api_user_center_auth_v2_auth_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_center_auth_v2_auth_error_proto_init() }
func file_api_user_center_auth_v2_auth_error_proto_init() {
	if File_api_user_center_auth_v2_auth_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_center_auth_v2_auth_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_center_auth_v2_auth_error_proto_goTypes,
		DependencyIndexes: file_api_user_center_auth_v2_auth_error_proto_depIdxs,
		EnumInfos:         file_api_user_center_auth_v2_auth_error_proto_enumTypes,
	}.Build()
	File_api_user_center_auth_v2_auth_error_proto = out.File
	file_api_user_center_auth_v2_auth_error_proto_rawDesc = nil
	file_api_user_center_auth_v2_auth_error_proto_goTypes = nil
	file_api_user_center_auth_v2_auth_error_proto_depIdxs = nil
}
