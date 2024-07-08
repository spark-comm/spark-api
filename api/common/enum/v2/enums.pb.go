// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: common/enum/v2/enums.proto

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

// 账号类型
type AccountType int32

const (
	AccountType_Phone    AccountType = 0 //手机号
	AccountType_Email    AccountType = 1 //邮箱
	AccountType_UserName AccountType = 2 //用户名密码注册
)

// Enum value maps for AccountType.
var (
	AccountType_name = map[int32]string{
		0: "Phone",
		1: "Email",
		2: "UserName",
	}
	AccountType_value = map[string]int32{
		"Phone":    0,
		"Email":    1,
		"UserName": 2,
	}
)

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}

func (x AccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_v2_enums_proto_enumTypes[0].Descriptor()
}

func (AccountType) Type() protoreflect.EnumType {
	return &file_common_enum_v2_enums_proto_enumTypes[0]
}

func (x AccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountType.Descriptor instead.
func (AccountType) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_v2_enums_proto_rawDescGZIP(), []int{0}
}

type Gender int32

const (
	Gender_GENDER_UNKNOWN Gender = 0 // 未知
	Gender_GENDER_FEMALE  Gender = 1 // 男
	Gender_GENDER_MALE    Gender = 2 // 女
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_UNKNOWN",
		1: "GENDER_FEMALE",
		2: "GENDER_MALE",
	}
	Gender_value = map[string]int32{
		"GENDER_UNKNOWN": 0,
		"GENDER_FEMALE":  1,
		"GENDER_MALE":    2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_v2_enums_proto_enumTypes[1].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_common_enum_v2_enums_proto_enumTypes[1]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_v2_enums_proto_rawDescGZIP(), []int{1}
}

// 账号状态
type AccountStatus int32

const (
	AccountStatus_ACCOUNT_STATUS_UNKNOWN AccountStatus = 0 // 状态未知
	AccountStatus_ACCOUNT_STATUS_NORMAL  AccountStatus = 1 // 状态正常
	AccountStatus_ACCOUNT_STATUS_FREEZE  AccountStatus = 2 // 状态冻结
)

// Enum value maps for AccountStatus.
var (
	AccountStatus_name = map[int32]string{
		0: "ACCOUNT_STATUS_UNKNOWN",
		1: "ACCOUNT_STATUS_NORMAL",
		2: "ACCOUNT_STATUS_FREEZE",
	}
	AccountStatus_value = map[string]int32{
		"ACCOUNT_STATUS_UNKNOWN": 0,
		"ACCOUNT_STATUS_NORMAL":  1,
		"ACCOUNT_STATUS_FREEZE":  2,
	}
)

func (x AccountStatus) Enum() *AccountStatus {
	p := new(AccountStatus)
	*p = x
	return p
}

func (x AccountStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_v2_enums_proto_enumTypes[2].Descriptor()
}

func (AccountStatus) Type() protoreflect.EnumType {
	return &file_common_enum_v2_enums_proto_enumTypes[2]
}

func (x AccountStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountStatus.Descriptor instead.
func (AccountStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_v2_enums_proto_rawDescGZIP(), []int{2}
}

var File_common_enum_v2_enums_proto protoreflect.FileDescriptor

var file_common_enum_v2_enums_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x76, 0x32,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x76, 0x32,
	0x2a, 0x31, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x10, 0x02, 0x2a, 0x40, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x0e, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d, 0x41,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x02, 0x2a, 0x61, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x46, 0x52, 0x45, 0x45, 0x5a, 0x45, 0x10, 0x02, 0x42, 0x4d, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x76, 0x32, 0x50, 0x01,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_enum_v2_enums_proto_rawDescOnce sync.Once
	file_common_enum_v2_enums_proto_rawDescData = file_common_enum_v2_enums_proto_rawDesc
)

func file_common_enum_v2_enums_proto_rawDescGZIP() []byte {
	file_common_enum_v2_enums_proto_rawDescOnce.Do(func() {
		file_common_enum_v2_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enum_v2_enums_proto_rawDescData)
	})
	return file_common_enum_v2_enums_proto_rawDescData
}

var file_common_enum_v2_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_enum_v2_enums_proto_goTypes = []any{
	(AccountType)(0),   // 0: api.common.enum.v2.AccountType
	(Gender)(0),        // 1: api.common.enum.v2.Gender
	(AccountStatus)(0), // 2: api.common.enum.v2.AccountStatus
}
var file_common_enum_v2_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enum_v2_enums_proto_init() }
func file_common_enum_v2_enums_proto_init() {
	if File_common_enum_v2_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enum_v2_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enum_v2_enums_proto_goTypes,
		DependencyIndexes: file_common_enum_v2_enums_proto_depIdxs,
		EnumInfos:         file_common_enum_v2_enums_proto_enumTypes,
	}.Build()
	File_common_enum_v2_enums_proto = out.File
	file_common_enum_v2_enums_proto_rawDesc = nil
	file_common_enum_v2_enums_proto_goTypes = nil
	file_common_enum_v2_enums_proto_depIdxs = nil
}
