// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/assignment/family.proto

package assignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FamilyInfo struct {
	Uid                  string          `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              int64           `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64           `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string          `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string          `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Sn                   string          `protobuf:"bytes,8,opt,name=sn,proto3" json:"sn,omitempty"`
	Remark               string          `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Address              string          `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	Cover                string          `protobuf:"bytes,11,opt,name=cover,proto3" json:"cover,omitempty"`
	Region               string          `protobuf:"bytes,12,opt,name=region,proto3" json:"region,omitempty"`
	Status               uint32          `protobuf:"varint,13,opt,name=status,proto3" json:"status,omitempty"`
	Location             string          `protobuf:"bytes,14,opt,name=location,proto3" json:"location,omitempty"`
	Passwords            string          `protobuf:"bytes,15,opt,name=passwords,proto3" json:"passwords,omitempty"`
	Master               string          `protobuf:"bytes,16,opt,name=master,proto3" json:"master,omitempty"`
	Assistants           []string        `protobuf:"bytes,26,rep,name=assistants,proto3" json:"assistants,omitempty"`
	Children             []string        `protobuf:"bytes,27,rep,name=children,proto3" json:"children,omitempty"`
	Tags                 []string        `protobuf:"bytes,28,rep,name=tags,proto3" json:"tags,omitempty"`
	Agents               []string        `protobuf:"bytes,29,rep,name=agents,proto3" json:"agents,omitempty"`
	Members              []*IdentifyInfo `protobuf:"bytes,30,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FamilyInfo) Reset()         { *m = FamilyInfo{} }
func (m *FamilyInfo) String() string { return proto.CompactTextString(m) }
func (*FamilyInfo) ProtoMessage()    {}
func (*FamilyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{0}
}

func (m *FamilyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FamilyInfo.Unmarshal(m, b)
}
func (m *FamilyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FamilyInfo.Marshal(b, m, deterministic)
}
func (m *FamilyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FamilyInfo.Merge(m, src)
}
func (m *FamilyInfo) XXX_Size() int {
	return xxx_messageInfo_FamilyInfo.Size(m)
}
func (m *FamilyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FamilyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FamilyInfo proto.InternalMessageInfo

func (m *FamilyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *FamilyInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FamilyInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FamilyInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *FamilyInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *FamilyInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *FamilyInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *FamilyInfo) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *FamilyInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *FamilyInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *FamilyInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *FamilyInfo) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *FamilyInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *FamilyInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *FamilyInfo) GetPasswords() string {
	if m != nil {
		return m.Passwords
	}
	return ""
}

func (m *FamilyInfo) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *FamilyInfo) GetAssistants() []string {
	if m != nil {
		return m.Assistants
	}
	return nil
}

func (m *FamilyInfo) GetChildren() []string {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *FamilyInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FamilyInfo) GetAgents() []string {
	if m != nil {
		return m.Agents
	}
	return nil
}

func (m *FamilyInfo) GetMembers() []*IdentifyInfo {
	if m != nil {
		return m.Members
	}
	return nil
}

type IdentifyInfo struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentifyInfo) Reset()         { *m = IdentifyInfo{} }
func (m *IdentifyInfo) String() string { return proto.CompactTextString(m) }
func (*IdentifyInfo) ProtoMessage()    {}
func (*IdentifyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{1}
}

