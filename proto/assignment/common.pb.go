// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/assignment/common.proto

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

type ResultStatus int32

const (
	ResultStatus_Success     ResultStatus = 0
	ResultStatus_MaxLimit    ResultStatus = 1
	ResultStatus_Repeated    ResultStatus = 2
	ResultStatus_NotExisted  ResultStatus = 3
	ResultStatus_DBException ResultStatus = 4
	ResultStatus_Empty       ResultStatus = 5
)

var ResultStatus_name = map[int32]string{
	0: "Success",
	1: "MaxLimit",
	2: "Repeated",
	3: "NotExisted",
	4: "DBException",
	5: "Empty",
}

var ResultStatus_value = map[string]int32{
	"Success":     0,
	"MaxLimit":    1,
	"Repeated":    2,
	"NotExisted":  3,
	"DBException": 4,
	"Empty":       5,
}

func (x ResultStatus) String() string {
	return proto.EnumName(ResultStatus_name, int32(x))
}

func (ResultStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{0}
}

type AddressInfo struct {
	Country              string   `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Province             string   `protobuf:"bytes,2,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Zone                 string   `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"`
	Location             string   `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressInfo) Reset()         { *m = AddressInfo{} }
func (m *AddressInfo) String() string { return proto.CompactTextString(m) }
func (*AddressInfo) ProtoMessage()    {}
func (*AddressInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{0}
}

func (m *AddressInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressInfo.Unmarshal(m, b)
}
func (m *AddressInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressInfo.Marshal(b, m, deterministic)
}
func (m *AddressInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressInfo.Merge(m, src)
}
func (m *AddressInfo) XXX_Size() int {
	return xxx_messageInfo_AddressInfo.Size(m)
}
func (m *AddressInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AddressInfo proto.InternalMessageInfo

func (m *AddressInfo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *AddressInfo) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *AddressInfo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *AddressInfo) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *AddressInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type RequestInfo struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{1}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestFilter struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32   `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	Values               []string `protobuf:"bytes,6,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFilter) Reset()         { *m = RequestFilter{} }
func (m *RequestFilter) String() string { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()    {}
func (*RequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{2}
}

func (m *RequestFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFilter.Unmarshal(m, b)
}
func (m *RequestFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFilter.Marshal(b, m, deterministic)
}
func (m *RequestFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFilter.Merge(m, src)
}
func (m *RequestFilter) XXX_Size() int {
	return xxx_messageInfo_RequestFilter.Size(m)
}
func (m *RequestFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFilter proto.InternalMessageInfo

func (m *RequestFilter) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestFilter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestFilter) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RequestFilter) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RequestFilter) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type RequestUpdate struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Values               []string `protobuf:"bytes,6,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUpdate) Reset()         { *m = RequestUpdate{} }
func (m *RequestUpdate) String() string { return proto.CompactTextString(m) }
func (*RequestUpdate) ProtoMessage()    {}
func (*RequestUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{3}
}

func (m *RequestUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUpdate.Unmarshal(m, b)
}
func (m *RequestUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUpdate.Marshal(b, m, deterministic)
}
func (m *RequestUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUpdate.Merge(m, src)
}
func (m *RequestUpdate) XXX_Size() int {
	return xxx_messageInfo_RequestUpdate.Size(m)
}
func (m *RequestUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUpdate proto.InternalMessageInfo

func (m *RequestUpdate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestUpdate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestUpdate) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestUpdate) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type ReplyStatus struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyStatus) Reset()         { *m = ReplyStatus{} }
func (m *ReplyStatus) String() string { return proto.CompactTextString(m) }
func (*ReplyStatus) ProtoMessage()    {}
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{4}
}

func (m *ReplyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatus.Unmarshal(m, b)
}
func (m *ReplyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatus.Marshal(b, m, deterministic)
}
func (m *ReplyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatus.Merge(m, src)
}
func (m *ReplyStatus) XXX_Size() int {
	return xxx_messageInfo_ReplyStatus.Size(m)
}
func (m *ReplyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatus proto.InternalMessageInfo

func (m *ReplyStatus) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReplyStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ReplyInfo struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyInfo) Reset()         { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()    {}
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{5}
}

func (m *ReplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInfo.Unmarshal(m, b)
}
func (m *ReplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInfo.Merge(m, src)
}
func (m *ReplyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyInfo.Size(m)
}
func (m *ReplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInfo proto.InternalMessageInfo

func (m *ReplyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type DateInfo struct {
	Begin                string   `protobuf:"bytes,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End                  string   `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateInfo) Reset()         { *m = DateInfo{} }
