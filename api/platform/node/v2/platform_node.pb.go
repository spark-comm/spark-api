// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: platform/node/v2/platform_node.proto

package v2

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/spark-comm/spark-api/api/common/enum/v2"
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

// PlatformNodeInfo
type PlatformNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`          // profile id 和 userId 等价
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // 地址
	Weight  int32  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`  //节点权重
}

func (x *PlatformNodeInfo) Reset() {
	*x = PlatformNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformNodeInfo) ProtoMessage() {}

func (x *PlatformNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformNodeInfo.ProtoReflect.Descriptor instead.
func (*PlatformNodeInfo) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{0}
}

func (x *PlatformNodeInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlatformNodeInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PlatformNodeInfo) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

// 新加节点
// -------------------------------------------
type AddPlatformNodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // 地址
	Weight  int32  `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`  //节点权重
}

func (x *AddPlatformNodeReq) Reset() {
	*x = AddPlatformNodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPlatformNodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlatformNodeReq) ProtoMessage() {}

func (x *AddPlatformNodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlatformNodeReq.ProtoReflect.Descriptor instead.
func (*AddPlatformNodeReq) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{1}
}

func (x *AddPlatformNodeReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddPlatformNodeReq) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type AddPlatformNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (x *AddPlatformNodeReply) Reset() {
	*x = AddPlatformNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPlatformNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlatformNodeReply) ProtoMessage() {}

func (x *AddPlatformNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlatformNodeReply.ProtoReflect.Descriptor instead.
func (*AddPlatformNodeReply) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{2}
}

func (x *AddPlatformNodeReply) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

// 删除节点
// -------------------------------------------
type DelPlatformNodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //节点id
}

func (x *DelPlatformNodeReq) Reset() {
	*x = DelPlatformNodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelPlatformNodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelPlatformNodeReq) ProtoMessage() {}

func (x *DelPlatformNodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelPlatformNodeReq.ProtoReflect.Descriptor instead.
func (*DelPlatformNodeReq) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{3}
}

func (x *DelPlatformNodeReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelPlatformNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelPlatformNodeReply) Reset() {
	*x = DelPlatformNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelPlatformNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelPlatformNodeReply) ProtoMessage() {}

func (x *DelPlatformNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelPlatformNodeReply.ProtoReflect.Descriptor instead.
func (*DelPlatformNodeReply) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{4}
}

// 编辑节点
// -------------------------------------------
type EditPlatformNodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                //节点id
	Address *string `protobuf:"bytes,2,opt,name=address,proto3,oneof" json:"address,omitempty"` // 地址
	Weight  *int32  `protobuf:"varint,3,opt,name=weight,proto3,oneof" json:"weight,omitempty"`  //节点权重
}

func (x *EditPlatformNodeReq) Reset() {
	*x = EditPlatformNodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditPlatformNodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditPlatformNodeReq) ProtoMessage() {}

func (x *EditPlatformNodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditPlatformNodeReq.ProtoReflect.Descriptor instead.
func (*EditPlatformNodeReq) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{5}
}

func (x *EditPlatformNodeReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditPlatformNodeReq) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *EditPlatformNodeReq) GetWeight() int32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

type EditPlatformNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EditPlatformNodeReply) Reset() {
	*x = EditPlatformNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditPlatformNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditPlatformNodeReply) ProtoMessage() {}

func (x *EditPlatformNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditPlatformNodeReply.ProtoReflect.Descriptor instead.
func (*EditPlatformNodeReply) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{6}
}

// 获取节点列表
// -------------------------------------------
type GetPlatformNodeListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlatformNodeListReq) Reset() {
	*x = GetPlatformNodeListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlatformNodeListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlatformNodeListReq) ProtoMessage() {}

func (x *GetPlatformNodeListReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlatformNodeListReq.ProtoReflect.Descriptor instead.
func (*GetPlatformNodeListReq) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{7}
}

type GetListPlatformNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*PlatformNodeInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"` //节点详情
}

func (x *GetListPlatformNodeReply) Reset() {
	*x = GetListPlatformNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListPlatformNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListPlatformNodeReply) ProtoMessage() {}

func (x *GetListPlatformNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListPlatformNodeReply.ProtoReflect.Descriptor instead.
func (*GetListPlatformNodeReply) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{8}
}

func (x *GetListPlatformNodeReply) GetList() []*PlatformNodeInfo {
	if x != nil {
		return x.List
	}
	return nil
}

// 获取所有节点列表
// -------------------------------------------
type GetAllAddressReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllAddressReq) Reset() {
	*x = GetAllAddressReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAddressReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAddressReq) ProtoMessage() {}

func (x *GetAllAddressReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAddressReq.ProtoReflect.Descriptor instead.
func (*GetAllAddressReq) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{9}
}

type GetAllAddressReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"` //节点详情
}

func (x *GetAllAddressReply) Reset() {
	*x = GetAllAddressReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_node_v2_platform_node_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAddressReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAddressReply) ProtoMessage() {}

