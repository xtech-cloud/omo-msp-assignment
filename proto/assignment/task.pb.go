// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/assignment/task.proto

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

type RecordInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Created              int64    `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Creator              string   `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Status               uint32   `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	Executor             string   `protobuf:"bytes,7,opt,name=executor,proto3" json:"executor,omitempty"`
	Tags                 []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string `protobuf:"bytes,9,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordInfo) Reset()         { *m = RecordInfo{} }
func (m *RecordInfo) String() string { return proto.CompactTextString(m) }
func (*RecordInfo) ProtoMessage()    {}
func (*RecordInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{0}
}

func (m *RecordInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordInfo.Unmarshal(m, b)
}
func (m *RecordInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordInfo.Marshal(b, m, deterministic)
}
func (m *RecordInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordInfo.Merge(m, src)
}
func (m *RecordInfo) XXX_Size() int {
	return xxx_messageInfo_RecordInfo.Size(m)
}
func (m *RecordInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RecordInfo proto.InternalMessageInfo

func (m *RecordInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RecordInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RecordInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RecordInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RecordInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RecordInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *RecordInfo) GetExecutor() string {
	if m != nil {
		return m.Executor
	}
	return ""
}

func (m *RecordInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *RecordInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type TaskInfo struct {
	Uid                  string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64         `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64         `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string        `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string        `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Target               string        `protobuf:"bytes,7,opt,name=target,proto3" json:"target,omitempty"`
	Name                 string        `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string        `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Type                 uint32        `protobuf:"varint,10,opt,name=type,proto3" json:"type,omitempty"`
	Status               uint32        `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	Owner                string        `protobuf:"bytes,12,opt,name=owner,proto3" json:"owner,omitempty"`
	Way                  string        `protobuf:"bytes,13,opt,name=way,proto3" json:"way,omitempty"`
	Duration             *DateInfo     `protobuf:"bytes,14,opt,name=duration,proto3" json:"duration,omitempty"`
	Executors            []string      `protobuf:"bytes,15,rep,name=executors,proto3" json:"executors,omitempty"`
	Pretasks             []string      `protobuf:"bytes,16,rep,name=pretasks,proto3" json:"pretasks,omitempty"`
	Tags                 []string      `protobuf:"bytes,17,rep,name=tags,proto3" json:"tags,omitempty"`
	Regions              []string      `protobuf:"bytes,18,rep,name=regions,proto3" json:"regions,omitempty"`
	Assets               []string      `protobuf:"bytes,19,rep,name=assets,proto3" json:"assets,omitempty"`
	Records              []*RecordInfo `protobuf:"bytes,20,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TaskInfo) Reset()         { *m = TaskInfo{} }
func (m *TaskInfo) String() string { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()    {}
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{1}
}

func (m *TaskInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInfo.Unmarshal(m, b)
}
func (m *TaskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInfo.Marshal(b, m, deterministic)
}
func (m *TaskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInfo.Merge(m, src)
}
func (m *TaskInfo) XXX_Size() int {
	return xxx_messageInfo_TaskInfo.Size(m)
}
func (m *TaskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInfo proto.InternalMessageInfo

func (m *TaskInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TaskInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TaskInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *TaskInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TaskInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *TaskInfo) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *TaskInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *TaskInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *TaskInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TaskInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TaskInfo) GetWay() string {
	if m != nil {
		return m.Way
	}
	return ""
}

func (m *TaskInfo) GetDuration() *DateInfo {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *TaskInfo) GetExecutors() []string {
	if m != nil {
		return m.Executors
	}
	return nil
}

func (m *TaskInfo) GetPretasks() []string {
	if m != nil {
		return m.Pretasks
	}
	return nil
}

func (m *TaskInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TaskInfo) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *TaskInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *TaskInfo) GetRecords() []*RecordInfo {
	if m != nil {
		return m.Records
	}
	return nil
}

type ReqTaskAdd struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 int32     `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Remark               string    `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string    `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Target               string    `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	Operator             string    `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Way                  string    `protobuf:"bytes,7,opt,name=way,proto3" json:"way,omitempty"`
	Duration             *DateInfo `protobuf:"bytes,8,opt,name=duration,proto3" json:"duration,omitempty"`
	Regions              []string  `protobuf:"bytes,9,rep,name=regions,proto3" json:"regions,omitempty"`
	Pretasks             []string  `protobuf:"bytes,10,rep,name=pretasks,proto3" json:"pretasks,omitempty"`
	Assets               []string  `protobuf:"bytes,11,rep,name=assets,proto3" json:"assets,omitempty"`
	Tags                 []string  `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReqTaskAdd) Reset()         { *m = ReqTaskAdd{} }
func (m *ReqTaskAdd) String() string { return proto.CompactTextString(m) }
func (*ReqTaskAdd) ProtoMessage()    {}
func (*ReqTaskAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{2}
}

func (m *ReqTaskAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTaskAdd.Unmarshal(m, b)
}
func (m *ReqTaskAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTaskAdd.Marshal(b, m, deterministic)
}
func (m *ReqTaskAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTaskAdd.Merge(m, src)
}
func (m *ReqTaskAdd) XXX_Size() int {
	return xxx_messageInfo_ReqTaskAdd.Size(m)
}
func (m *ReqTaskAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTaskAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTaskAdd proto.InternalMessageInfo

func (m *ReqTaskAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTaskAdd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqTaskAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqTaskAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqTaskAdd) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqTaskAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTaskAdd) GetWay() string {
	if m != nil {
		return m.Way
	}
	return ""
}

func (m *ReqTaskAdd) GetDuration() *DateInfo {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *ReqTaskAdd) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *ReqTaskAdd) GetPretasks() []string {
	if m != nil {
		return m.Pretasks
	}
	return nil
}

