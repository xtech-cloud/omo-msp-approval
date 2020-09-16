// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/approval/task.proto

package approval

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

// Client API for Task service

type TaskService interface {
	// 提交一个任务
	Submit(ctx context.Context, in *TaskSubmitRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 通过一个任务
	Accept(ctx context.Context, in *TaskAcceptRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 驳回一个任务
	Reject(ctx context.Context, in *TaskRejectRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 获取任务
	Get(ctx context.Context, in *TaskGetRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 列举任务
	List(ctx context.Context, in *TaskGetRequest, opts ...client.CallOption) (*BlankResponse, error)
}

type taskService struct {
	c    client.Client
	name string
}

func NewTaskService(name string, c client.Client) TaskService {
	return &taskService{
		c:    c,
		name: name,
	}
}

func (c *taskService) Submit(ctx context.Context, in *TaskSubmitRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Task.Submit", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Accept(ctx context.Context, in *TaskAcceptRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Task.Accept", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Reject(ctx context.Context, in *TaskRejectRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Task.Reject", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Get(ctx context.Context, in *TaskGetRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Task.Get", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) List(ctx context.Context, in *TaskGetRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Task.List", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Task service

type TaskHandler interface {
	// 提交一个任务
	Submit(context.Context, *TaskSubmitRequest, *BlankResponse) error
	// 通过一个任务
	Accept(context.Context, *TaskAcceptRequest, *BlankResponse) error
	// 驳回一个任务
	Reject(context.Context, *TaskRejectRequest, *BlankResponse) error
	// 获取任务
	Get(context.Context, *TaskGetRequest, *BlankResponse) error
	// 列举任务
	List(context.Context, *TaskGetRequest, *BlankResponse) error
}

func RegisterTaskHandler(s server.Server, hdlr TaskHandler, opts ...server.HandlerOption) error {
	type task interface {
		Submit(ctx context.Context, in *TaskSubmitRequest, out *BlankResponse) error
		Accept(ctx context.Context, in *TaskAcceptRequest, out *BlankResponse) error
		Reject(ctx context.Context, in *TaskRejectRequest, out *BlankResponse) error
		Get(ctx context.Context, in *TaskGetRequest, out *BlankResponse) error
		List(ctx context.Context, in *TaskGetRequest, out *BlankResponse) error
	}
	type Task struct {
		task
	}
	h := &taskHandler{hdlr}
	return s.Handle(s.NewHandler(&Task{h}, opts...))
}

type taskHandler struct {
	TaskHandler
}

func (h *taskHandler) Submit(ctx context.Context, in *TaskSubmitRequest, out *BlankResponse) error {
	return h.TaskHandler.Submit(ctx, in, out)
}

func (h *taskHandler) Accept(ctx context.Context, in *TaskAcceptRequest, out *BlankResponse) error {
	return h.TaskHandler.Accept(ctx, in, out)
}

func (h *taskHandler) Reject(ctx context.Context, in *TaskRejectRequest, out *BlankResponse) error {
	return h.TaskHandler.Reject(ctx, in, out)
}

func (h *taskHandler) Get(ctx context.Context, in *TaskGetRequest, out *BlankResponse) error {
	return h.TaskHandler.Get(ctx, in, out)
}

func (h *taskHandler) List(ctx context.Context, in *TaskGetRequest, out *BlankResponse) error {
	return h.TaskHandler.List(ctx, in, out)
}
