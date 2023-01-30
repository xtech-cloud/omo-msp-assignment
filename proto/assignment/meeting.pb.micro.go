// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/assignment/meeting.proto

package assignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for MeetingService service

type MeetingService interface {
	AddOne(ctx context.Context, in *ReqMeetingAdd, opts ...client.CallOption) (*ReplyMeetingOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyMeetingOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyMeetingList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateStatus(ctx context.Context, in *RequestIntFlag, opts ...client.CallOption) (*ReplyInfo, error)
}

type meetingService struct {
	c    client.Client
	name string
}

func NewMeetingService(name string, c client.Client) MeetingService {
	return &meetingService{
		c:    c,
		name: name,
	}
}

func (c *meetingService) AddOne(ctx context.Context, in *ReqMeetingAdd, opts ...client.CallOption) (*ReplyMeetingOne, error) {
	req := c.c.NewRequest(c.name, "MeetingService.AddOne", in)
	out := new(ReplyMeetingOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyMeetingOne, error) {
	req := c.c.NewRequest(c.name, "MeetingService.GetOne", in)
	out := new(ReplyMeetingOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "MeetingService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyMeetingList, error) {
	req := c.c.NewRequest(c.name, "MeetingService.GetListByFilter", in)
	out := new(ReplyMeetingList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "MeetingService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "MeetingService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingService) UpdateStatus(ctx context.Context, in *RequestIntFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "MeetingService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MeetingService service

type MeetingServiceHandler interface {
	AddOne(context.Context, *ReqMeetingAdd, *ReplyMeetingOne) error
	GetOne(context.Context, *RequestInfo, *ReplyMeetingOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyMeetingList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	UpdateStatus(context.Context, *RequestIntFlag, *ReplyInfo) error
}

func RegisterMeetingServiceHandler(s server.Server, hdlr MeetingServiceHandler, opts ...server.HandlerOption) error {
	type meetingService interface {
		AddOne(ctx context.Context, in *ReqMeetingAdd, out *ReplyMeetingOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyMeetingOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyMeetingList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		UpdateStatus(ctx context.Context, in *RequestIntFlag, out *ReplyInfo) error
	}
	type MeetingService struct {
		meetingService
	}
	h := &meetingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MeetingService{h}, opts...))
}

type meetingServiceHandler struct {
	MeetingServiceHandler
}

func (h *meetingServiceHandler) AddOne(ctx context.Context, in *ReqMeetingAdd, out *ReplyMeetingOne) error {
	return h.MeetingServiceHandler.AddOne(ctx, in, out)
}

func (h *meetingServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyMeetingOne) error {
	return h.MeetingServiceHandler.GetOne(ctx, in, out)
}

func (h *meetingServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.MeetingServiceHandler.RemoveOne(ctx, in, out)
}

func (h *meetingServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyMeetingList) error {
	return h.MeetingServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *meetingServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.MeetingServiceHandler.GetStatistic(ctx, in, out)
}

func (h *meetingServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.MeetingServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *meetingServiceHandler) UpdateStatus(ctx context.Context, in *RequestIntFlag, out *ReplyInfo) error {
	return h.MeetingServiceHandler.UpdateStatus(ctx, in, out)
}
