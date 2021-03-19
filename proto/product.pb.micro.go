// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/product.proto

package product

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

// Api Endpoints for Product service

func NewProductEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Product service

type ProductService interface {
	AddProduct(ctx context.Context, in *Request_ProductInfo, opts ...client.CallOption) (*Response_Product, error)
	DelProduct(ctx context.Context, in *Request_ProductID, opts ...client.CallOption) (*Response_DelProduct, error)
	ChangeProduct(ctx context.Context, in *Request_ProductInfo, opts ...client.CallOption) (*Response_Product, error)
	FindProductByID(ctx context.Context, in *Request_ProductID, opts ...client.CallOption) (*Response_ProductInfo, error)
	FindProductByRFID(ctx context.Context, in *Request_ProductRFID, opts ...client.CallOption) (*Response_ProductInfo, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) AddProduct(ctx context.Context, in *Request_ProductInfo, opts ...client.CallOption) (*Response_Product, error) {
	req := c.c.NewRequest(c.name, "Product.AddProduct", in)
	out := new(Response_Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) DelProduct(ctx context.Context, in *Request_ProductID, opts ...client.CallOption) (*Response_DelProduct, error) {
	req := c.c.NewRequest(c.name, "Product.DelProduct", in)
	out := new(Response_DelProduct)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) ChangeProduct(ctx context.Context, in *Request_ProductInfo, opts ...client.CallOption) (*Response_Product, error) {
	req := c.c.NewRequest(c.name, "Product.ChangeProduct", in)
	out := new(Response_Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) FindProductByID(ctx context.Context, in *Request_ProductID, opts ...client.CallOption) (*Response_ProductInfo, error) {
	req := c.c.NewRequest(c.name, "Product.FindProductByID", in)
	out := new(Response_ProductInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) FindProductByRFID(ctx context.Context, in *Request_ProductRFID, opts ...client.CallOption) (*Response_ProductInfo, error) {
	req := c.c.NewRequest(c.name, "Product.FindProductByRFID", in)
	out := new(Response_ProductInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Product service

type ProductHandler interface {
	AddProduct(context.Context, *Request_ProductInfo, *Response_Product) error
	DelProduct(context.Context, *Request_ProductID, *Response_DelProduct) error
	ChangeProduct(context.Context, *Request_ProductInfo, *Response_Product) error
	FindProductByID(context.Context, *Request_ProductID, *Response_ProductInfo) error
	FindProductByRFID(context.Context, *Request_ProductRFID, *Response_ProductInfo) error
}

func RegisterProductHandler(s server.Server, hdlr ProductHandler, opts ...server.HandlerOption) error {
	type product interface {
		AddProduct(ctx context.Context, in *Request_ProductInfo, out *Response_Product) error
		DelProduct(ctx context.Context, in *Request_ProductID, out *Response_DelProduct) error
		ChangeProduct(ctx context.Context, in *Request_ProductInfo, out *Response_Product) error
		FindProductByID(ctx context.Context, in *Request_ProductID, out *Response_ProductInfo) error
		FindProductByRFID(ctx context.Context, in *Request_ProductRFID, out *Response_ProductInfo) error
	}
	type Product struct {
		product
	}
	h := &productHandler{hdlr}
	return s.Handle(s.NewHandler(&Product{h}, opts...))
}

type productHandler struct {
	ProductHandler
}

func (h *productHandler) AddProduct(ctx context.Context, in *Request_ProductInfo, out *Response_Product) error {
	return h.ProductHandler.AddProduct(ctx, in, out)
}

func (h *productHandler) DelProduct(ctx context.Context, in *Request_ProductID, out *Response_DelProduct) error {
	return h.ProductHandler.DelProduct(ctx, in, out)
}

func (h *productHandler) ChangeProduct(ctx context.Context, in *Request_ProductInfo, out *Response_Product) error {
	return h.ProductHandler.ChangeProduct(ctx, in, out)
}

func (h *productHandler) FindProductByID(ctx context.Context, in *Request_ProductID, out *Response_ProductInfo) error {
	return h.ProductHandler.FindProductByID(ctx, in, out)
}

func (h *productHandler) FindProductByRFID(ctx context.Context, in *Request_ProductRFID, out *Response_ProductInfo) error {
	return h.ProductHandler.FindProductByRFID(ctx, in, out)
}