func (m *DateInfo) String() string { return proto.CompactTextString(m) }
func (*DateInfo) ProtoMessage()    {}
func (*DateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{6}
}

func (m *DateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateInfo.Unmarshal(m, b)
}
func (m *DateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateInfo.Marshal(b, m, deterministic)
}
func (m *DateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateInfo.Merge(m, src)
}
func (m *DateInfo) XXX_Size() int {
	return xxx_messageInfo_DateInfo.Size(m)
}
func (m *DateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DateInfo proto.InternalMessageInfo

func (m *DateInfo) GetBegin() string {
	if m != nil {
		return m.Begin
	}
	return ""
}

func (m *DateInfo) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type RequestAddress struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Country              string   `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Province             string   `protobuf:"bytes,3,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Zone                 string   `protobuf:"bytes,5,opt,name=zone,proto3" json:"zone,omitempty"`
	Location             string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAddress) Reset()         { *m = RequestAddress{} }
func (m *RequestAddress) String() string { return proto.CompactTextString(m) }
func (*RequestAddress) ProtoMessage()    {}
func (*RequestAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{7}
}

func (m *RequestAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAddress.Unmarshal(m, b)
}
func (m *RequestAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAddress.Marshal(b, m, deterministic)
}
func (m *RequestAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAddress.Merge(m, src)
}
func (m *RequestAddress) XXX_Size() int {
	return xxx_messageInfo_RequestAddress.Size(m)
}
func (m *RequestAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAddress.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAddress proto.InternalMessageInfo

func (m *RequestAddress) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestAddress) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *RequestAddress) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *RequestAddress) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *RequestAddress) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *RequestAddress) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *RequestAddress) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestFlag struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 string   `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFlag) Reset()         { *m = RequestFlag{} }
func (m *RequestFlag) String() string { return proto.CompactTextString(m) }
func (*RequestFlag) ProtoMessage()    {}
func (*RequestFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{8}
}

func (m *RequestFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFlag.Unmarshal(m, b)
}
func (m *RequestFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFlag.Marshal(b, m, deterministic)
}
func (m *RequestFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFlag.Merge(m, src)
}
func (m *RequestFlag) XXX_Size() int {
	return xxx_messageInfo_RequestFlag.Size(m)
}
func (m *RequestFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFlag.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFlag proto.InternalMessageInfo

func (m *RequestFlag) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestFlag) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *RequestFlag) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestIntFlag struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 int32    `protobuf:"varint,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestIntFlag) Reset()         { *m = RequestIntFlag{} }
func (m *RequestIntFlag) String() string { return proto.CompactTextString(m) }
func (*RequestIntFlag) ProtoMessage()    {}
func (*RequestIntFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{9}
}

func (m *RequestIntFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestIntFlag.Unmarshal(m, b)
}
func (m *RequestIntFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestIntFlag.Marshal(b, m, deterministic)
}
func (m *RequestIntFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestIntFlag.Merge(m, src)
}
func (m *RequestIntFlag) XXX_Size() int {
	return xxx_messageInfo_RequestIntFlag.Size(m)
}
func (m *RequestIntFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestIntFlag.DiscardUnknown(m)
}

var xxx_messageInfo_RequestIntFlag proto.InternalMessageInfo

