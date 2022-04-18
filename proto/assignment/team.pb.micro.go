// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/assignment/team.proto

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

// Client API for TeamService service

type TeamService interface {
	AddOne(ctx context.Context, in *ReqTeamAdd, opts ...client.CallOption) (*ReplyTeamInfo, error)
	UpdateBase(ctx context.Context, in *ReqTeamUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTeamInfo, error)
	Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTeamList, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTeamList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateStatus(ctx context.Context, in *RequestIntFlag, opts ...client.CallOption) (*ReplyInfo, error)
	AppendMember(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error)
	SubtractMember(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error)
}

type teamService struct {
	c    client.Client
	name string
}

func NewTeamService(name string, c client.Client) TeamService {
	return &teamService{
		c:    c,
		name: name,
	}
}

func (c *teamService) AddOne(ctx context.Context, in *ReqTeamAdd, opts ...client.CallOption) (*ReplyTeamInfo, error) {
	req := c.c.NewRequest(c.name, "TeamService.AddOne", in)
	out := new(ReplyTeamInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) UpdateBase(ctx context.Context, in *ReqTeamUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TeamService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TeamService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTeamInfo, error) {
	req := c.c.NewRequest(c.name, "TeamService.GetOne", in)
	out := new(ReplyTeamInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTeamList, error) {
	req := c.c.NewRequest(c.name, "TeamService.Search", in)
	out := new(ReplyTeamList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTeamList, error) {
	req := c.c.NewRequest(c.name, "TeamService.GetListByFilter", in)
	out := new(ReplyTeamList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "TeamService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TeamService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) UpdateStatus(ctx context.Context, in *RequestIntFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TeamService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) AppendMember(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "TeamService.AppendMember", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamService) SubtractMember(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "TeamService.SubtractMember", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TeamService service

type TeamServiceHandler interface {
	AddOne(context.Context, *ReqTeamAdd, *ReplyTeamInfo) error
	UpdateBase(context.Context, *ReqTeamUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyTeamInfo) error
	Search(context.Context, *RequestInfo, *ReplyTeamList) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyTeamList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	UpdateStatus(context.Context, *RequestIntFlag, *ReplyInfo) error
	AppendMember(context.Context, *RequestList, *ReplyList) error
	SubtractMember(context.Context, *RequestList, *ReplyList) error
}

func RegisterTeamServiceHandler(s server.Server, hdlr TeamServiceHandler, opts ...server.HandlerOption) error {
	type teamService interface {
		AddOne(ctx context.Context, in *ReqTeamAdd, out *ReplyTeamInfo) error
		UpdateBase(ctx context.Context, in *ReqTeamUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyTeamInfo) error
		Search(ctx context.Context, in *RequestInfo, out *ReplyTeamList) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyTeamList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		UpdateStatus(ctx context.Context, in *RequestIntFlag, out *ReplyInfo) error
		AppendMember(ctx context.Context, in *RequestList, out *ReplyList) error
		SubtractMember(ctx context.Context, in *RequestList, out *ReplyList) error
	}
	type TeamService struct {
		teamService
	}
	h := &teamServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TeamService{h}, opts...))
}

type teamServiceHandler struct {
	TeamServiceHandler
}

func (h *teamServiceHandler) AddOne(ctx context.Context, in *ReqTeamAdd, out *ReplyTeamInfo) error {
	return h.TeamServiceHandler.AddOne(ctx, in, out)
}

func (h *teamServiceHandler) UpdateBase(ctx context.Context, in *ReqTeamUpdate, out *ReplyInfo) error {
	return h.TeamServiceHandler.UpdateBase(ctx, in, out)
}

func (h *teamServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.TeamServiceHandler.RemoveOne(ctx, in, out)
}

func (h *teamServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyTeamInfo) error {
	return h.TeamServiceHandler.GetOne(ctx, in, out)
}

func (h *teamServiceHandler) Search(ctx context.Context, in *RequestInfo, out *ReplyTeamList) error {
	return h.TeamServiceHandler.Search(ctx, in, out)
}

func (h *teamServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyTeamList) error {
	return h.TeamServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *teamServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.TeamServiceHandler.GetStatistic(ctx, in, out)
}

func (h *teamServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.TeamServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *teamServiceHandler) UpdateStatus(ctx context.Context, in *RequestIntFlag, out *ReplyInfo) error {
	return h.TeamServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *teamServiceHandler) AppendMember(ctx context.Context, in *RequestList, out *ReplyList) error {
	return h.TeamServiceHandler.AppendMember(ctx, in, out)
}

func (h *teamServiceHandler) SubtractMember(ctx context.Context, in *RequestList, out *ReplyList) error {
	return h.TeamServiceHandler.SubtractMember(ctx, in, out)
}