func (m *IdentifyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentifyInfo.Unmarshal(m, b)
}
func (m *IdentifyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentifyInfo.Marshal(b, m, deterministic)
}
func (m *IdentifyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifyInfo.Merge(m, src)
}
func (m *IdentifyInfo) XXX_Size() int {
	return xxx_messageInfo_IdentifyInfo.Size(m)
}
func (m *IdentifyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifyInfo proto.InternalMessageInfo

func (m *IdentifyInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *IdentifyInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type ReqFamilyAdd struct {
	Sn                   string   `protobuf:"bytes,1,opt,name=sn,proto3" json:"sn,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Passwords            string   `protobuf:"bytes,5,opt,name=passwords,proto3" json:"passwords,omitempty"`
	Location             string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Address              string   `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Region               string   `protobuf:"bytes,8,opt,name=region,proto3" json:"region,omitempty"`
	Cover                string   `protobuf:"bytes,9,opt,name=cover,proto3" json:"cover,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFamilyAdd) Reset()         { *m = ReqFamilyAdd{} }
func (m *ReqFamilyAdd) String() string { return proto.CompactTextString(m) }
func (*ReqFamilyAdd) ProtoMessage()    {}
func (*ReqFamilyAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{2}
}

func (m *ReqFamilyAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFamilyAdd.Unmarshal(m, b)
}
func (m *ReqFamilyAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFamilyAdd.Marshal(b, m, deterministic)
}
func (m *ReqFamilyAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFamilyAdd.Merge(m, src)
}
func (m *ReqFamilyAdd) XXX_Size() int {
	return xxx_messageInfo_ReqFamilyAdd.Size(m)
}
func (m *ReqFamilyAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFamilyAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFamilyAdd proto.InternalMessageInfo

func (m *ReqFamilyAdd) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ReqFamilyAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqFamilyAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqFamilyAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFamilyAdd) GetPasswords() string {
	if m != nil {
		return m.Passwords
	}
	return ""
}

func (m *ReqFamilyAdd) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *ReqFamilyAdd) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ReqFamilyAdd) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *ReqFamilyAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

type ReqFamilyUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFamilyUpdate) Reset()         { *m = ReqFamilyUpdate{} }
func (m *ReqFamilyUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqFamilyUpdate) ProtoMessage()    {}
func (*ReqFamilyUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{3}
}

func (m *ReqFamilyUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFamilyUpdate.Unmarshal(m, b)
}
func (m *ReqFamilyUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFamilyUpdate.Marshal(b, m, deterministic)
}
func (m *ReqFamilyUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFamilyUpdate.Merge(m, src)
}
func (m *ReqFamilyUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqFamilyUpdate.Size(m)
}
func (m *ReqFamilyUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFamilyUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFamilyUpdate proto.InternalMessageInfo

func (m *ReqFamilyUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqFamilyUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqFamilyUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqFamilyUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyFamilyInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *FamilyInfo  `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyFamilyInfo) Reset()         { *m = ReplyFamilyInfo{} }
func (m *ReplyFamilyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyFamilyInfo) ProtoMessage()    {}
func (*ReplyFamilyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{4}
}

func (m *ReplyFamilyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFamilyInfo.Unmarshal(m, b)
}
func (m *ReplyFamilyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFamilyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyFamilyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFamilyInfo.Merge(m, src)
}
func (m *ReplyFamilyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyFamilyInfo.Size(m)
}
func (m *ReplyFamilyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFamilyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFamilyInfo proto.InternalMessageInfo

func (m *ReplyFamilyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyFamilyInfo) GetInfo() *FamilyInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyFamilyList struct {
	Status               *ReplyStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*FamilyInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyFamilyList) Reset()         { *m = ReplyFamilyList{} }
func (m *ReplyFamilyList) String() string { return proto.CompactTextString(m) }
func (*ReplyFamilyList) ProtoMessage()    {}
func (*ReplyFamilyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{5}
}

func (m *ReplyFamilyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFamilyList.Unmarshal(m, b)
}
func (m *ReplyFamilyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFamilyList.Marshal(b, m, deterministic)
}
func (m *ReplyFamilyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFamilyList.Merge(m, src)
}
func (m *ReplyFamilyList) XXX_Size() int {
	return xxx_messageInfo_ReplyFamilyList.Size(m)
}
func (m *ReplyFamilyList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFamilyList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFamilyList proto.InternalMessageInfo

func (m *ReplyFamilyList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyFamilyList) GetList() []*FamilyInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*FamilyInfo)(nil), "omo.msp.assignment.FamilyInfo")
	proto.RegisterType((*IdentifyInfo)(nil), "omo.msp.assignment.IdentifyInfo")
	proto.RegisterType((*ReqFamilyAdd)(nil), "omo.msp.assignment.ReqFamilyAdd")
	proto.RegisterType((*ReqFamilyUpdate)(nil), "omo.msp.assignment.ReqFamilyUpdate")
	proto.RegisterType((*ReplyFamilyInfo)(nil), "omo.msp.assignment.ReplyFamilyInfo")
	proto.RegisterType((*ReplyFamilyList)(nil), "omo.msp.assignment.ReplyFamilyList")
}

func init() {
	proto.RegisterFile("proto/assignment/family.proto", fileDescriptor_b5f82ff24828bc16)
}

var fileDescriptor_b5f82ff24828bc16 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcd, 0x4e, 0x1b, 0x31,
	0x10, 0xc7, 0xc9, 0x07, 0x09, 0x99, 0x84, 0x80, 0xac, 0xaa, 0xb2, 0x52, 0xa0, 0xdb, 0xe5, 0x92,
	0x53, 0x90, 0xd2, 0x43, 0x25, 0x6e, 0x70, 0x00, 0x21, 0x15, 0xb5, 0xdd, 0xb4, 0x6a, 0xd5, 0x9e,
	0xcc, 0xda, 0x49, 0x2d, 0xb2, 0xf6, 0x62, 0x3b, 0x54, 0x5c, 0xfa, 0x40, 0x7d, 0x82, 0x3e, 0x52,
	0x1f, 0xa3, 0xb2, 0xbd, 0xd9, 0x6c, 0x42, 0x08, 0x91, 0xe0, 0xe6, 0xf9, 0xfa, 0xcf, 0xac, 0xe7,
	0xe7, 0x28, 0xb0, 0x9f, 0x2a, 0x69, 0xe4, 0x11, 0xd1, 0x9a, 0x8f, 0x44, 0xc2, 0x84, 0x39, 0x1a,
	0x92, 0x84, 0x8f, 0xef, 0x7a, 0xce, 0x8f, 0x90, 0x4c, 0x64, 0x2f, 0xd1, 0x69, 0x6f, 0x96, 0xd0,
	0xb9, 0x5f, 0x12, 0xcb, 0x24, 0x91, 0xc2, 0x97, 0x84, 0x7f, 0xaa, 0x00, 0x67, 0x4e, 0xe3, 0x42,
	0x0c, 0x25, 0xda, 0x85, 0xca, 0x84, 0x53, 0x5c, 0x0a, 0x4a, 0xdd, 0x46, 0x64, 0x8f, 0xa8, 0x0d,
	0x65, 0x4e, 0x71, 0x39, 0x28, 0x75, 0xab, 0x51, 0x99, 0x53, 0x84, 0xa0, 0x2a, 0x48, 0xc2, 0x70,
	0xc5, 0xa5, 0xb8, 0x33, 0xc2, 0x50, 0x8f, 0x15, 0x23, 0x86, 0x51, 0x5c, 0x0d, 0x4a, 0xdd, 0x4a,
	0x34, 0x35, 0x6d, 0x64, 0x92, 0x52, 0x17, 0xd9, 0xf4, 0x91, 0xcc, 0xcc, 0x6b, 0xa4, 0xc2, 0x35,
	0x27, 0x35, 0x35, 0x51, 0x07, 0xb6, 0x64, 0xca, 0x94, 0x0b, 0xd5, 0x5d, 0x28, 0xb7, 0xed, 0x34,
	0x5a, 0xe0, 0x2d, 0xe7, 0x2d, 0x6b, 0x81, 0x5e, 0x42, 0x4d, 0xb1, 0x84, 0xa8, 0x6b, 0xdc, 0x70,
	0xbe, 0xcc, 0xb2, 0xea, 0x84, 0x52, 0xc5, 0xb4, 0xc6, 0xe0, 0xd5, 0x33, 0x13, 0xbd, 0x80, 0xcd,
	0x58, 0xde, 0x32, 0x85, 0x9b, 0xce, 0xef, 0x0d, 0xaf, 0x33, 0xe2, 0x52, 0xe0, 0xd6, 0x54, 0xc7,
	0x5a, 0xd6, 0xaf, 0x0d, 0x31, 0x13, 0x8d, 0xb7, 0x83, 0x52, 0x77, 0x3b, 0xca, 0x2c, 0x3b, 0xe3,
	0x58, 0xc6, 0xc4, 0xd8, 0x8a, 0xb6, 0x9f, 0x71, 0x6a, 0xa3, 0x3d, 0x68, 0xa4, 0x44, 0xeb, 0x5f,
	0x52, 0x51, 0x8d, 0x77, 0x5c, 0x70, 0xe6, 0xb0, 0x8a, 0x09, 0xd1, 0x86, 0x29, 0xbc, 0xeb, 0x3b,
	0x79, 0x0b, 0x1d, 0x00, 0xd8, 0x1d, 0x69, 0x43, 0x84, 0xd1, 0xb8, 0x13, 0x54, 0xba, 0x8d, 0xa8,
	0xe0, 0xb1, 0x1d, 0xe3, 0x9f, 0x7c, 0x4c, 0x15, 0x13, 0xf8, 0x95, 0x8b, 0xe6, 0xb6, 0xdd, 0x89,
	0x21, 0x23, 0x8d, 0xf7, 0x9c, 0xdf, 0x9d, 0x6d, 0x1f, 0x32, 0x62, 0x56, 0x6b, 0xdf, 0x79, 0x33,
	0x0b, 0x1d, 0x43, 0x3d, 0x61, 0xc9, 0x15, 0x53, 0x1a, 0x1f, 0x04, 0x95, 0x6e, 0xb3, 0x1f, 0xf4,
	0xee, 0x53, 0xd3, 0xbb, 0xa0, 0x4c, 0x18, 0x3e, 0x74, 0x50, 0x44, 0xd3, 0x82, 0xf0, 0x18, 0x5a,
	0xc5, 0x80, 0xed, 0x3b, 0xd1, 0x4c, 0x65, 0xb8, 0xb8, 0x73, 0x61, 0x23, 0xe5, 0xe2, 0x46, 0xc2,
	0x7f, 0x25, 0x68, 0x45, 0xec, 0xc6, 0xb3, 0x76, 0x42, 0x69, 0xb6, 0xca, 0x52, 0xbe, 0xca, 0x29,
	0x58, 0xe5, 0x02, 0x58, 0x33, 0xb1, 0xca, 0xdc, 0x7a, 0x8b, 0x88, 0x54, 0x17, 0x10, 0x99, 0xbb,
	0xfe, 0xcd, 0xc5, 0xeb, 0x2f, 0x2e, 0xae, 0xb6, 0xb0, 0xb8, 0x02, 0x34, 0xf5, 0x79, 0x68, 0x66,
	0x78, 0x6c, 0xcd, 0xe1, 0x91, 0xc3, 0xd4, 0x28, 0xc0, 0x14, 0x5e, 0xc3, 0x4e, 0xfe, 0xa5, 0x5f,
	0x1c, 0xee, 0x4b, 0xde, 0xd5, 0x33, 0x7d, 0x6e, 0xf8, 0xdb, 0x36, 0x4b, 0xc7, 0x77, 0x85, 0x47,
	0xfc, 0x2e, 0x87, 0xd6, 0xf6, 0x6b, 0xf6, 0x5f, 0x2f, 0xdb, 0xb0, 0x2b, 0x1a, 0xb8, 0xb4, 0x9c,
	0xea, 0x3e, 0x54, 0xb9, 0x18, 0x4a, 0x37, 0x53, 0xb3, 0x7f, 0xb0, 0xac, 0x6c, 0xd6, 0x26, 0x72,
	0xb9, 0x0b, 0xfd, 0xdf, 0x73, 0x6d, 0x9e, 0xd4, 0x7f, 0xcc, 0xb5, 0xc1, 0x65, 0x07, 0xe6, 0xa3,
	0xfd, 0x6d, 0x6e, 0xff, 0x6f, 0x1d, 0xb6, 0xbd, 0x73, 0xc0, 0xd4, 0x2d, 0x8f, 0x19, 0xfa, 0x04,
	0xb5, 0x13, 0x4a, 0x3f, 0x08, 0x86, 0x82, 0xe5, 0x8d, 0x67, 0x10, 0x76, 0x0e, 0x1f, 0x1c, 0x6d,
	0xd6, 0x28, 0xdc, 0x40, 0x11, 0x80, 0x5f, 0xe4, 0x29, 0xd1, 0x0c, 0x1d, 0xae, 0x94, 0xf5, 0x89,
	0x9d, 0xfd, 0x07, 0x95, 0x33, 0xcd, 0x4b, 0x68, 0x44, 0x2c, 0x91, 0xb7, 0xcc, 0x4e, 0xfa, 0xc0,
	0x15, 0xdd, 0x4c, 0x98, 0x36, 0x36, 0xff, 0x71, 0xb9, 0x8f, 0x50, 0x3b, 0x67, 0x66, 0x2d, 0xad,
	0x35, 0x3f, 0xfa, 0x1b, 0xb4, 0xbd, 0xe2, 0xe9, 0xdd, 0x19, 0x1f, 0xdb, 0xdf, 0xa8, 0xe7, 0x52,
	0xfe, 0x01, 0x3b, 0xe7, 0xcc, 0x58, 0x56, 0x72, 0xe9, 0x37, 0x2b, 0xa4, 0x7d, 0xca, 0xa3, 0xe2,
	0x56, 0x2f, 0xdc, 0x40, 0x5f, 0xa1, 0x75, 0xce, 0x8c, 0x25, 0x8b, 0x6b, 0xc3, 0xe3, 0x75, 0x94,
	0xc3, 0x95, 0x80, 0x3a, 0x99, 0x70, 0x03, 0x7d, 0x86, 0x76, 0x06, 0xc1, 0x3a, 0x43, 0xaf, 0x8b,
	0xc1, 0x00, 0x5a, 0x3e, 0xd5, 0xbf, 0x05, 0x14, 0xae, 0xbc, 0x63, 0x73, 0x36, 0x26, 0xa3, 0x75,
	0x60, 0x68, 0x9d, 0xa4, 0x29, 0x13, 0xf4, 0xd2, 0xfd, 0x72, 0x3f, 0x05, 0xaf, 0xec, 0x56, 0x23,
	0x68, 0x0f, 0x26, 0x57, 0x46, 0x91, 0xd8, 0x3c, 0x97, 0xe6, 0x29, 0xfa, 0xbe, 0xbb, 0xf8, 0xe7,
	0xe4, 0xaa, 0xe6, 0x3c, 0x6f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xf2, 0x76, 0xf6, 0xea,
	0x08, 0x00, 0x00,
}