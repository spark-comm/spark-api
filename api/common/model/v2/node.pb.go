// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: common/model/v2/node.proto

package v2

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// 节点基本信息
type NodeBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId        string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`                // profile id 和 userId 等价
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                    // 名称
	Code          string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`                    // 节点邀请码
	StartTime     int64  `protobuf:"varint,4,opt,name=startTime,proto3" json:"startTime,omitempty"`         // 开始时间
	EndTime       int64  `protobuf:"varint,5,opt,name=endTime,proto3" json:"endTime,omitempty"`             // 结束时间
	MasterApi     string `protobuf:"bytes,6,opt,name=masterApi,proto3" json:"masterApi,omitempty"`          // 管理节点地址
	SlaveApi      string `protobuf:"bytes,7,opt,name=slaveApi,proto3" json:"slaveApi,omitempty"`            // 应用节点接口地址
	ManageApi     string `protobuf:"bytes,8,opt,name=manageApi,proto3" json:"manageApi,omitempty"`          // 管理节点地址
	ImApiUrl      string `protobuf:"bytes,9,opt,name=imApiUrl,proto3" json:"imApiUrl,omitempty"`            // 客户imapi域名地址,兼容2.0
	ImWsUrl       string `protobuf:"bytes,10,opt,name=imWsUrl,proto3" json:"imWsUrl,omitempty"`             // im ws api 地址
	Version       string `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`             // 版本
	AvAppid       string `protobuf:"bytes,12,opt,name=avAppid,proto3" json:"avAppid,omitempty"`             // 腾讯音视频app id
	AvSecret      string `protobuf:"bytes,13,opt,name=avSecret,proto3" json:"avSecret,omitempty"`           // 腾讯音视频密钥
	Logo          string `protobuf:"bytes,14,opt,name=logo,proto3" json:"logo,omitempty"`                   // logo
	NodeType      int32  `protobuf:"varint,15,opt,name=nodeType,proto3" json:"nodeType,omitempty"`          //节点类型 0 独立节点 1 代理节点
	PromotionCode string `protobuf:"bytes,16,opt,name=promotionCode,proto3" json:"promotionCode,omitempty"` // 代理者推广码
	CreatedAt     int64  `protobuf:"varint,17,opt,name=createdAt,proto3" json:"createdAt,omitempty"`        // 创建时间
	UpdatedAt     int64  `protobuf:"varint,18,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`        // 更新时间
	DeletedAt     int64  `protobuf:"varint,19,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`        // 删除时间
	Status        int32  `protobuf:"varint,20,opt,name=status,proto3" json:"status,omitempty"`              // 1：启用 2：停用 3: 时间限制
	Sort          int32  `protobuf:"varint,21,opt,name=sort,proto3" json:"sort,omitempty"`                  //排序
	MasterNodeId  string `protobuf:"bytes,22,opt,name=masterNodeId,proto3" json:"masterNodeId,omitempty"`   // 主几点id// 当为租户节点时
}

func (x *NodeBase) Reset() {
	*x = NodeBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeBase) ProtoMessage() {}

func (x *NodeBase) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeBase.ProtoReflect.Descriptor instead.
func (*NodeBase) Descriptor() ([]byte, []int) {
	return file_common_model_v2_node_proto_rawDescGZIP(), []int{0}
}

func (x *NodeBase) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeBase) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *NodeBase) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *NodeBase) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *NodeBase) GetMasterApi() string {
	if x != nil {
		return x.MasterApi
	}
	return ""
}

func (x *NodeBase) GetSlaveApi() string {
	if x != nil {
		return x.SlaveApi
	}
	return ""
}

func (x *NodeBase) GetManageApi() string {
	if x != nil {
		return x.ManageApi
	}
	return ""
}

func (x *NodeBase) GetImApiUrl() string {
	if x != nil {
		return x.ImApiUrl
	}
	return ""
}

func (x *NodeBase) GetImWsUrl() string {
	if x != nil {
		return x.ImWsUrl
	}
	return ""
}

func (x *NodeBase) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NodeBase) GetAvAppid() string {
	if x != nil {
		return x.AvAppid
	}
	return ""
}

func (x *NodeBase) GetAvSecret() string {
	if x != nil {
		return x.AvSecret
	}
	return ""
}

func (x *NodeBase) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *NodeBase) GetNodeType() int32 {
	if x != nil {
		return x.NodeType
	}
	return 0
}

func (x *NodeBase) GetPromotionCode() string {
	if x != nil {
		return x.PromotionCode
	}
	return ""
}

