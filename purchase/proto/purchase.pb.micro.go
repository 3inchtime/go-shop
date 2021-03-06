// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: purchase/proto/purchase.proto

package go_micro_service_purchase

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Purchase service

func NewPurchaseEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Purchase service

type PurchaseService interface {
	QueryUserPurchase(ctx context.Context, in *QueryPurchaseRequest, opts ...client.CallOption) (*QueryPurchaseResponse, error)
	QueryPurchaseDetail(ctx context.Context, in *QueryPurchaseDetailRequest, opts ...client.CallOption) (*PurchaseDetail, error)
}

type purchaseService struct {
	c    client.Client
	name string
}

func NewPurchaseService(name string, c client.Client) PurchaseService {
	return &purchaseService{
		c:    c,
		name: name,
	}
}

func (c *purchaseService) QueryUserPurchase(ctx context.Context, in *QueryPurchaseRequest, opts ...client.CallOption) (*QueryPurchaseResponse, error) {
	req := c.c.NewRequest(c.name, "Purchase.QueryUserPurchase", in)
	out := new(QueryPurchaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseService) QueryPurchaseDetail(ctx context.Context, in *QueryPurchaseDetailRequest, opts ...client.CallOption) (*PurchaseDetail, error) {
	req := c.c.NewRequest(c.name, "Purchase.QueryPurchaseDetail", in)
	out := new(PurchaseDetail)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Purchase service

type PurchaseHandler interface {
	QueryUserPurchase(context.Context, *QueryPurchaseRequest, *QueryPurchaseResponse) error
	QueryPurchaseDetail(context.Context, *QueryPurchaseDetailRequest, *PurchaseDetail) error
}

func RegisterPurchaseHandler(s server.Server, hdlr PurchaseHandler, opts ...server.HandlerOption) error {
	type purchase interface {
		QueryUserPurchase(ctx context.Context, in *QueryPurchaseRequest, out *QueryPurchaseResponse) error
		QueryPurchaseDetail(ctx context.Context, in *QueryPurchaseDetailRequest, out *PurchaseDetail) error
	}
	type Purchase struct {
		purchase
	}
	h := &purchaseHandler{hdlr}
	return s.Handle(s.NewHandler(&Purchase{h}, opts...))
}

type purchaseHandler struct {
	PurchaseHandler
}

func (h *purchaseHandler) QueryUserPurchase(ctx context.Context, in *QueryPurchaseRequest, out *QueryPurchaseResponse) error {
	return h.PurchaseHandler.QueryUserPurchase(ctx, in, out)
}

func (h *purchaseHandler) QueryPurchaseDetail(ctx context.Context, in *QueryPurchaseDetailRequest, out *PurchaseDetail) error {
	return h.PurchaseHandler.QueryPurchaseDetail(ctx, in, out)
}