func (m *RequestIntFlag) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestIntFlag) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *RequestIntFlag) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestList struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []string `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestList) Reset()         { *m = RequestList{} }
func (m *RequestList) String() string { return proto.CompactTextString(m) }
func (*RequestList) ProtoMessage()    {}
func (*RequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{10}
}

func (m *RequestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestList.Unmarshal(m, b)
}
func (m *RequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestList.Marshal(b, m, deterministic)
}
func (m *RequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestList.Merge(m, src)
}
func (m *RequestList) XXX_Size() int {
	return xxx_messageInfo_RequestList.Size(m)
}
func (m *RequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestList proto.InternalMessageInfo

func (m *RequestList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestList) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyList struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	List                 []string     `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyList) Reset()         { *m = ReplyList{} }
func (m *ReplyList) String() string { return proto.CompactTextString(m) }
func (*ReplyList) ProtoMessage()    {}
func (*ReplyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{11}
}

func (m *ReplyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyList.Unmarshal(m, b)
}
func (m *ReplyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyList.Marshal(b, m, deterministic)
}
func (m *ReplyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyList.Merge(m, src)
}
func (m *ReplyList) XXX_Size() int {
	return xxx_messageInfo_ReplyList.Size(m)
}
func (m *ReplyList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyList proto.InternalMessageInfo

func (m *ReplyList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyStatistic struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Key                  string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Owner                string       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Count                uint32       `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyStatistic) Reset()         { *m = ReplyStatistic{} }
func (m *ReplyStatistic) String() string { return proto.CompactTextString(m) }
func (*ReplyStatistic) ProtoMessage()    {}
func (*ReplyStatistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_571eee8a2220d5e7, []int{12}
}

func (m *ReplyStatistic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatistic.Unmarshal(m, b)
}
func (m *ReplyStatistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatistic.Marshal(b, m, deterministic)
}
func (m *ReplyStatistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatistic.Merge(m, src)
}
func (m *ReplyStatistic) XXX_Size() int {
	return xxx_messageInfo_ReplyStatistic.Size(m)
}
func (m *ReplyStatistic) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatistic.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatistic proto.InternalMessageInfo

func (m *ReplyStatistic) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyStatistic) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReplyStatistic) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReplyStatistic) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterEnum("omo.msp.assignment.ResultStatus", ResultStatus_name, ResultStatus_value)
	proto.RegisterType((*AddressInfo)(nil), "omo.msp.assignment.AddressInfo")
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.assignment.RequestInfo")
	proto.RegisterType((*RequestFilter)(nil), "omo.msp.assignment.RequestFilter")
	proto.RegisterType((*RequestUpdate)(nil), "omo.msp.assignment.RequestUpdate")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.assignment.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.assignment.ReplyInfo")
	proto.RegisterType((*DateInfo)(nil), "omo.msp.assignment.DateInfo")
	proto.RegisterType((*RequestAddress)(nil), "omo.msp.assignment.RequestAddress")
	proto.RegisterType((*RequestFlag)(nil), "omo.msp.assignment.RequestFlag")
	proto.RegisterType((*RequestIntFlag)(nil), "omo.msp.assignment.RequestIntFlag")
	proto.RegisterType((*RequestList)(nil), "omo.msp.assignment.RequestList")
	proto.RegisterType((*ReplyList)(nil), "omo.msp.assignment.ReplyList")
	proto.RegisterType((*ReplyStatistic)(nil), "omo.msp.assignment.ReplyStatistic")
}

func init() {
	proto.RegisterFile("proto/assignment/common.proto", fileDescriptor_571eee8a2220d5e7)
}

var fileDescriptor_571eee8a2220d5e7 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xc1, 0x4e, 0xdc, 0x30,
	0x10, 0x6d, 0x36, 0xc9, 0xc2, 0xce, 0xb2, 0x34, 0xb2, 0x50, 0x15, 0x55, 0xaa, 0x8a, 0x72, 0x42,
	0x3d, 0x2c, 0x12, 0x3d, 0x70, 0x2e, 0x02, 0x24, 0x24, 0xda, 0xaa, 0x41, 0xed, 0xa1, 0x37, 0x93,
	0x0c, 0x2b, 0x97, 0xc4, 0x4e, 0x6d, 0x87, 0xb2, 0xfd, 0x01, 0xce, 0x55, 0x7f, 0xa6, 0x9f, 0x57,
	0xd9, 0x71, 0x76, 0x03, 0x4d, 0x2a, 0xb8, 0xcd, 0x9b, 0xf5, 0xbc, 0x79, 0x33, 0x7e, 0xce, 0xc2,
	0xab, 0x4a, 0x0a, 0x2d, 0xf6, 0xa9, 0x52, 0x6c, 0xc1, 0x4b, 0xe4, 0x7a, 0x3f, 0x13, 0x65, 0x29,
	0xf8, 0xdc, 0xe6, 0x09, 0x11, 0xa5, 0x98, 0x97, 0xaa, 0x9a, 0xaf, 0x0f, 0x24, 0x77, 0x1e, 0x4c,
	0xdf, 0xe5, 0xb9, 0x44, 0xa5, 0xce, 0xf8, 0x95, 0x20, 0x31, 0x6c, 0x64, 0xa2, 0xe6, 0x5a, 0x2e,
	0x63, 0x6f, 0xd7, 0xdb, 0x9b, 0xa4, 0x2d, 0x24, 0x2f, 0x61, 0xb3, 0x92, 0xe2, 0x86, 0xf1, 0x0c,
	0xe3, 0x91, 0xfd, 0x69, 0x85, 0x09, 0x81, 0x20, 0x63, 0x7a, 0x19, 0xfb, 0x36, 0x6f, 0x63, 0x93,
	0xfb, 0x29, 0x38, 0xc6, 0x41, 0x93, 0x33, 0xb1, 0xe1, 0x28, 0x44, 0x46, 0x35, 0x13, 0x3c, 0x0e,
	0x1b, 0x8e, 0x16, 0x27, 0x9f, 0x60, 0x9a, 0xe2, 0xf7, 0x1a, 0x95, 0xb6, 0x42, 0x76, 0x20, 0x14,
	0x3f, 0x38, 0x4a, 0x27, 0xa3, 0x01, 0x24, 0x02, 0xbf, 0x66, 0xb9, 0xeb, 0x6f, 0x42, 0x43, 0x29,
	0x2a, 0x94, 0x54, 0x0b, 0xe9, 0xda, 0xaf, 0x70, 0xf2, 0xcb, 0x83, 0x99, 0xe3, 0x3c, 0x65, 0x85,
	0x46, 0x39, 0xcc, 0x7a, 0x8d, 0xcb, 0x96, 0xf5, 0x1a, 0x97, 0xe6, 0xdc, 0x0d, 0x2d, 0x6a, 0x74,
	0x94, 0x0d, 0x30, 0x23, 0x55, 0x74, 0xd1, 0x8c, 0x34, 0x4b, 0x6d, 0x4c, 0x5e, 0xc0, 0x98, 0xd7,
	0xe5, 0x25, 0x4a, 0x3b, 0xd0, 0x2c, 0x75, 0xc8, 0xe4, 0x6d, 0x91, 0x8a, 0xc7, 0xbb, 0xfe, 0xde,
	0x24, 0x75, 0x28, 0xf9, 0xbd, 0xd6, 0xf4, 0xb9, 0xca, 0xa9, 0xc6, 0x47, 0x4f, 0xea, 0x54, 0xfa,
	0x3d, 0x2a, 0x83, 0xae, 0xca, 0xee, 0x46, 0xc2, 0xfb, 0x1b, 0x19, 0x54, 0x75, 0x68, 0x96, 0x5f,
	0x15, 0xcb, 0x0b, 0x4d, 0x75, 0xad, 0xec, 0x7d, 0x8a, 0x1c, 0xad, 0xa2, 0x59, 0x6a, 0x63, 0xd3,
	0x0c, 0xa5, 0x14, 0xd2, 0x49, 0x6a, 0x40, 0xf2, 0x05, 0x26, 0xb6, 0xd0, 0xde, 0x99, 0xd3, 0xec,
	0xad, 0x35, 0x1f, 0xc2, 0x58, 0x59, 0x4a, 0x5b, 0x35, 0x3d, 0x78, 0x3d, 0xff, 0xd7, 0x83, 0xf3,
	0x4e, 0xe7, 0xd4, 0x1d, 0x4f, 0x0e, 0x60, 0xf3, 0x98, 0x6a, 0x6c, 0xad, 0x70, 0x89, 0x0b, 0xc6,
	0xdb, 0x05, 0x59, 0x60, 0x9a, 0x21, 0x5f, 0x2d, 0x08, 0x79, 0x9e, 0xfc, 0xf1, 0x60, 0xdb, 0xad,
	0xd6, 0x59, 0xba, 0x47, 0x51, 0xc7, 0xe0, 0xa3, 0x61, 0x83, 0xfb, 0x03, 0x06, 0x0f, 0x7a, 0x0c,
	0x1e, 0x0e, 0x18, 0x7c, 0x7c, 0xdf, 0xe0, 0xf7, 0xee, 0x65, 0xe3, 0x81, 0x53, 0x3f, 0xae, 0xcc,
	0x7f, 0x5a, 0xd0, 0x45, 0x8f, 0x6c, 0x02, 0xc1, 0x55, 0x41, 0x17, 0x4e, 0xb3, 0x8d, 0xff, 0x6b,
	0xfd, 0x74, 0xb5, 0x8a, 0x33, 0xfe, 0x18, 0xce, 0xf0, 0x11, 0x9c, 0x6b, 0x91, 0xe7, 0x4c, 0xe9,
	0x1e, 0xc2, 0x6e, 0xf1, 0xe8, 0x81, 0xf3, 0x08, 0x04, 0x05, 0x53, 0x3a, 0xf6, 0xad, 0xef, 0x6c,
	0x9c, 0x7c, 0x73, 0xe6, 0x19, 0xa0, 0x6b, 0x4b, 0x46, 0xeb, 0x92, 0x8e, 0xa1, 0xfc, 0xa7, 0x19,
	0xea, 0xce, 0x9a, 0xc3, 0xe5, 0x99, 0xd2, 0x2c, 0xeb, 0x70, 0x79, 0x4f, 0xe2, 0xea, 0xff, 0x5e,
	0x34, 0x6f, 0xd8, 0xef, 0xbe, 0xe1, 0x1d, 0x08, 0xad, 0xb9, 0xdc, 0x07, 0xa3, 0x01, 0x6f, 0x32,
	0xd8, 0x4a, 0x51, 0xd5, 0x85, 0x76, 0x8f, 0x6d, 0x0a, 0x1b, 0x17, 0x75, 0x96, 0xa1, 0x52, 0xd1,
	0x33, 0xb2, 0x05, 0x9b, 0xef, 0xe9, 0xed, 0x39, 0x2b, 0x99, 0x8e, 0x3c, 0x83, 0x52, 0xac, 0x90,
	0x6a, 0xcc, 0xa3, 0x11, 0xd9, 0x06, 0xf8, 0x20, 0xf4, 0xc9, 0x2d, 0x53, 0x06, 0xfb, 0xe4, 0x39,
	0x4c, 0x8f, 0x8f, 0x4e, 0x6e, 0x33, 0xac, 0x8c, 0xbf, 0xa2, 0x80, 0x4c, 0x20, 0x3c, 0x29, 0x2b,
	0xbd, 0x8c, 0xc2, 0x23, 0xf2, 0x35, 0x7a, 0xf8, 0x67, 0x70, 0x39, 0xb6, 0x99, 0xb7, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x00, 0xff, 0xf7, 0x87, 0x27, 0x06, 0x00, 0x00,
}