// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/assignment/question.proto

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

type QuestionInfo struct {
	Uid                  string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64            `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64             `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64             `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string            `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string            `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string            `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string            `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Quote                string            `protobuf:"bytes,10,opt,name=quote,proto3" json:"quote,omitempty"`
	Cd                   uint32            `protobuf:"varint,11,opt,name=cd,proto3" json:"cd,omitempty"`
	Answers              []uint32          `protobuf:"varint,21,rep,packed,name=answers,proto3" json:"answers,omitempty"`
	Assets               []string          `protobuf:"bytes,22,rep,name=assets,proto3" json:"assets,omitempty"`
	Options              []*QuestionOption `protobuf:"bytes,31,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *QuestionInfo) Reset()         { *m = QuestionInfo{} }
func (m *QuestionInfo) String() string { return proto.CompactTextString(m) }
func (*QuestionInfo) ProtoMessage()    {}
func (*QuestionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf27896c1e9fe1e7, []int{0}
}

func (m *QuestionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionInfo.Unmarshal(m, b)
}
func (m *QuestionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionInfo.Marshal(b, m, deterministic)
}
func (m *QuestionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionInfo.Merge(m, src)
}
func (m *QuestionInfo) XXX_Size() int {
	return xxx_messageInfo_QuestionInfo.Size(m)
}
func (m *QuestionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionInfo proto.InternalMessageInfo

func (m *QuestionInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *QuestionInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QuestionInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *QuestionInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *QuestionInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *QuestionInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *QuestionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QuestionInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *QuestionInfo) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *QuestionInfo) GetCd() uint32 {
	if m != nil {
		return m.Cd
	}
	return 0
}

func (m *QuestionInfo) GetAnswers() []uint32 {
	if m != nil {
		return m.Answers
	}
	return nil
}

func (m *QuestionInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *QuestionInfo) GetOptions() []*QuestionOption {
	if m != nil {
		return m.Options
	}
	return nil
}

type QuestionOption struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuestionOption) Reset()         { *m = QuestionOption{} }
func (m *QuestionOption) String() string { return proto.CompactTextString(m) }
func (*QuestionOption) ProtoMessage()    {}
func (*QuestionOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf27896c1e9fe1e7, []int{1}
}

func (m *QuestionOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionOption.Unmarshal(m, b)
}
func (m *QuestionOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionOption.Marshal(b, m, deterministic)
}
func (m *QuestionOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionOption.Merge(m, src)
}
func (m *QuestionOption) XXX_Size() int {
	return xxx_messageInfo_QuestionOption.Size(m)
}
func (m *QuestionOption) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionOption.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionOption proto.InternalMessageInfo

func (m *QuestionOption) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QuestionOption) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type ReqQuestionAdd struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string            `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Category             string            `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Quote                string            `protobuf:"bytes,5,opt,name=quote,proto3" json:"quote,omitempty"`
	Operator             string            `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Cd                   uint32            `protobuf:"varint,7,opt,name=cd,proto3" json:"cd,omitempty"`
	Answers              []uint32          `protobuf:"varint,20,rep,packed,name=answers,proto3" json:"answers,omitempty"`
	Options              []*QuestionOption `protobuf:"bytes,21,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReqQuestionAdd) Reset()         { *m = ReqQuestionAdd{} }
func (m *ReqQuestionAdd) String() string { return proto.CompactTextString(m) }
func (*ReqQuestionAdd) ProtoMessage()    {}
func (*ReqQuestionAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf27896c1e9fe1e7, []int{2}
}

func (m *ReqQuestionAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqQuestionAdd.Unmarshal(m, b)
}
func (m *ReqQuestionAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqQuestionAdd.Marshal(b, m, deterministic)
}
func (m *ReqQuestionAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQuestionAdd.Merge(m, src)
}
func (m *ReqQuestionAdd) XXX_Size() int {
	return xxx_messageInfo_ReqQuestionAdd.Size(m)
}
func (m *ReqQuestionAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQuestionAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQuestionAdd proto.InternalMessageInfo

func (m *ReqQuestionAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqQuestionAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqQuestionAdd) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ReqQuestionAdd) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *ReqQuestionAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqQuestionAdd) GetCd() uint32 {
	if m != nil {
		return m.Cd
	}
	return 0
}

func (m *ReqQuestionAdd) GetAnswers() []uint32 {
	if m != nil {
		return m.Answers
	}
	return nil
}

func (m *ReqQuestionAdd) GetOptions() []*QuestionOption {
	if m != nil {
		return m.Options
	}
	return nil
}