func (x *NodeBase) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NodeBase) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *NodeBase) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *NodeBase) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *NodeBase) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *NodeBase) GetMasterNodeId() string {
	if x != nil {
		return x.MasterNodeId
	}
	return ""
}

// 租户下发信息信息
type TenantDistributeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                             //租户主键
	CreatedAt        int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`              //创建时间
	UpdatedAt        int64  `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`              //创建时间
	DeletedAt        int64  `protobuf:"varint,4,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`              //删除时间
	AssociatedUserID string `protobuf:"bytes,5,opt,name=associatedUserID,proto3" json:"associatedUserID,omitempty"` //开通用户id
	Name             string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`                         //租户名称
	Logo             string `protobuf:"bytes,7,opt,name=logo,proto3" json:"logo,omitempty"`                         //租户logo
	Introduce        string `protobuf:"bytes,8,opt,name=introduce,proto3" json:"introduce,omitempty"`               //简介
	WebsiteAddress   string `protobuf:"bytes,9,opt,name=websiteAddress,proto3" json:"websiteAddress,omitempty"`     //网站地址
	WebsiteStatus    int32  `protobuf:"varint,10,opt,name=websiteStatus,proto3" json:"websiteStatus,omitempty"`     //网站状态1，申请中，0未开通，2:可用
	PackageType      string `protobuf:"bytes,11,opt,name=packageType,proto3" json:"packageType,omitempty"`          //套餐类型
	MaxNumber        int64  `protobuf:"varint,12,opt,name=maxNumber,proto3" json:"maxNumber,omitempty"`             // 最大用户数
	MaxGroupNumber   int64  `protobuf:"varint,13,opt,name=maxGroupNumber,proto3" json:"maxGroupNumber,omitempty"`   //最大群人数
	Phone            string `protobuf:"bytes,14,opt,name=phone,proto3" json:"phone,omitempty"`                      //手机号
	Status           int32  `protobuf:"varint,15,opt,name=status,proto3" json:"status,omitempty"`                   //状态，1，可用，0未开通，2过期
	StartTime        int64  `protobuf:"varint,16,opt,name=startTime,proto3" json:"startTime,omitempty"`             //开始时间
	EndTime          int64  `protobuf:"varint,17,opt,name=endTime,proto3" json:"endTime,omitempty"`                 //结束时间
	Code             string `protobuf:"bytes,18,opt,name=code,proto3" json:"code,omitempty"`                        //租户code
	PubKey           string `protobuf:"bytes,19,opt,name=pubKey,proto3" json:"pubKey,omitempty"`                    //租户公钥
	PriKey           string `protobuf:"bytes,20,opt,name=priKey,proto3" json:"priKey,omitempty"`                    // 租户私钥
}

func (x *TenantDistributeInfo) Reset() {
	*x = TenantDistributeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantDistributeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantDistributeInfo) ProtoMessage() {}

func (x *TenantDistributeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantDistributeInfo.ProtoReflect.Descriptor instead.
func (*TenantDistributeInfo) Descriptor() ([]byte, []int) {
	return file_common_model_v2_node_proto_rawDescGZIP(), []int{1}
}

func (x *TenantDistributeInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TenantDistributeInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TenantDistributeInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *TenantDistributeInfo) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *TenantDistributeInfo) GetAssociatedUserID() string {
	if x != nil {
		return x.AssociatedUserID
	}
	return ""
}

func (x *TenantDistributeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TenantDistributeInfo) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *TenantDistributeInfo) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

func (x *TenantDistributeInfo) GetWebsiteAddress() string {
	if x != nil {
		return x.WebsiteAddress
	}
	return ""
}

func (x *TenantDistributeInfo) GetWebsiteStatus() int32 {
	if x != nil {
		return x.WebsiteStatus
	}
	return 0
}

func (x *TenantDistributeInfo) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

func (x *TenantDistributeInfo) GetMaxNumber() int64 {
	if x != nil {
		return x.MaxNumber
	}
	return 0
}

func (x *TenantDistributeInfo) GetMaxGroupNumber() int64 {
	if x != nil {
		return x.MaxGroupNumber
	}
	return 0
}

func (x *TenantDistributeInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *TenantDistributeInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TenantDistributeInfo) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TenantDistributeInfo) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *TenantDistributeInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TenantDistributeInfo) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *TenantDistributeInfo) GetPriKey() string {
	if x != nil {
		return x.PriKey
	}
	return ""
}

// 租户上报信息信息
type TenantEscalationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                               //租户主键
	CreatedAt         int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`                //创建时间
	UpdatedAt         int64  `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`                //创建时间
	DeletedAt         int64  `protobuf:"varint,4,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`                //删除时间
	AssociatedUsersID string `protobuf:"bytes,5,opt,name=associatedUsersID,proto3" json:"associatedUsersID,omitempty"` //开通用户id
	Name              string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`                           //租户名称
	Logo              string `protobuf:"bytes,7,opt,name=logo,proto3" json:"logo,omitempty"`                           //租户logo
	Introduce         string `protobuf:"bytes,8,opt,name=introduce,proto3" json:"introduce,omitempty"`                 //简介
	WebsiteAddress    string `protobuf:"bytes,9,opt,name=websiteAddress,proto3" json:"websiteAddress,omitempty"`       //网站地址
	WebsiteStatus     int32  `protobuf:"varint,10,opt,name=websiteStatus,proto3" json:"websiteStatus,omitempty"`       //网站状态1，申请中，0未开通，2:可用
	MaxNumber         int64  `protobuf:"varint,12,opt,name=maxNumber,proto3" json:"maxNumber,omitempty"`               // 最大用户数
	MaxGroupNumber    int64  `protobuf:"varint,13,opt,name=maxGroupNumber,proto3" json:"maxGroupNumber,omitempty"`     //最大群人数
	Phone             string `protobuf:"bytes,14,opt,name=phone,proto3" json:"phone,omitempty"`                        //手机号
	Status            int32  `protobuf:"varint,15,opt,name=status,proto3" json:"status,omitempty"`                     //状态，1，可用，0未开通，2过期
	StartTime         int64  `protobuf:"varint,16,opt,name=startTime,proto3" json:"startTime,omitempty"`               //开始时间
	EndTime           int64  `protobuf:"varint,17,opt,name=endTime,proto3" json:"endTime,omitempty"`                   //结束时间
}

func (x *TenantEscalationInfo) Reset() {
	*x = TenantEscalationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantEscalationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantEscalationInfo) ProtoMessage() {}

func (x *TenantEscalationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantEscalationInfo.ProtoReflect.Descriptor instead.
func (*TenantEscalationInfo) Descriptor() ([]byte, []int) {
	return file_common_model_v2_node_proto_rawDescGZIP(), []int{2}
}

func (x *TenantEscalationInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TenantEscalationInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TenantEscalationInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *TenantEscalationInfo) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *TenantEscalationInfo) GetAssociatedUsersID() string {
	if x != nil {
		return x.AssociatedUsersID
	}
	return ""
}

func (x *TenantEscalationInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TenantEscalationInfo) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *TenantEscalationInfo) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

func (x *TenantEscalationInfo) GetWebsiteAddress() string {
	if x != nil {
		return x.WebsiteAddress
	}
	return ""
}

func (x *TenantEscalationInfo) GetWebsiteStatus() int32 {
	if x != nil {
		return x.WebsiteStatus
	}
	return 0
}

func (x *TenantEscalationInfo) GetMaxNumber() int64 {
	if x != nil {
		return x.MaxNumber
	}
	return 0
}

func (x *TenantEscalationInfo) GetMaxGroupNumber() int64 {
	if x != nil {
		return x.MaxGroupNumber
	}
	return 0
}

func (x *TenantEscalationInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *TenantEscalationInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TenantEscalationInfo) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TenantEscalationInfo) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

// 主节点信息同步
type NodeMasterSyncInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId       string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`             // 节点id
	Enable       string `protobuf:"bytes,2,opt,name=enable,proto3" json:"enable,omitempty"`             // 是否可用0 可用 1 不可用
	MasterApi    string `protobuf:"bytes,3,opt,name=masterApi,proto3" json:"masterApi,omitempty"`       // 主节点地址
	MasterPubKey string `protobuf:"bytes,4,opt,name=masterPubKey,proto3" json:"masterPubKey,omitempty"` // 主节点公钥匙
	Functions    string `protobuf:"bytes,5,opt,name=functions,proto3" json:"functions,omitempty"`       // 节点功能
	NodePriKey   string `protobuf:"bytes,6,opt,name=nodePriKey,proto3" json:"nodePriKey,omitempty"`     // 应用节点接口地址
}

