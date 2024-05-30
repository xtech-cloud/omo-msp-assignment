// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/assignment/category.proto

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

// Client API for CategoryService service

type CategoryService interface {
	AddOne(ctx context.Context, in *ReqCategoryAdd, opts ...client.CallOption) (*ReplyCategoryOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyCategoryOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyCategoryList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateBase(ctx context.Context, in *ReqCategoryUpdate, opts ...client.CallOption) (*ReplyCategoryOne, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
}

type categoryService struct {
	c    client.Client
	name string
}

func NewCategoryService(name string, c client.Client) CategoryService {
	return &categoryService{
		c:    c,
		name: name,
	}
}

func (c *categoryService) AddOne(ctx context.Context, in *ReqCategoryAdd, opts ...client.CallOption) (*ReplyCategoryOne, error) {
	req := c.c.NewRequest(c.name, "CategoryService.AddOne", in)
	out := new(ReplyCategoryOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyCategoryOne, error) {
	req := c.c.NewRequest(c.name, "CategoryService.GetOne", in)
	out := new(ReplyCategoryOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "CategoryService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyCategoryList, error) {
	req := c.c.NewRequest(c.name, "CategoryService.GetListByFilter", in)
	out := new(ReplyCategoryList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "CategoryService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) UpdateBase(ctx context.Context, in *ReqCategoryUpdate, opts ...client.CallOption) (*ReplyCategoryOne, error) {
	req := c.c.NewRequest(c.name, "CategoryService.UpdateBase", in)
	out := new(ReplyCategoryOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "CategoryService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CategoryService service

type CategoryServiceHandler interface {
	AddOne(context.Context, *ReqCategoryAdd, *ReplyCategoryOne) error
	GetOne(context.Context, *RequestInfo, *ReplyCategoryOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyCategoryList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateBase(context.Context, *ReqCategoryUpdate, *ReplyCategoryOne) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
}

func RegisterCategoryServiceHandler(s server.Server, hdlr CategoryServiceHandler, opts ...server.HandlerOption) error {
	type categoryService interface {
		AddOne(ctx context.Context, in *ReqCategoryAdd, out *ReplyCategoryOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyCategoryOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyCategoryList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateBase(ctx context.Context, in *ReqCategoryUpdate, out *ReplyCategoryOne) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
	}
	type CategoryService struct {
		categoryService
	}
	h := &categoryServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CategoryService{h}, opts...))
}

type categoryServiceHandler struct {
	CategoryServiceHandler
}

func (h *categoryServiceHandler) AddOne(ctx context.Context, in *ReqCategoryAdd, out *ReplyCategoryOne) error {
	return h.CategoryServiceHandler.AddOne(ctx, in, out)
}

func (h *categoryServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyCategoryOne) error {
	return h.CategoryServiceHandler.GetOne(ctx, in, out)
}

func (h *categoryServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.CategoryServiceHandler.RemoveOne(ctx, in, out)
}

func (h *categoryServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyCategoryList) error {
	return h.CategoryServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *categoryServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.CategoryServiceHandler.GetStatistic(ctx, in, out)
}

func (h *categoryServiceHandler) UpdateBase(ctx context.Context, in *ReqCategoryUpdate, out *ReplyCategoryOne) error {
	return h.CategoryServiceHandler.UpdateBase(ctx, in, out)
}

func (h *categoryServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.CategoryServiceHandler.UpdateByFilter(ctx, in, out)
}