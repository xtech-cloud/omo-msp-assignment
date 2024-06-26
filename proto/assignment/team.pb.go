// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/assignment/team.proto

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

type TeamInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string   `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string   `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Master               string   `protobuf:"bytes,10,opt,name=master,proto3" json:"master,omitempty"`
	Region               string   `protobuf:"bytes,11,opt,name=region,proto3" json:"region,omitempty"`
	Status               uint32   `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	Limit                uint32   `protobuf:"varint,13,opt,name=limit,proto3" json:"limit,omitempty"`
	Parent               string   `protobuf:"bytes,14,opt,name=parent,proto3" json:"parent,omitempty"`
	Assistants           []string `protobuf:"bytes,15,rep,name=assistants,proto3" json:"assistants,omitempty"`
	Tags                 []string `protobuf:"bytes,16,rep,name=tags,proto3" json:"tags,omitempty"`
	Members              []string `protobuf:"bytes,17,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamInfo) Reset()         { *m = TeamInfo{} }
func (m *TeamInfo) String() string { return proto.CompactTextString(m) }
func (*TeamInfo) ProtoMessage()    {}
func (*TeamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_28bce25f9dd8d840, []int{0}
}

func (m *TeamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamInfo.Unmarshal(m, b)
}
func (m *TeamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamInfo.Marshal(b, m, deterministic)
}
func (m *TeamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamInfo.Merge(m, src)
}
func (m *TeamInfo) XXX_Size() int {
	return xxx_messageInfo_TeamInfo.Size(m)
}
func (m *TeamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeamInfo proto.InternalMessageInfo

func (m *TeamInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TeamInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TeamInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TeamInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TeamInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *TeamInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TeamInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *TeamInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TeamInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *TeamInfo) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *TeamInfo) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *TeamInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TeamInfo) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TeamInfo) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *TeamInfo) GetAssistants() []string {
	if m != nil {
		return m.Assistants
	}
	return nil
}

func (m *TeamInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TeamInfo) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type ReqTeamAdd struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Region               string   `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Master               string   `protobuf:"bytes,6,opt,name=master,proto3" json:"master,omitempty"`
	Parent               string   `protobuf:"bytes,7,opt,name=parent,proto3" json:"parent,omitempty"`
	Limit                uint32   `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`
	Tags                 []string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTeamAdd) Reset()         { *m = ReqTeamAdd{} }
func (m *ReqTeamAdd) String() string { return proto.CompactTextString(m) }
func (*ReqTeamAdd) ProtoMessage()    {}
func (*ReqTeamAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_28bce25f9dd8d840, []int{1}
}

func (m *ReqTeamAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTeamAdd.Unmarshal(m, b)
}
func (m *ReqTeamAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTeamAdd.Marshal(b, m, deterministic)
}
func (m *ReqTeamAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTeamAdd.Merge(m, src)
}
func (m *ReqTeamAdd) XXX_Size() int {
	return xxx_messageInfo_ReqTeamAdd.Size(m)
}
func (m *ReqTeamAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTeamAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTeamAdd proto.InternalMessageInfo

func (m *ReqTeamAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqTeamAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTeamAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqTeamAdd) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *ReqTeamAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTeamAdd) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ReqTeamAdd) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReqTeamAdd) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReqTeamAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqTeamUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTeamUpdate) Reset()         { *m = ReqTeamUpdate{} }
func (m *ReqTeamUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqTeamUpdate) ProtoMessage()    {}
func (*ReqTeamUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_28bce25f9dd8d840, []int{2}
}

func (m *ReqTeamUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTeamUpdate.Unmarshal(m, b)
}
func (m *ReqTeamUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTeamUpdate.Marshal(b, m, deterministic)
}
func (m *ReqTeamUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTeamUpdate.Merge(m, src)
}
func (m *ReqTeamUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqTeamUpdate.Size(m)
}
func (m *ReqTeamUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTeamUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTeamUpdate proto.InternalMessageInfo

func (m *ReqTeamUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqTeamUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTeamUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqTeamUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyTeamInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *TeamInfo    `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyTeamInfo) Reset()         { *m = ReplyTeamInfo{} }
func (m *ReplyTeamInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyTeamInfo) ProtoMessage()    {}
func (*ReplyTeamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_28bce25f9dd8d840, []int{3}
}

