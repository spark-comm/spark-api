// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: api/platform/task/v2/task_error.proto

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

type ErrorTaskReason int32

const (
	// 节点信息未找到
	ErrorTaskReason_NODE_NOT_FOUND ErrorTaskReason = 0
	// 添加任务失败
	ErrorTaskReason_ADD_TASK_ERROR ErrorTaskReason = 1
	// 执行任务失败
	ErrorTaskReason_EXECUTION_TASK_ERROR ErrorTaskReason = 3
)

// Enum value maps for ErrorTaskReason.
var (
	ErrorTaskReason_name = map[int32]string{
		0: "NODE_NOT_FOUND",
		1: "ADD_TASK_ERROR",
		3: "EXECUTION_TASK_ERROR",
	}
	ErrorTaskReason_value = map[string]int32{
		"NODE_NOT_FOUND":       0,
		"ADD_TASK_ERROR":       1,
		"EXECUTION_TASK_ERROR": 3,
	}
)

func (x ErrorTaskReason) Enum() *ErrorTaskReason {
	p := new(ErrorTaskReason)
	*p = x
	return p
}

func (x ErrorTaskReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorTaskReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_platform_task_v2_task_error_proto_enumTypes[0].Descriptor()
}

func (ErrorTaskReason) Type() protoreflect.EnumType {
	return &file_api_platform_task_v2_task_error_proto_enumTypes[0]
}

func (x ErrorTaskReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorTaskReason.Descriptor instead.
func (ErrorTaskReason) EnumDescriptor() ([]byte, []int) {
	return file_api_platform_task_v2_task_error_proto_rawDescGZIP(), []int{0}
}

var File_api_platform_task_v2_task_error_proto protoreflect.FileDescriptor

var file_api_platform_task_v2_task_error_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x32, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x6b, 0x0a, 0x0f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x0e, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12,
	0x18, 0x0a, 0x0e, 0x41, 0x44, 0x44, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x45, 0x58, 0x45,
	0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42,
	0x51, 0x0a, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x32, 0x3b,
	0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_platform_task_v2_task_error_proto_rawDescOnce sync.Once
	file_api_platform_task_v2_task_error_proto_rawDescData = file_api_platform_task_v2_task_error_proto_rawDesc
)

func file_api_platform_task_v2_task_error_proto_rawDescGZIP() []byte {
	file_api_platform_task_v2_task_error_proto_rawDescOnce.Do(func() {
		file_api_platform_task_v2_task_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_platform_task_v2_task_error_proto_rawDescData)
	})
	return file_api_platform_task_v2_task_error_proto_rawDescData
}

var file_api_platform_task_v2_task_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_platform_task_v2_task_error_proto_goTypes = []interface{}{
	(ErrorTaskReason)(0), // 0: api.platform.task.v2.ErrorTaskReason
}
var file_api_platform_task_v2_task_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_platform_task_v2_task_error_proto_init() }
func file_api_platform_task_v2_task_error_proto_init() {
	if File_api_platform_task_v2_task_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_platform_task_v2_task_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_platform_task_v2_task_error_proto_goTypes,
		DependencyIndexes: file_api_platform_task_v2_task_error_proto_depIdxs,
		EnumInfos:         file_api_platform_task_v2_task_error_proto_enumTypes,
	}.Build()
	File_api_platform_task_v2_task_error_proto = out.File
	file_api_platform_task_v2_task_error_proto_rawDesc = nil
	file_api_platform_task_v2_task_error_proto_goTypes = nil
	file_api_platform_task_v2_task_error_proto_depIdxs = nil
}
