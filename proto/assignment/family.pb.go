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
	Created              int64           `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64           `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string          `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string          `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string          `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string          `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Address              string          `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	Cover                string          `protobuf:"bytes,10,opt,name=cover,proto3" json:"cover,omitempty"`
	Region               string          `protobuf:"bytes,11,opt,name=region,proto3" json:"region,omitempty"`
	Status               uint32          `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	Location             string          `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	Passwords            string          `protobuf:"bytes,14,opt,name=passwords,proto3" json:"passwords,omitempty"`
	Master               string          `protobuf:"bytes,15,opt,name=master,proto3" json:"master,omitempty"`
	Assistants           []string        `protobuf:"bytes,26,rep,name=assistants,proto3" json:"assistants,omitempty"`
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

func (m *FamilyInfo) GetName() string {
	if m != nil {
		return m.Name
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

type ReqFamilyAdd struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string          `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string          `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Passwords            string          `protobuf:"bytes,4,opt,name=passwords,proto3" json:"passwords,omitempty"`
	Location             string          `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Address              string          `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Region               string          `protobuf:"bytes,7,opt,name=region,proto3" json:"region,omitempty"`
	Cover                string          `protobuf:"bytes,8,opt,name=cover,proto3" json:"cover,omitempty"`
	Master               string          `protobuf:"bytes,9,opt,name=master,proto3" json:"master,omitempty"`
	Members              []*IdentifyInfo `protobuf:"bytes,10,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReqFamilyAdd) Reset()         { *m = ReqFamilyAdd{} }
func (m *ReqFamilyAdd) String() string { return proto.CompactTextString(m) }
func (*ReqFamilyAdd) ProtoMessage()    {}
func (*ReqFamilyAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{1}
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

func (m *ReqFamilyAdd) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ReqFamilyAdd) GetMembers() []*IdentifyInfo {
	if m != nil {
		return m.Members
	}
	return nil
}

type ReqFamilyUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Passwords            string   `protobuf:"bytes,4,opt,name=passwords,proto3" json:"passwords,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFamilyUpdate) Reset()         { *m = ReqFamilyUpdate{} }
func (m *ReqFamilyUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqFamilyUpdate) ProtoMessage()    {}
func (*ReqFamilyUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{2}
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

func (m *ReqFamilyUpdate) GetPasswords() string {
	if m != nil {
		return m.Passwords
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
	return fileDescriptor_b5f82ff24828bc16, []int{3}
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
	Total                uint32        `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32        `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyFamilyList) Reset()         { *m = ReplyFamilyList{} }
func (m *ReplyFamilyList) String() string { return proto.CompactTextString(m) }
func (*ReplyFamilyList) ProtoMessage()    {}
func (*ReplyFamilyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{4}
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

func (m *ReplyFamilyList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyFamilyList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

type ReplyFamilyMembers struct {
	Status               *ReplyStatus    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*IdentifyInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyFamilyMembers) Reset()         { *m = ReplyFamilyMembers{} }
func (m *ReplyFamilyMembers) String() string { return proto.CompactTextString(m) }
func (*ReplyFamilyMembers) ProtoMessage()    {}
func (*ReplyFamilyMembers) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f82ff24828bc16, []int{5}
}

func (m *ReplyFamilyMembers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFamilyMembers.Unmarshal(m, b)
}
func (m *ReplyFamilyMembers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFamilyMembers.Marshal(b, m, deterministic)
}
func (m *ReplyFamilyMembers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFamilyMembers.Merge(m, src)
}
func (m *ReplyFamilyMembers) XXX_Size() int {
	return xxx_messageInfo_ReplyFamilyMembers.Size(m)
}
func (m *ReplyFamilyMembers) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFamilyMembers.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFamilyMembers proto.InternalMessageInfo

func (m *ReplyFamilyMembers) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyFamilyMembers) GetList() []*IdentifyInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*FamilyInfo)(nil), "omo.msp.assignment.FamilyInfo")
	proto.RegisterType((*ReqFamilyAdd)(nil), "omo.msp.assignment.ReqFamilyAdd")
	proto.RegisterType((*ReqFamilyUpdate)(nil), "omo.msp.assignment.ReqFamilyUpdate")
	proto.RegisterType((*ReplyFamilyInfo)(nil), "omo.msp.assignment.ReplyFamilyInfo")
	proto.RegisterType((*ReplyFamilyList)(nil), "omo.msp.assignment.ReplyFamilyList")
	proto.RegisterType((*ReplyFamilyMembers)(nil), "omo.msp.assignment.ReplyFamilyMembers")
}

func init() {
	proto.RegisterFile("proto/assignment/family.proto", fileDescriptor_b5f82ff24828bc16)
}

var fileDescriptor_b5f82ff24828bc16 = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4d, 0x6f, 0x13, 0x3b,
	0x14, 0xed, 0x64, 0xf2, 0xd1, 0xdc, 0x7c, 0xb4, 0xb2, 0x9e, 0x9e, 0xac, 0xa8, 0xed, 0x9b, 0x37,
	0x95, 0x50, 0x56, 0xa9, 0x14, 0x90, 0x90, 0xd8, 0xb5, 0x8b, 0x56, 0x95, 0xa8, 0x28, 0x13, 0x10,
	0x02, 0x56, 0xee, 0x8c, 0x13, 0x2c, 0x32, 0xe3, 0xa9, 0xed, 0x14, 0x75, 0xc3, 0x86, 0x15, 0x7f,
	0x05, 0xf1, 0x2b, 0xf8, 0x61, 0x08, 0xd9, 0x9e, 0x4c, 0x32, 0x69, 0x9a, 0x04, 0xca, 0x6e, 0xce,
	0xf5, 0xf5, 0xf1, 0xf1, 0x39, 0xd7, 0x51, 0x60, 0x3f, 0x15, 0x5c, 0xf1, 0x23, 0x22, 0x25, 0x1b,
	0x25, 0x31, 0x4d, 0xd4, 0xd1, 0x90, 0xc4, 0x6c, 0x7c, 0xdb, 0x33, 0x75, 0x84, 0x78, 0xcc, 0x7b,
	0xb1, 0x4c, 0x7b, 0xb3, 0x86, 0xce, 0xdd, 0x2d, 0x21, 0x8f, 0x63, 0x9e, 0xd8, 0x2d, 0xfe, 0x4f,
	0x17, 0xe0, 0xd4, 0x70, 0x9c, 0x27, 0x43, 0x8e, 0x76, 0xc1, 0x9d, 0xb0, 0x08, 0x3b, 0x9e, 0xd3,
	0xad, 0x07, 0xfa, 0x13, 0xb5, 0xa1, 0xc4, 0x22, 0x5c, 0xf2, 0x9c, 0x6e, 0x39, 0x28, 0xb1, 0x08,
	0x61, 0xa8, 0x85, 0x82, 0x12, 0x45, 0x23, 0xec, 0x7a, 0x4e, 0xd7, 0x0d, 0xa6, 0x50, 0xaf, 0x4c,
	0xd2, 0xc8, 0xac, 0x94, 0xed, 0x4a, 0x06, 0xf3, 0x3d, 0x5c, 0xe0, 0x8a, 0x61, 0x9e, 0x42, 0xd4,
	0x81, 0x6d, 0x9e, 0x52, 0x61, 0x96, 0xaa, 0x66, 0x29, 0xc7, 0x08, 0x41, 0x39, 0x21, 0x31, 0xc5,
	0x35, 0x53, 0x37, 0xdf, 0xe8, 0x5f, 0xa8, 0x0a, 0x1a, 0x13, 0xf1, 0x11, 0x6f, 0x9b, 0x6a, 0x86,
	0xf4, 0x09, 0x24, 0x8a, 0x04, 0x95, 0x12, 0xd7, 0xed, 0x09, 0x19, 0x44, 0xff, 0x40, 0x25, 0xe4,
	0x37, 0x54, 0x60, 0x30, 0x75, 0x0b, 0x2c, 0xcf, 0x88, 0xf1, 0x04, 0x37, 0xa6, 0x3c, 0x1a, 0xe9,
	0xba, 0x54, 0x44, 0x4d, 0x24, 0x6e, 0x7a, 0x4e, 0xb7, 0x15, 0x64, 0x48, 0xeb, 0x1c, 0xf3, 0x90,
	0x28, 0xbd, 0xa3, 0x65, 0x75, 0x4e, 0x31, 0xda, 0x83, 0x7a, 0x4a, 0xa4, 0xfc, 0xc4, 0x45, 0x24,
	0x71, 0xdb, 0x2c, 0xce, 0x0a, 0x9a, 0x31, 0x26, 0x52, 0x51, 0x81, 0x77, 0xec, 0x49, 0x16, 0xa1,
	0x03, 0x00, 0x9d, 0x89, 0x54, 0x24, 0x51, 0x12, 0x77, 0x3c, 0xb7, 0x5b, 0x0f, 0xe6, 0x2a, 0xfa,
	0xf6, 0x8a, 0x8c, 0x24, 0xde, 0x33, 0x2b, 0xe6, 0x5b, 0x73, 0x91, 0x11, 0xd5, 0xfd, 0xfb, 0xa6,
	0x9a, 0x21, 0xf4, 0x0c, 0x6a, 0x31, 0x8d, 0xaf, 0xa8, 0x90, 0xf8, 0xc0, 0x73, 0xbb, 0x8d, 0xbe,
	0xd7, 0xbb, 0x3b, 0x09, 0xbd, 0xf3, 0x88, 0x26, 0x8a, 0x0d, 0x4d, 0xd0, 0xc1, 0x74, 0x83, 0xff,
	0xbd, 0x04, 0xcd, 0x80, 0x5e, 0xdb, 0x19, 0x38, 0x8e, 0xa2, 0xdc, 0x76, 0x67, 0xa9, 0xed, 0xa5,
	0x82, 0xed, 0xf3, 0xf1, 0xb9, 0x0b, 0xf1, 0x15, 0x6c, 0x29, 0x2f, 0xda, 0x32, 0x6f, 0x68, 0x65,
	0xc1, 0xd0, 0xb9, 0x30, 0xab, 0xc5, 0x30, 0x67, 0xb1, 0xd5, 0x0a, 0xb1, 0xe5, 0x21, 0x6f, 0x2f,
	0x84, 0x9c, 0x59, 0x5f, 0x2f, 0x58, 0x3f, 0x67, 0x17, 0xfc, 0xae, 0x5d, 0x5f, 0x1d, 0xd8, 0xc9,
	0xed, 0x7a, 0x6d, 0xe6, 0x7b, 0xc9, 0xa3, 0x99, 0x7a, 0x58, 0x5a, 0xea, 0xa1, 0x5b, 0xf0, 0x70,
	0xad, 0x4f, 0xb9, 0xc3, 0x95, 0xa2, 0xc3, 0xfe, 0x67, 0x2d, 0x25, 0x1d, 0xdf, 0xce, 0xbd, 0xdf,
	0xa7, 0xf9, 0xfc, 0x6a, 0x35, 0x8d, 0xfe, 0x7f, 0xcb, 0x6e, 0x66, 0x36, 0x0d, 0x4c, 0x5b, 0x3e,
	0xe0, 0x7d, 0x28, 0xb3, 0x64, 0xc8, 0x8d, 0xe2, 0x46, 0xff, 0x60, 0xd9, 0xb6, 0xd9, 0x31, 0x81,
	0xe9, 0xf5, 0xbf, 0x39, 0x05, 0x01, 0xcf, 0x99, 0x54, 0x0f, 0x12, 0x30, 0x66, 0x52, 0xe1, 0x92,
	0x49, 0x64, 0xad, 0x00, 0xdd, 0xab, 0x63, 0x57, 0x5c, 0x91, 0xb1, 0x71, 0xb4, 0x15, 0x58, 0xa0,
	0xab, 0x29, 0x19, 0x51, 0x6b, 0x66, 0x2b, 0xb0, 0xc0, 0xff, 0xe2, 0x00, 0x9a, 0x13, 0x7b, 0x61,
	0xf3, 0xfc, 0x73, 0xbd, 0x4f, 0x0a, 0x7a, 0xd7, 0x4f, 0x90, 0xe9, 0xee, 0xff, 0xa8, 0x41, 0xcb,
	0x0a, 0x18, 0x50, 0x71, 0xc3, 0x42, 0x8a, 0x5e, 0x42, 0xf5, 0x38, 0x8a, 0x5e, 0x24, 0x14, 0x79,
	0xcb, 0x8f, 0x9e, 0x3d, 0xcd, 0xce, 0xe1, 0xbd, 0xe2, 0x66, 0xd6, 0xf8, 0x5b, 0x28, 0x00, 0xb0,
	0x93, 0x79, 0x42, 0x24, 0x45, 0x87, 0x2b, 0x69, 0x6d, 0x63, 0x67, 0xff, 0x5e, 0xe6, 0x8c, 0xf3,
	0x02, 0xea, 0x01, 0x8d, 0xf9, 0x0d, 0xd5, 0x4a, 0xef, 0x31, 0xe9, 0x7a, 0x42, 0xa5, 0xd2, 0xfd,
	0xeb, 0xe9, 0x2e, 0xa1, 0x7a, 0x46, 0xd5, 0x46, 0x5c, 0x1b, 0x5e, 0xfa, 0x12, 0xaa, 0x03, 0x4a,
	0x44, 0xf8, 0xe1, 0xe1, 0x8c, 0x7a, 0x90, 0xfd, 0x2d, 0xf4, 0x1e, 0x76, 0xce, 0xa8, 0xd2, 0xe0,
	0xe4, 0xf6, 0x94, 0x8d, 0xf5, 0x2f, 0xc7, 0xff, 0x2b, 0xa8, 0x6d, 0xcb, 0xa6, 0xe4, 0x6f, 0xa0,
	0x79, 0x46, 0x95, 0x9e, 0x29, 0x26, 0x15, 0x0b, 0x37, 0x61, 0xf6, 0x57, 0x8e, 0xa6, 0xa1, 0xf1,
	0xb7, 0xd0, 0x2b, 0x68, 0x67, 0xe1, 0x6f, 0x22, 0x7a, 0xd3, 0xf8, 0x07, 0xd0, 0xb4, 0xad, 0xf6,
	0x15, 0x20, 0x7f, 0xa5, 0xc7, 0xea, 0x74, 0x4c, 0x46, 0xeb, 0x49, 0xdf, 0x42, 0xf3, 0x38, 0x4d,
	0x69, 0x12, 0xd9, 0xc7, 0xb8, 0x3e, 0xb8, 0x47, 0x6b, 0xbc, 0xcd, 0x1e, 0xb5, 0xc9, 0xae, 0x3d,
	0x98, 0x5c, 0x29, 0x41, 0x42, 0xf5, 0xd7, 0xc9, 0x4f, 0xd0, 0xbb, 0xdd, 0xc5, 0x3f, 0x55, 0x57,
	0x55, 0x53, 0x79, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x35, 0x45, 0xae, 0x93, 0xa2, 0x09, 0x00,
	0x00,
}