func (m *ReplyTeamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTeamInfo.Unmarshal(m, b)
}
func (m *ReplyTeamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTeamInfo.Marshal(b, m, deterministic)
}
func (m *ReplyTeamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTeamInfo.Merge(m, src)
}
func (m *ReplyTeamInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyTeamInfo.Size(m)
}
func (m *ReplyTeamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTeamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTeamInfo proto.InternalMessageInfo

func (m *ReplyTeamInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTeamInfo) GetInfo() *TeamInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyTeamList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*TeamInfo  `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	Total                uint32       `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32       `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyTeamList) Reset()         { *m = ReplyTeamList{} }
func (m *ReplyTeamList) String() string { return proto.CompactTextString(m) }
func (*ReplyTeamList) ProtoMessage()    {}
func (*ReplyTeamList) Descriptor() ([]byte, []int) {
	return fileDescriptor_28bce25f9dd8d840, []int{4}
}

func (m *ReplyTeamList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTeamList.Unmarshal(m, b)
}
func (m *ReplyTeamList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTeamList.Marshal(b, m, deterministic)
}
func (m *ReplyTeamList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTeamList.Merge(m, src)
}
func (m *ReplyTeamList) XXX_Size() int {
	return xxx_messageInfo_ReplyTeamList.Size(m)
}
func (m *ReplyTeamList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTeamList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTeamList proto.InternalMessageInfo

func (m *ReplyTeamList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTeamList) GetList() []*TeamInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyTeamList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyTeamList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func init() {
	proto.RegisterType((*TeamInfo)(nil), "omo.msp.assignment.TeamInfo")
	proto.RegisterType((*ReqTeamAdd)(nil), "omo.msp.assignment.ReqTeamAdd")
	proto.RegisterType((*ReqTeamUpdate)(nil), "omo.msp.assignment.ReqTeamUpdate")
	proto.RegisterType((*ReplyTeamInfo)(nil), "omo.msp.assignment.ReplyTeamInfo")
	proto.RegisterType((*ReplyTeamList)(nil), "omo.msp.assignment.ReplyTeamList")
}

func init() {
	proto.RegisterFile("proto/assignment/team.proto", fileDescriptor_28bce25f9dd8d840)
}

var fileDescriptor_28bce25f9dd8d840 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x4f, 0x1b, 0x3f,
	0x10, 0x65, 0xb3, 0x61, 0x93, 0x4c, 0xfe, 0xc0, 0xcf, 0xfa, 0xa9, 0xb2, 0xd2, 0x42, 0xc3, 0x9e,
	0x72, 0x0a, 0x15, 0x3d, 0xf4, 0x0c, 0x07, 0x50, 0xa5, 0x42, 0xd1, 0x86, 0xaa, 0x6a, 0x6f, 0x26,
	0x3b, 0xa4, 0x56, 0xe3, 0xf5, 0xb2, 0x76, 0xa8, 0xe8, 0x67, 0xe9, 0xa9, 0x87, 0x7e, 0xaa, 0x7e,
	0x98, 0xca, 0xf6, 0x6e, 0xb2, 0x04, 0x48, 0x52, 0xd1, 0xdb, 0xbe, 0x99, 0xe7, 0x37, 0x9e, 0x37,
	0xe3, 0x04, 0x9e, 0xa7, 0x99, 0xd4, 0x72, 0x9f, 0x29, 0xc5, 0xc7, 0x89, 0xc0, 0x44, 0xef, 0x6b,
	0x64, 0x62, 0x60, 0xa3, 0x84, 0x48, 0x21, 0x07, 0x42, 0xa5, 0x83, 0x79, 0xba, 0xbb, 0x73, 0xef,
	0xc0, 0x48, 0x0a, 0x21, 0x13, 0x77, 0x24, 0xfc, 0xe1, 0x43, 0xfd, 0x02, 0x99, 0x78, 0x9b, 0x5c,
	0x49, 0xb2, 0x0d, 0xfe, 0x94, 0xc7, 0xd4, 0xeb, 0x79, 0xfd, 0x46, 0x64, 0x3e, 0x49, 0x07, 0x2a,
	0x3c, 0xa6, 0x95, 0x9e, 0xd7, 0xaf, 0x46, 0x15, 0x1e, 0x13, 0x02, 0xd5, 0x84, 0x09, 0xa4, 0xbe,
	0xa5, 0xd8, 0x6f, 0x42, 0xa1, 0x36, 0xca, 0x90, 0x69, 0x8c, 0x69, 0xb5, 0xe7, 0xf5, 0xfd, 0xa8,
	0x80, 0x26, 0x33, 0x4d, 0x63, 0x9b, 0xd9, 0x74, 0x99, 0x1c, 0xce, 0xce, 0xc8, 0x8c, 0x06, 0x56,
	0xaa, 0x80, 0xa4, 0x0b, 0x75, 0x99, 0x62, 0x66, 0x53, 0x35, 0x9b, 0x9a, 0x61, 0xf2, 0x3f, 0x6c,
	0xca, 0x6f, 0x09, 0x66, 0xb4, 0x6e, 0x13, 0x0e, 0x90, 0x67, 0x10, 0x64, 0x28, 0x58, 0xf6, 0x95,
	0x36, 0x6c, 0x38, 0x47, 0x26, 0x2e, 0x98, 0xd2, 0x98, 0x51, 0x70, 0x71, 0x87, 0x1c, 0x7f, 0xcc,
	0x65, 0x42, 0x9b, 0x05, 0xdf, 0x20, 0x13, 0x57, 0x9a, 0xe9, 0xa9, 0xa2, 0xad, 0x9e, 0xd7, 0x6f,
	0x47, 0x39, 0x32, 0x55, 0x27, 0x5c, 0x70, 0x4d, 0xdb, 0x36, 0xec, 0x80, 0x61, 0xa7, 0x2c, 0xc3,
	0x44, 0xd3, 0x8e, 0x53, 0x71, 0x88, 0xec, 0x02, 0x18, 0xaf, 0x95, 0x66, 0x89, 0x56, 0x74, 0xab,
	0xe7, 0xf7, 0x1b, 0x51, 0x29, 0x62, 0x1c, 0xd4, 0x6c, 0xac, 0xe8, 0xb6, 0xcd, 0xd8, 0x6f, 0xe3,
	0x86, 0x40, 0x71, 0x89, 0x99, 0xa2, 0xff, 0xd9, 0x70, 0x01, 0xc3, 0xdf, 0x1e, 0x40, 0x84, 0xd7,
	0x66, 0x42, 0x87, 0x71, 0x3c, 0x37, 0xc0, 0x2b, 0x1b, 0x50, 0x0c, 0xa5, 0x52, 0x1a, 0xca, 0xdc,
	0x14, 0x7f, 0xd1, 0x94, 0xbc, 0xf9, 0xea, 0x9d, 0xe6, 0xcb, 0xb6, 0x6f, 0x2e, 0xd8, 0x3e, 0x37,
	0x32, 0x58, 0x34, 0x32, 0xb7, 0xa0, 0x76, 0xc7, 0x82, 0x99, 0x61, 0xf5, 0xb2, 0x61, 0x45, 0xe3,
	0x8d, 0x79, 0xe3, 0x21, 0x87, 0x76, 0xde, 0xdd, 0x07, 0xbb, 0x18, 0x0f, 0x6c, 0xe0, 0xdf, 0x34,
	0x57, 0x6e, 0xa2, 0x7a, 0xb7, 0x89, 0xf0, 0xbb, 0x29, 0x95, 0x4e, 0x6e, 0x67, 0xcb, 0xfe, 0x66,
	0x36, 0x6e, 0x53, 0xad, 0x79, 0xf0, 0x72, 0x70, 0xff, 0xf5, 0x0c, 0xec, 0x91, 0xa1, 0xa5, 0xcd,
	0xf6, 0xe1, 0x15, 0x54, 0x79, 0x72, 0x25, 0xed, 0x8d, 0x9a, 0x07, 0x2f, 0x1e, 0x3a, 0x56, 0x14,
	0x89, 0x2c, 0x33, 0xfc, 0xe5, 0x95, 0x8a, 0xbf, 0xe3, 0x4a, 0x3f, 0xa9, 0xf8, 0x84, 0x2b, 0x4d,
	0x2b, 0x3d, 0x7f, 0x75, 0x71, 0xc3, 0x34, 0xd3, 0xd0, 0x52, 0xb3, 0x89, 0xf5, 0xaa, 0x1d, 0x39,
	0x60, 0xa2, 0x29, 0x1b, 0xa3, 0xb2, 0x3e, 0xb5, 0x23, 0x07, 0x0e, 0x7e, 0xd6, 0xa0, 0x69, 0x8e,
	0x0f, 0x31, 0xbb, 0xe1, 0x23, 0x24, 0xa7, 0x10, 0x1c, 0xc6, 0xf1, 0xfb, 0x04, 0xc9, 0xee, 0xc3,
	0x17, 0x2c, 0x36, 0xb3, 0xbb, 0xf7, 0x68, 0x03, 0xc5, 0x75, 0xc2, 0x0d, 0x72, 0x0e, 0xe0, 0xe6,
	0x7c, 0xc4, 0x14, 0x92, 0xbd, 0x25, 0x92, 0x8e, 0xd6, 0xdd, 0x79, 0x54, 0x35, 0x57, 0x3c, 0x85,
	0x46, 0x84, 0x42, 0xde, 0xa0, 0xb9, 0xe3, 0x23, 0x26, 0x5e, 0x4f, 0x51, 0x69, 0xc3, 0x5f, 0x2d,
	0x77, 0x06, 0xc1, 0x09, 0xea, 0xb5, 0xb4, 0xd6, 0x6a, 0xf8, 0x0c, 0x82, 0x21, 0xb2, 0x6c, 0xf4,
	0xe5, 0xa9, 0x7a, 0x66, 0x69, 0xc2, 0x0d, 0xf2, 0x09, 0xb6, 0x4e, 0x50, 0x1b, 0x70, 0x74, 0x7b,
	0xcc, 0x27, 0xe6, 0x11, 0xee, 0x2d, 0x11, 0x76, 0x94, 0xf5, 0xa4, 0x3f, 0x42, 0xeb, 0x04, 0xb5,
	0xd9, 0x36, 0xae, 0x34, 0x1f, 0xad, 0xa3, 0x1b, 0x2e, 0x5d, 0x5a, 0x2b, 0x13, 0x6e, 0x90, 0x0b,
	0xe8, 0xe4, 0x43, 0x5f, 0xe7, 0xca, 0xeb, 0x0e, 0x7e, 0x08, 0x2d, 0x47, 0x75, 0xef, 0x83, 0x84,
	0x4b, 0xfd, 0xd5, 0xc7, 0x13, 0x36, 0x5e, 0x2d, 0x7a, 0x0e, 0xad, 0xc3, 0x34, 0xc5, 0x24, 0x3e,
	0xb5, 0x3f, 0xbf, 0x4b, 0x87, 0x66, 0x6c, 0x5b, 0xa2, 0x98, 0xbb, 0x1a, 0x41, 0x67, 0x38, 0xbd,
	0xd4, 0x19, 0x1b, 0xe9, 0x7f, 0xa5, 0x79, 0x44, 0x3e, 0x6f, 0x2f, 0xfe, 0xa7, 0x5f, 0x06, 0x36,
	0xf2, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xcb, 0xa3, 0x60, 0x1f, 0x08, 0x00, 0x00,
}