func (x *NodeMasterSyncInfo) Reset() {
	*x = NodeMasterSyncInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMasterSyncInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMasterSyncInfo) ProtoMessage() {}

func (x *NodeMasterSyncInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMasterSyncInfo.ProtoReflect.Descriptor instead.
func (*NodeMasterSyncInfo) Descriptor() ([]byte, []int) {
	return file_common_model_v2_node_proto_rawDescGZIP(), []int{3}
}

func (x *NodeMasterSyncInfo) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeMasterSyncInfo) GetEnable() string {
	if x != nil {
		return x.Enable
	}
	return ""
}

func (x *NodeMasterSyncInfo) GetMasterApi() string {
	if x != nil {
		return x.MasterApi
	}
	return ""
}

func (x *NodeMasterSyncInfo) GetMasterPubKey() string {
	if x != nil {
		return x.MasterPubKey
	}
	return ""
}

func (x *NodeMasterSyncInfo) GetFunctions() string {
	if x != nil {
		return x.Functions
	}
	return ""
}

func (x *NodeMasterSyncInfo) GetNodePriKey() string {
	if x != nil {
		return x.NodePriKey
	}
	return ""
}

// 获取节点信息同步请求
type NodeMasterSyncInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"` // 节点id
}

func (x *NodeMasterSyncInfoReq) Reset() {
	*x = NodeMasterSyncInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMasterSyncInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMasterSyncInfoReq) ProtoMessage() {}