func (m *ReqTaskAdd) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *ReqTaskAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReplyTaskOne struct {
	Info                 *TaskInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyTaskOne) Reset()         { *m = ReplyTaskOne{} }
func (m *ReplyTaskOne) String() string { return proto.CompactTextString(m) }
func (*ReplyTaskOne) ProtoMessage()    {}
func (*ReplyTaskOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{3}
}

func (m *ReplyTaskOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTaskOne.Unmarshal(m, b)
}
func (m *ReplyTaskOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTaskOne.Marshal(b, m, deterministic)
}
func (m *ReplyTaskOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTaskOne.Merge(m, src)
}
func (m *ReplyTaskOne) XXX_Size() int {
	return xxx_messageInfo_ReplyTaskOne.Size(m)
}
func (m *ReplyTaskOne) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTaskOne.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTaskOne proto.InternalMessageInfo

func (m *ReplyTaskOne) GetInfo() *TaskInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyTaskOne) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyTaskList struct {
	Total                uint32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32       `protobuf:"varint,2,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32       `protobuf:"varint,3,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*TaskInfo  `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyTaskList) Reset()         { *m = ReplyTaskList{} }
func (m *ReplyTaskList) String() string { return proto.CompactTextString(m) }
func (*ReplyTaskList) ProtoMessage()    {}
func (*ReplyTaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{4}
}

func (m *ReplyTaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTaskList.Unmarshal(m, b)
}
func (m *ReplyTaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTaskList.Marshal(b, m, deterministic)
}
func (m *ReplyTaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTaskList.Merge(m, src)
}
func (m *ReplyTaskList) XXX_Size() int {
	return xxx_messageInfo_ReplyTaskList.Size(m)
}
func (m *ReplyTaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTaskList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTaskList proto.InternalMessageInfo

func (m *ReplyTaskList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyTaskList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplyTaskList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplyTaskList) GetList() []*TaskInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyTaskList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqTaskUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTaskUpdate) Reset()         { *m = ReqTaskUpdate{} }
func (m *ReqTaskUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqTaskUpdate) ProtoMessage()    {}
func (*ReqTaskUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{5}
}

func (m *ReqTaskUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTaskUpdate.Unmarshal(m, b)
}
func (m *ReqTaskUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTaskUpdate.Marshal(b, m, deterministic)
}
func (m *ReqTaskUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTaskUpdate.Merge(m, src)
}
func (m *ReqTaskUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqTaskUpdate.Size(m)
}
func (m *ReqTaskUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTaskUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTaskUpdate proto.InternalMessageInfo

func (m *ReqTaskUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqTaskUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTaskUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqTaskUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqTaskRecord struct {
	Task                 string   `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Creator              string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Executor             string   `protobuf:"bytes,5,opt,name=executor,proto3" json:"executor,omitempty"`
	Status               uint32   `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string `protobuf:"bytes,8,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTaskRecord) Reset()         { *m = ReqTaskRecord{} }
func (m *ReqTaskRecord) String() string { return proto.CompactTextString(m) }
func (*ReqTaskRecord) ProtoMessage()    {}
func (*ReqTaskRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{6}
}

func (m *ReqTaskRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTaskRecord.Unmarshal(m, b)
}
func (m *ReqTaskRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTaskRecord.Marshal(b, m, deterministic)
}
func (m *ReqTaskRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTaskRecord.Merge(m, src)
}
func (m *ReqTaskRecord) XXX_Size() int {
	return xxx_messageInfo_ReqTaskRecord.Size(m)
}
func (m *ReqTaskRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTaskRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTaskRecord proto.InternalMessageInfo

func (m *ReqTaskRecord) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *ReqTaskRecord) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ReqTaskRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTaskRecord) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqTaskRecord) GetExecutor() string {
	if m != nil {
		return m.Executor
	}
	return ""
}

func (m *ReqTaskRecord) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReqTaskRecord) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ReqTaskRecord) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReplyTaskRecords struct {
	Status               *ReplyStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Task                 string        `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	List                 []*RecordInfo `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyTaskRecords) Reset()         { *m = ReplyTaskRecords{} }
func (m *ReplyTaskRecords) String() string { return proto.CompactTextString(m) }
func (*ReplyTaskRecords) ProtoMessage()    {}
func (*ReplyTaskRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{7}
}

func (m *ReplyTaskRecords) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTaskRecords.Unmarshal(m, b)
}
func (m *ReplyTaskRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTaskRecords.Marshal(b, m, deterministic)
}
func (m *ReplyTaskRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTaskRecords.Merge(m, src)
}
func (m *ReplyTaskRecords) XXX_Size() int {
	return xxx_messageInfo_ReplyTaskRecords.Size(m)
}
func (m *ReplyTaskRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTaskRecords.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTaskRecords proto.InternalMessageInfo

func (m *ReplyTaskRecords) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTaskRecords) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *ReplyTaskRecords) GetList() []*RecordInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*RecordInfo)(nil), "omo.msp.assignment.RecordInfo")
	proto.RegisterType((*TaskInfo)(nil), "omo.msp.assignment.TaskInfo")
	proto.RegisterType((*ReqTaskAdd)(nil), "omo.msp.assignment.ReqTaskAdd")
	proto.RegisterType((*ReplyTaskOne)(nil), "omo.msp.assignment.ReplyTaskOne")
	proto.RegisterType((*ReplyTaskList)(nil), "omo.msp.assignment.ReplyTaskList")
	proto.RegisterType((*ReqTaskUpdate)(nil), "omo.msp.assignment.ReqTaskUpdate")
	proto.RegisterType((*ReqTaskRecord)(nil), "omo.msp.assignment.ReqTaskRecord")
	proto.RegisterType((*ReplyTaskRecords)(nil), "omo.msp.assignment.ReplyTaskRecords")
}

func init() {
	proto.RegisterFile("proto/assignment/task.proto", fileDescriptor_eb1a60d5b88542d3)
}

var fileDescriptor_eb1a60d5b88542d3 = []byte{
	// 921 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x4b, 0x8e, 0x1b, 0x37,
	0x10, 0x9d, 0x56, 0xb7, 0x7e, 0x25, 0x69, 0x3c, 0x61, 0x8c, 0x80, 0x50, 0xec, 0x44, 0x6e, 0x64,
	0xa1, 0x95, 0x1c, 0x4c, 0x16, 0xf1, 0x76, 0x8c, 0xc0, 0x83, 0x00, 0xf6, 0x24, 0x68, 0x39, 0x08,
	0x9c, 0x1d, 0xad, 0xa6, 0x95, 0xc6, 0x48, 0xcd, 0x36, 0x49, 0x59, 0xd6, 0x39, 0x72, 0x83, 0x9c,
	0x20, 0x87, 0x98, 0x4d, 0x2e, 0x91, 0xb3, 0x04, 0x45, 0xb2, 0x7f, 0x1a, 0xfd, 0x66, 0xd7, 0xc5,
	0x2a, 0x3e, 0x55, 0xbd, 0xf7, 0x58, 0x10, 0x7c, 0x9d, 0x49, 0xa1, 0xc5, 0x73, 0xa6, 0x54, 0x32,
	0x4f, 0x97, 0x3c, 0xd5, 0xcf, 0x35, 0x53, 0xb7, 0x13, 0x73, 0x4a, 0x88, 0x58, 0x8a, 0xc9, 0x52,
	0x65, 0x93, 0x32, 0x3d, 0x7c, 0x7a, 0xef, 0xc2, 0x4c, 0x2c, 0x97, 0x22, 0xb5, 0x57, 0xc2, 0xff,
	0x3c, 0x80, 0x88, 0xcf, 0x84, 0x8c, 0x7f, 0x4e, 0x3f, 0x08, 0x72, 0x01, 0xfe, 0x2a, 0x89, 0xa9,
	0x37, 0xf2, 0xc6, 0xdd, 0x08, 0x3f, 0x09, 0x85, 0xf6, 0x4c, 0x72, 0xa6, 0x79, 0x4c, 0x1b, 0x23,
	0x6f, 0xec, 0x47, 0x79, 0x58, 0x64, 0x84, 0xa4, 0xbe, 0xa9, 0xcf, 0x43, 0x42, 0x20, 0x48, 0xd9,
	0x92, 0xd3, 0xc0, 0x1c, 0x9b, 0x6f, 0xf2, 0x15, 0xb4, 0x94, 0x66, 0x7a, 0xa5, 0x68, 0x73, 0xe4,
	0x8d, 0x07, 0x91, 0x8b, 0xf0, 0x5c, 0xf2, 0x25, 0x93, 0xb7, 0xb4, 0x65, 0xaa, 0x5d, 0x44, 0x86,
	0xd0, 0xe1, 0x9f, 0xf9, 0x6c, 0x85, 0xf0, 0x6d, 0x93, 0x29, 0x62, 0xc4, 0xd7, 0x6c, 0xae, 0x68,
	0x67, 0xe4, 0x23, 0x3e, 0x7e, 0x23, 0x0e, 0x53, 0x8a, 0x6b, 0x45, 0xbb, 0xe6, 0xd4, 0x45, 0xe1,
	0xdf, 0x01, 0x74, 0xde, 0x32, 0x75, 0xbb, 0x67, 0xbc, 0x73, 0x68, 0x24, 0x76, 0xb2, 0x20, 0x6a,
	0xd4, 0xc7, 0xf5, 0xef, 0x8d, 0xbb, 0xca, 0x62, 0x93, 0x09, 0x6c, 0xc6, 0x85, 0x55, 0x22, 0x9a,
	0x75, 0x22, 0x86, 0xd0, 0x11, 0x19, 0x97, 0x26, 0x65, 0xc7, 0x2b, 0x62, 0x6c, 0x58, 0x33, 0x39,
	0xe7, 0xda, 0x8d, 0xe7, 0xa2, 0x82, 0xbc, 0x4e, 0x9d, 0x3c, 0x47, 0x52, 0xb7, 0x46, 0x12, 0x12,
	0xb1, 0xc9, 0x38, 0x05, 0x43, 0xa9, 0xf9, 0xae, 0x10, 0xdd, 0xab, 0x11, 0xfd, 0x18, 0x9a, 0x62,
	0x9d, 0x72, 0x49, 0xfb, 0x06, 0xc2, 0x06, 0xc8, 0xc8, 0x9a, 0x6d, 0xe8, 0xc0, 0x32, 0xb2, 0x66,
	0x1b, 0xf2, 0x02, 0x3a, 0xf1, 0x4a, 0x32, 0x9d, 0x88, 0x94, 0x9e, 0x8f, 0xbc, 0x71, 0xef, 0xf2,
	0xc9, 0xe4, 0xbe, 0xaf, 0x26, 0x3f, 0x31, 0xcd, 0x91, 0xd3, 0xa8, 0xa8, 0x26, 0x4f, 0xa0, 0x9b,
	0x4b, 0xa4, 0xe8, 0x23, 0xa3, 0x42, 0x79, 0x80, 0x5c, 0x64, 0x92, 0xa3, 0x5b, 0x15, 0xbd, 0x30,
	0xc9, 0x22, 0x2e, 0x04, 0xfd, 0xa2, 0x22, 0x28, 0x85, 0xb6, 0xe4, 0xf3, 0x44, 0xa4, 0x8a, 0x12,
	0x73, 0x9c, 0x87, 0x15, 0xa9, 0xbf, 0xac, 0x4a, 0x4d, 0x5e, 0xe0, 0x0d, 0xb4, 0xb2, 0xa2, 0x8f,
	0x47, 0xfe, 0xb8, 0x77, 0xf9, 0xcd, 0xae, 0xc6, 0x4b, 0xb7, 0x47, 0x79, 0x79, 0x78, 0xd7, 0xc0,
	0x57, 0xf0, 0x11, 0x7d, 0x72, 0x15, 0xc7, 0x85, 0x04, 0x5e, 0x45, 0x82, 0x9c, 0x6a, 0xb4, 0x4a,
	0xb3, 0xa4, 0xda, 0xc9, 0xe2, 0xd7, 0x64, 0x29, 0xa8, 0x0e, 0xaa, 0x54, 0x97, 0x82, 0x37, 0x6b,
	0x82, 0x1f, 0x32, 0x89, 0x93, 0xa7, 0xbd, 0x5b, 0x9e, 0xce, 0x83, 0xe4, 0xa9, 0x10, 0xda, 0xad,
	0x13, 0x5a, 0x95, 0x06, 0xb6, 0xa4, 0x29, 0xc9, 0xee, 0xd5, 0xc8, 0xce, 0x25, 0xeb, 0x97, 0x92,
	0x85, 0x1b, 0xe8, 0x47, 0x3c, 0x5b, 0x6c, 0x90, 0xc7, 0x5f, 0x52, 0x4e, 0xbe, 0x87, 0x20, 0x49,
	0x3f, 0x08, 0xc3, 0xe3, 0x9e, 0x3e, 0xf3, 0xa7, 0x19, 0x99, 0x4a, 0xf2, 0x63, 0x61, 0xde, 0x86,
	0xb9, 0xf3, 0xed, 0x6e, 0x05, 0xb3, 0xc5, 0x66, 0x6a, 0xca, 0x72, 0x77, 0x87, 0x77, 0x1e, 0x0c,
	0x8a, 0xdf, 0x7e, 0x9d, 0x28, 0x8d, 0x22, 0x68, 0xa1, 0xd9, 0xc2, 0xfc, 0xfa, 0x20, 0xb2, 0x01,
	0x92, 0x90, 0xb1, 0x39, 0x7f, 0xc3, 0x3e, 0x9b, 0x5f, 0x18, 0x44, 0x79, 0x98, 0x67, 0x6e, 0xc4,
	0xda, 0xa8, 0xe9, 0x32, 0x37, 0x62, 0x8d, 0x63, 0x2c, 0x12, 0xa5, 0x69, 0x60, 0x4c, 0x75, 0x64,
	0x0c, 0xac, 0xac, 0x8c, 0xd1, 0x7c, 0xd8, 0x18, 0x09, 0x4e, 0x61, 0x7c, 0xf8, 0x9b, 0x59, 0x2e,
	0x3b, 0x36, 0x56, 0x6e, 0xce, 0xc6, 0xce, 0xfd, 0xe0, 0x6f, 0x2f, 0xd1, 0xc2, 0x5a, 0x41, 0xdd,
	0x5a, 0xe1, 0xbf, 0x5e, 0xf1, 0x5b, 0xf6, 0x49, 0x58, 0x49, 0xd5, 0x6d, 0x6e, 0x7b, 0xfc, 0xae,
	0xee, 0xb6, 0xc6, 0xee, 0x25, 0xef, 0xef, 0xec, 0x23, 0xd8, 0xbb, 0xcc, 0x9b, 0x5b, 0xcb, 0xbc,
	0xdc, 0x57, 0xad, 0xda, 0xbe, 0xca, 0x0d, 0xd6, 0xde, 0xb9, 0xe4, 0x3b, 0xb5, 0x25, 0xff, 0x97,
	0x07, 0x17, 0x85, 0xfa, 0x76, 0x1a, 0x55, 0x11, 0xc1, 0x7b, 0x90, 0x08, 0x05, 0x0f, 0x8d, 0x0a,
	0x0f, 0x97, 0xce, 0x03, 0xfe, 0x49, 0x8b, 0xc5, 0xd4, 0x5e, 0xfe, 0xd3, 0x86, 0x1e, 0x36, 0x34,
	0xe5, 0xf2, 0x53, 0x32, 0xe3, 0xe4, 0x35, 0xb4, 0xae, 0xe2, 0x18, 0x1f, 0xc6, 0x9e, 0xfb, 0xf9,
	0x02, 0x1a, 0x8e, 0xf6, 0xb6, 0xea, 0x9e, 0x56, 0x78, 0x46, 0xde, 0x40, 0xeb, 0x9a, 0x6b, 0x44,
	0xdb, 0x33, 0xd8, 0xc7, 0x15, 0x57, 0x1a, 0xdb, 0x39, 0x11, 0xae, 0x1b, 0xf1, 0xa5, 0xf8, 0xc4,
	0x4f, 0x42, 0x7c, 0xba, 0x17, 0x11, 0xd3, 0xe1, 0x19, 0xb9, 0x81, 0xd6, 0x94, 0x33, 0x39, 0xfb,
	0xf3, 0x38, 0xd6, 0xb3, 0x83, 0xdd, 0xe1, 0x5b, 0x0e, 0xcf, 0xc8, 0x3b, 0x78, 0x74, 0xcd, 0x35,
	0x06, 0x2f, 0x37, 0xaf, 0x92, 0x85, 0xe6, 0x92, 0x3c, 0x3b, 0x00, 0x6c, 0x4b, 0x4e, 0x83, 0xfe,
	0x1d, 0xfa, 0xd7, 0x5c, 0xa3, 0x07, 0x12, 0xa5, 0x93, 0xd9, 0x29, 0xb8, 0xe1, 0x41, 0x2b, 0x19,
	0x98, 0xf0, 0x8c, 0xbc, 0x85, 0x73, 0xfb, 0x8a, 0x4f, 0x6a, 0xd9, 0x96, 0x1e, 0x67, 0xf6, 0x57,
	0x00, 0x87, 0xca, 0x14, 0xdf, 0x8b, 0x58, 0xae, 0x90, 0xe3, 0x88, 0x53, 0xe8, 0xdb, 0x52, 0xfb,
	0x0e, 0x48, 0x78, 0x50, 0x31, 0xfd, 0x6a, 0xc1, 0xe6, 0xc7, 0x41, 0xdf, 0x41, 0xff, 0x2a, 0xcb,
	0x78, 0x1a, 0xbb, 0xe5, 0x72, 0xa8, 0x51, 0x5b, 0x32, 0xfc, 0xee, 0xa0, 0x5a, 0xee, 0x59, 0x1b,
	0xe8, 0xf3, 0xe9, 0xea, 0xbd, 0x96, 0x6c, 0xa6, 0x1d, 0xf8, 0x51, 0x8f, 0x9d, 0x08, 0xfd, 0x92,
	0xfc, 0x71, 0xb1, 0xfd, 0x7f, 0xf9, 0x7d, 0xcb, 0x9c, 0xfc, 0xf0, 0x7f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x44, 0xfb, 0x9f, 0x8f, 0x7b, 0x0b, 0x00, 0x00,
}