func (x *GetAllAddressReply) ProtoReflect() protoreflect.Message {
	mi := &file_platform_node_v2_platform_node_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAddressReply.ProtoReflect.Descriptor instead.
func (*GetAllAddressReply) Descriptor() ([]byte, []int) {
	return file_platform_node_v2_platform_node_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllAddressReply) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

var File_platform_node_v2_platform_node_proto protoreflect.FileDescriptor

var file_platform_node_v2_platform_node_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x54, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x46, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x2e, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22,
	0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x78, 0x0a,
	0x13, 0x45, 0x64, 0x69, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x45, 0x64, 0x69, 0x74, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e,
	0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x56, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x32, 0xf4, 0x03, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x5b, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5b,
	0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a, 0x04, 0x45,
	0x64, 0x69, 0x74, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x67, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x61, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x51, 0x0a, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x32, 0x50,
	0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_platform_node_v2_platform_node_proto_rawDescOnce sync.Once
	file_platform_node_v2_platform_node_proto_rawDescData = file_platform_node_v2_platform_node_proto_rawDesc
)

func file_platform_node_v2_platform_node_proto_rawDescGZIP() []byte {
	file_platform_node_v2_platform_node_proto_rawDescOnce.Do(func() {
		file_platform_node_v2_platform_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_node_v2_platform_node_proto_rawDescData)
	})
	return file_platform_node_v2_platform_node_proto_rawDescData
}

var file_platform_node_v2_platform_node_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_platform_node_v2_platform_node_proto_goTypes = []interface{}{
	(*PlatformNodeInfo)(nil),         // 0: api.platform.node.v2.PlatformNodeInfo
	(*AddPlatformNodeReq)(nil),       // 1: api.platform.node.v2.AddPlatformNodeReq
	(*AddPlatformNodeReply)(nil),     // 2: api.platform.node.v2.AddPlatformNodeReply
	(*DelPlatformNodeReq)(nil),       // 3: api.platform.node.v2.DelPlatformNodeReq
	(*DelPlatformNodeReply)(nil),     // 4: api.platform.node.v2.DelPlatformNodeReply
	(*EditPlatformNodeReq)(nil),      // 5: api.platform.node.v2.EditPlatformNodeReq
	(*EditPlatformNodeReply)(nil),    // 6: api.platform.node.v2.EditPlatformNodeReply
	(*GetPlatformNodeListReq)(nil),   // 7: api.platform.node.v2.GetPlatformNodeListReq
	(*GetListPlatformNodeReply)(nil), // 8: api.platform.node.v2.GetListPlatformNodeReply
	(*GetAllAddressReq)(nil),         // 9: api.platform.node.v2.GetAllAddressReq
	(*GetAllAddressReply)(nil),       // 10: api.platform.node.v2.GetAllAddressReply
}
var file_platform_node_v2_platform_node_proto_depIdxs = []int32{
	0,  // 0: api.platform.node.v2.GetListPlatformNodeReply.list:type_name -> api.platform.node.v2.PlatformNodeInfo
	1,  // 1: api.platform.node.v2.PlatformNode.Add:input_type -> api.platform.node.v2.AddPlatformNodeReq
	3,  // 2: api.platform.node.v2.PlatformNode.Del:input_type -> api.platform.node.v2.DelPlatformNodeReq
	5,  // 3: api.platform.node.v2.PlatformNode.Edit:input_type -> api.platform.node.v2.EditPlatformNodeReq
	7,  // 4: api.platform.node.v2.PlatformNode.GetList:input_type -> api.platform.node.v2.GetPlatformNodeListReq
	9,  // 5: api.platform.node.v2.PlatformNode.GetAllAddress:input_type -> api.platform.node.v2.GetAllAddressReq
	2,  // 6: api.platform.node.v2.PlatformNode.Add:output_type -> api.platform.node.v2.AddPlatformNodeReply
	4,  // 7: api.platform.node.v2.PlatformNode.Del:output_type -> api.platform.node.v2.DelPlatformNodeReply
	6,  // 8: api.platform.node.v2.PlatformNode.Edit:output_type -> api.platform.node.v2.EditPlatformNodeReply
	8,  // 9: api.platform.node.v2.PlatformNode.GetList:output_type -> api.platform.node.v2.GetListPlatformNodeReply
	10, // 10: api.platform.node.v2.PlatformNode.GetAllAddress:output_type -> api.platform.node.v2.GetAllAddressReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_platform_node_v2_platform_node_proto_init() }
func file_platform_node_v2_platform_node_proto_init() {
	if File_platform_node_v2_platform_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_node_v2_platform_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformNodeInfo); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPlatformNodeReq); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPlatformNodeReply); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelPlatformNodeReq); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelPlatformNodeReply); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditPlatformNodeReq); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditPlatformNodeReply); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlatformNodeListReq); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListPlatformNodeReply); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAddressReq); i {
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
		file_platform_node_v2_platform_node_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAddressReply); i {
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
	file_platform_node_v2_platform_node_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_platform_node_v2_platform_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_platform_node_v2_platform_node_proto_goTypes,
		DependencyIndexes: file_platform_node_v2_platform_node_proto_depIdxs,
		MessageInfos:      file_platform_node_v2_platform_node_proto_msgTypes,
	}.Build()
	File_platform_node_v2_platform_node_proto = out.File
	file_platform_node_v2_platform_node_proto_rawDesc = nil
	file_platform_node_v2_platform_node_proto_goTypes = nil
	file_platform_node_v2_platform_node_proto_depIdxs = nil
}
