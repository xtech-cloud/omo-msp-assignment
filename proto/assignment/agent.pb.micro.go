// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/assignment/agent.proto

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

// Client API for AgentService service

type AgentService interface {
	AddOne(ctx context.Context, in *ReqAgentAdd, opts ...client.CallOption) (*ReplyAgentOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAgentOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyAgentList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateBase(ctx context.Context, in *ReqAgentUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
}

type agentService struct {
	c    client.Client
	name string
}

func NewAgentService(name string, c client.Client) AgentService {
	return &agentService{
		c:    c,
		name: name,
	}
}

func (c *agentService) AddOne(ctx context.Context, in *ReqAgentAdd, opts ...client.CallOption) (*ReplyAgentOne, error) {
	req := c.c.NewRequest(c.name, "AgentService.AddOne", in)
	out := new(ReplyAgentOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAgentOne, error) {
	req := c.c.NewRequest(c.name, "AgentService.GetOne", in)
	out := new(ReplyAgentOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "AgentService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyAgentList, error) {
	req := c.c.NewRequest(c.name, "AgentService.GetListByFilter", in)
	out := new(ReplyAgentList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "AgentService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentService) UpdateBase(ctx context.Context, in *ReqAgentUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "AgentService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "AgentService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AgentService service

type AgentServiceHandler interface {
	AddOne(context.Context, *ReqAgentAdd, *ReplyAgentOne) error
	GetOne(context.Context, *RequestInfo, *ReplyAgentOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyAgentList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateBase(context.Context, *ReqAgentUpdate, *ReplyInfo) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
}

func RegisterAgentServiceHandler(s server.Server, hdlr AgentServiceHandler, opts ...server.HandlerOption) error {
	type agentService interface {
		AddOne(ctx context.Context, in *ReqAgentAdd, out *ReplyAgentOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyAgentOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyAgentList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateBase(ctx context.Context, in *ReqAgentUpdate, out *ReplyInfo) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
	}
	type AgentService struct {
		agentService
	}
	h := &agentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AgentService{h}, opts...))
}

type agentServiceHandler struct {
	AgentServiceHandler
}

func (h *agentServiceHandler) AddOne(ctx context.Context, in *ReqAgentAdd, out *ReplyAgentOne) error {
	return h.AgentServiceHandler.AddOne(ctx, in, out)
}

func (h *agentServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyAgentOne) error {
	return h.AgentServiceHandler.GetOne(ctx, in, out)
}

func (h *agentServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.AgentServiceHandler.RemoveOne(ctx, in, out)
}

func (h *agentServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyAgentList) error {
	return h.AgentServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *agentServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.AgentServiceHandler.GetStatistic(ctx, in, out)
}

func (h *agentServiceHandler) UpdateBase(ctx context.Context, in *ReqAgentUpdate, out *ReplyInfo) error {
	return h.AgentServiceHandler.UpdateBase(ctx, in, out)
}

func (h *agentServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.AgentServiceHandler.UpdateByFilter(ctx, in, out)
}