func (x *NodeMasterSyncInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMasterSyncInfoReq.ProtoReflect.Descriptor instead.
func (*NodeMasterSyncInfoReq) Descriptor() ([]byte, []int) {
	return file_common_model_v2_node_proto_rawDescGZIP(), []int{4}
}

func (x *NodeMasterSyncInfoReq) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

var File_common_model_v2_node_proto protoreflect.FileDescriptor

var file_common_model_v2_node_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x32, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x32, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x04, 0x0a, 0x08, 0x4e,
	0x6f, 0x64, 0x65, 0x42, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x69, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x69, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x41, 0x70, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x41, 0x70, 0x69, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x41, 0x70, 0x69,
	0x55, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x41, 0x70, 0x69,
	0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x57, 0x73, 0x55, 0x72, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x57, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x41, 0x70, 0x70,
	0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x76, 0x41, 0x70, 0x70, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x76, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x76, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0xd2, 0x04,
	0x0a, 0x14, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x69, 0x4b, 0x65, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x69, 0x4b,
	0x65, 0x79, 0x22, 0xee, 0x03, 0x0a, 0x14, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x45, 0x73, 0x63,
	0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d,
	0x61, 0x78, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x69, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x4b, 0x65, 0x79, 0x22, 0x2f, 0x0a, 0x15, 0x4e, 0x6f,
	0x64, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x42, 0x4f, 0x0a, 0x13, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x32, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_model_v2_node_proto_rawDescOnce sync.Once
	file_common_model_v2_node_proto_rawDescData = file_common_model_v2_node_proto_rawDesc
)

func file_common_model_v2_node_proto_rawDescGZIP() []byte {
	file_common_model_v2_node_proto_rawDescOnce.Do(func() {
		file_common_model_v2_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_model_v2_node_proto_rawDescData)
	})
	return file_common_model_v2_node_proto_rawDescData
}

var file_common_model_v2_node_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_model_v2_node_proto_goTypes = []interface{}{
	(*NodeBase)(nil),              // 0: api.common.model.v2.NodeBase
	(*TenantDistributeInfo)(nil),  // 1: api.common.model.v2.TenantDistributeInfo
	(*TenantEscalationInfo)(nil),  // 2: api.common.model.v2.TenantEscalationInfo
	(*NodeMasterSyncInfo)(nil),    // 3: api.common.model.v2.NodeMasterSyncInfo
	(*NodeMasterSyncInfoReq)(nil), // 4: api.common.model.v2.NodeMasterSyncInfoReq
}
var file_common_model_v2_node_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_model_v2_node_proto_init() }
func file_common_model_v2_node_proto_init() {
	if File_common_model_v2_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_model_v2_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeBase); i {
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
		file_common_model_v2_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantDistributeInfo); i {
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
		file_common_model_v2_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantEscalationInfo); i {
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
		file_common_model_v2_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMasterSyncInfo); i {
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
		file_common_model_v2_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMasterSyncInfoReq); i {
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
			RawDescriptor: file_common_model_v2_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_model_v2_node_proto_goTypes,
		DependencyIndexes: file_common_model_v2_node_proto_depIdxs,
		MessageInfos:      file_common_model_v2_node_proto_msgTypes,
	}.Build()
	File_common_model_v2_node_proto = out.File
	file_common_model_v2_node_proto_rawDesc = nil
	file_common_model_v2_node_proto_goTypes = nil
	file_common_model_v2_node_proto_depIdxs = nil
}