type ReqQuestionUpdate struct {
	Uid                  string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string            `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string            `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Cd                   uint32            `protobuf:"varint,5,opt,name=cd,proto3" json:"cd,omitempty"`
	Category             string            `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Answers              []uint32          `protobuf:"varint,20,rep,packed,name=answers,proto3" json:"answers,omitempty"`
	Options              []*QuestionOption `protobuf:"bytes,21,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReqQuestionUpdate) Reset()         { *m = ReqQuestionUpdate{} }
func (m *ReqQuestionUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqQuestionUpdate) ProtoMessage()    {}
func (*ReqQuestionUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf27896c1e9fe1e7, []int{3}
}

func (m *ReqQuestionUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqQuestionUpdate.Unmarshal(m, b)
}
func (m *ReqQuestionUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqQuestionUpdate.Marshal(b, m, deterministic)
}
func (m *ReqQuestionUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQuestionUpdate.Merge(m, src)
}
func (m *ReqQuestionUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqQuestionUpdate.Size(m)
}
func (m *ReqQuestionUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQuestionUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQuestionUpdate proto.InternalMessageInfo

func (m *ReqQuestionUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqQuestionUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqQuestionUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqQuestionUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqQuestionUpdate) GetCd() uint32 {
	if m != nil {
		return m.Cd
	}
	return 0
}

func (m *ReqQuestionUpdate) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ReqQuestionUpdate) GetAnswers() []uint32 {
	if m != nil {
		return m.Answers
	}
	return nil
}

func (m *ReqQuestionUpdate) GetOptions() []*QuestionOption {
	if m != nil {
		return m.Options
	}
	return nil
}

type ReplyQuestionOne struct {
	Info                 *QuestionInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyQuestionOne) Reset()         { *m = ReplyQuestionOne{} }
func (m *ReplyQuestionOne) String() string { return proto.CompactTextString(m) }
func (*ReplyQuestionOne) ProtoMessage()    {}
func (*ReplyQuestionOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf27896c1e9fe1e7, []int{4}
}

func (m *ReplyQuestionOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyQuestionOne.Unmarshal(m, b)
}
func (m *ReplyQuestionOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyQuestionOne.Marshal(b, m, deterministic)
}
func (m *ReplyQuestionOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyQuestionOne.Merge(m, src)
}
func (m *ReplyQuestionOne) XXX_Size() int {
	return xxx_messageInfo_ReplyQuestionOne.Size(m)
}
func (m *ReplyQuestionOne) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyQuestionOne.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyQuestionOne proto.InternalMessageInfo

func (m *ReplyQuestionOne) GetInfo() *QuestionInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyQuestionOne) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyQuestionList struct {
	Total                uint32          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32          `protobuf:"varint,2,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32          `protobuf:"varint,3,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*QuestionInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus    `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyQuestionList) Reset()         { *m = ReplyQuestionList{} }
func (m *ReplyQuestionList) String() string { return proto.CompactTextString(m) }
func (*ReplyQuestionList) ProtoMessage()    {}
func (*ReplyQuestionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf27896c1e9fe1e7, []int{5}
}

func (m *ReplyQuestionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyQuestionList.Unmarshal(m, b)
}
func (m *ReplyQuestionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyQuestionList.Marshal(b, m, deterministic)
}
func (m *ReplyQuestionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyQuestionList.Merge(m, src)
}
func (m *ReplyQuestionList) XXX_Size() int {
	return xxx_messageInfo_ReplyQuestionList.Size(m)
}
func (m *ReplyQuestionList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyQuestionList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyQuestionList proto.InternalMessageInfo

func (m *ReplyQuestionList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyQuestionList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplyQuestionList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplyQuestionList) GetList() []*QuestionInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyQuestionList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*QuestionInfo)(nil), "omo.msp.assignment.QuestionInfo")
	proto.RegisterType((*QuestionOption)(nil), "omo.msp.assignment.QuestionOption")
	proto.RegisterType((*ReqQuestionAdd)(nil), "omo.msp.assignment.ReqQuestionAdd")
	proto.RegisterType((*ReqQuestionUpdate)(nil), "omo.msp.assignment.ReqQuestionUpdate")
	proto.RegisterType((*ReplyQuestionOne)(nil), "omo.msp.assignment.ReplyQuestionOne")
	proto.RegisterType((*ReplyQuestionList)(nil), "omo.msp.assignment.ReplyQuestionList")
}

func init() {
	proto.RegisterFile("proto/assignment/question.proto", fileDescriptor_bf27896c1e9fe1e7)
}

var fileDescriptor_bf27896c1e9fe1e7 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xc7, 0xe3, 0x8f, 0x38, 0x64, 0x02, 0x81, 0xae, 0x00, 0x59, 0x91, 0x10, 0xae, 0x55, 0xa4,
	0x9c, 0x82, 0x94, 0x22, 0xf5, 0xd2, 0x0b, 0x1c, 0x8a, 0x2a, 0x15, 0x2a, 0x16, 0xaa, 0x4a, 0xad,
	0x7a, 0x70, 0xed, 0x05, 0xad, 0x1a, 0x7b, 0x8d, 0x77, 0x03, 0xe5, 0xd8, 0xf7, 0xeb, 0x23, 0xf4,
	0x25, 0x5a, 0xf5, 0x01, 0xaa, 0x1d, 0x7f, 0x60, 0xc0, 0x24, 0xe1, 0xd0, 0x9b, 0x67, 0x67, 0xe6,
	0xef, 0xf9, 0xff, 0x66, 0x9d, 0xc0, 0x76, 0x9a, 0x09, 0x25, 0x76, 0x03, 0x29, 0xf9, 0x45, 0x12,
	0xb3, 0x44, 0xed, 0x5e, 0x4e, 0x99, 0x54, 0x5c, 0x24, 0x23, 0xcc, 0x10, 0x22, 0x62, 0x31, 0x8a,
	0x65, 0x3a, 0xba, 0x2d, 0x19, 0x6c, 0x3d, 0x68, 0x0a, 0x45, 0x1c, 0x97, 0x2d, 0xfe, 0x2f, 0x13,
	0x96, 0x4f, 0x0a, 0x95, 0xb7, 0xc9, 0xb9, 0x20, 0x6b, 0x60, 0x4d, 0x79, 0xe4, 0x1a, 0x9e, 0x31,
	0xec, 0x52, 0xfd, 0x48, 0xfa, 0x60, 0xf2, 0xc8, 0x35, 0x3d, 0x63, 0x68, 0x53, 0x93, 0x47, 0xc4,
	0x85, 0x4e, 0x98, 0xb1, 0x40, 0xb1, 0xc8, 0xb5, 0x3c, 0x63, 0x68, 0xd1, 0x32, 0xd4, 0x99, 0x69,
	0x1a, 0x61, 0xc6, 0xce, 0x33, 0x45, 0x58, 0xf5, 0x88, 0xcc, 0x6d, 0xa3, 0x72, 0x19, 0x92, 0x01,
	0x2c, 0x89, 0x94, 0x65, 0x98, 0x72, 0x30, 0x55, 0xc5, 0x84, 0x80, 0x9d, 0x04, 0x31, 0x73, 0x3b,
	0x78, 0x8e, 0xcf, 0x64, 0x13, 0x9c, 0x8c, 0xc5, 0x41, 0xf6, 0xcd, 0x5d, 0xc2, 0xd3, 0x22, 0x22,
	0xeb, 0xd0, 0xbe, 0x9c, 0x0a, 0xc5, 0x5c, 0xc0, 0xe3, 0x3c, 0xd0, 0xb3, 0x87, 0x91, 0xdb, 0xf3,
	0x8c, 0xe1, 0x0a, 0x35, 0x43, 0x9c, 0x23, 0x48, 0xe4, 0x35, 0xcb, 0xa4, 0xbb, 0xe1, 0x59, 0xc3,
	0x15, 0x5a, 0x86, 0x5a, 0x37, 0x90, 0x92, 0x29, 0xe9, 0x6e, 0x7a, 0x96, 0xd6, 0xcd, 0x23, 0xf2,
	0x1a, 0x3a, 0x22, 0xd5, 0x74, 0xa4, 0xbb, 0xed, 0x59, 0xc3, 0xde, 0xd8, 0x1f, 0x3d, 0xa4, 0x3c,
	0x2a, 0x11, 0xbe, 0xc7, 0x52, 0x5a, 0xb6, 0xf8, 0x7b, 0xd0, 0xbf, 0x9b, 0x2a, 0x68, 0x1a, 0xf9,
	0x44, 0x3c, 0xd2, 0x1e, 0x23, 0x26, 0x43, 0xe4, 0xdb, 0xa5, 0xf8, 0xec, 0xff, 0x31, 0xa0, 0x4f,
	0xd9, 0x65, 0xd9, 0xb9, 0x1f, 0x45, 0x15, 0x0a, 0xa3, 0x11, 0x85, 0x79, 0x07, 0xc5, 0x00, 0x96,
	0xc2, 0x40, 0xb1, 0x0b, 0x91, 0xdd, 0xe0, 0x86, 0xba, 0xb4, 0x8a, 0x6f, 0x31, 0xb5, 0xeb, 0x98,
	0x66, 0x2d, 0x21, 0x47, 0xd8, 0x69, 0x42, 0xb8, 0x7e, 0x17, 0x61, 0x0d, 0xd5, 0xc6, 0xd3, 0x51,
	0xfd, 0x36, 0xe0, 0x59, 0xcd, 0xf4, 0x07, 0xbc, 0x39, 0x0d, 0xd7, 0xb1, 0x24, 0x61, 0x36, 0x92,
	0xb0, 0xee, 0x93, 0xa8, 0x7c, 0xd9, 0x8d, 0xbe, 0xda, 0x95, 0xaf, 0x3a, 0x35, 0xe7, 0x1e, 0xb5,
	0xff, 0xe5, 0xf9, 0x87, 0x01, 0x6b, 0x94, 0xa5, 0x93, 0x9b, 0xaa, 0x20, 0x61, 0x64, 0x0f, 0x6c,
	0x9e, 0x9c, 0x0b, 0xf4, 0xdc, 0x1b, 0x7b, 0xb3, 0xf4, 0xf4, 0x17, 0x4b, 0xb1, 0x9a, 0xbc, 0x02,
	0x47, 0xaa, 0x40, 0x4d, 0x25, 0x82, 0xe9, 0x8d, 0xb7, 0x9b, 0xfa, 0xf0, 0x5d, 0xa7, 0x58, 0x46,
	0x8b, 0x72, 0xff, 0x27, 0x72, 0xaf, 0xcd, 0xf0, 0x8e, 0x4b, 0xa5, 0xef, 0x89, 0x12, 0x2a, 0x98,
	0x14, 0x37, 0x35, 0x0f, 0x34, 0x87, 0x34, 0xb8, 0x60, 0x47, 0xc1, 0x77, 0x7c, 0xcb, 0x0a, 0x2d,
	0xc3, 0x32, 0x73, 0x2c, 0xae, 0x71, 0x05, 0x45, 0xe6, 0x58, 0x5c, 0x6b, 0x3b, 0x13, 0x2e, 0x95,
	0x6b, 0x23, 0x9e, 0x05, 0xec, 0xe8, 0xea, 0x9a, 0x9d, 0xf6, 0x93, 0xec, 0x8c, 0xff, 0xda, 0xb0,
	0x5a, 0xea, 0x9d, 0xb2, 0xec, 0x8a, 0x87, 0x8c, 0x9c, 0x81, 0xb3, 0x1f, 0x45, 0x9a, 0xad, 0xdf,
	0x2c, 0x53, 0xff, 0xd4, 0x06, 0x2f, 0x1e, 0x7d, 0x55, 0x6d, 0x4b, 0x7e, 0x8b, 0x9c, 0x80, 0x73,
	0xc8, 0x94, 0x56, 0x7d, 0x64, 0x38, 0xfc, 0x75, 0xd6, 0x9e, 0x16, 0x96, 0x3c, 0x82, 0x2e, 0x65,
	0xb1, 0xb8, 0x62, 0x0b, 0xa9, 0x6e, 0x3d, 0xaa, 0xaa, 0xd3, 0x7e, 0x8b, 0x7c, 0x81, 0xd5, 0x43,
	0xa6, 0xf4, 0x3e, 0x0f, 0x6e, 0xde, 0xf0, 0x89, 0x62, 0x19, 0x79, 0x3e, 0x43, 0x34, 0x2f, 0x19,
	0xec, 0xcc, 0x1d, 0x56, 0x2b, 0xfa, 0x2d, 0xf2, 0x11, 0x96, 0x0f, 0x99, 0xd2, 0xfc, 0xb9, 0x54,
	0x3c, 0x5c, 0x44, 0xdb, 0x9f, 0xb9, 0x46, 0x94, 0xf1, 0x5b, 0xe4, 0x33, 0x40, 0xfe, 0xf9, 0x1f,
	0x04, 0x92, 0x91, 0x9d, 0x39, 0x3b, 0xcb, 0x4b, 0x17, 0x66, 0x7c, 0x06, 0xfd, 0x42, 0x7c, 0x11,
	0x26, 0x85, 0xf8, 0x3c, 0xd4, 0x07, 0xe4, 0xd3, 0xda, 0xfd, 0x3f, 0xda, 0xaf, 0x0e, 0x9e, 0xbc,
	0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x62, 0x41, 0xe6, 0x94, 0xb8, 0x07, 0x00, 0x00,
}
