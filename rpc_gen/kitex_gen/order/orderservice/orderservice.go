// Code generated by Kitex v0.9.1. DO NOT EDIT.

package orderservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	order "github.com/py/biz-demo/gomall/rpc_gen/kitex_gen/order"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"PlaceOrder": kitex.NewMethodInfo(
		placeOrderHandler,
		newPlaceOrderArgs,
		newPlaceOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ListOrder": kitex.NewMethodInfo(
		listOrderHandler,
		newListOrderArgs,
		newListOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	orderServiceServiceInfo                = NewServiceInfo()
	orderServiceServiceInfoForClient       = NewServiceInfoForClient()
	orderServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "OrderService"
	handlerType := (*order.OrderService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "order",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func placeOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.PlaceOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).PlaceOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PlaceOrderArgs:
		success, err := handler.(order.OrderService).PlaceOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PlaceOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPlaceOrderArgs() interface{} {
	return &PlaceOrderArgs{}
}

func newPlaceOrderResult() interface{} {
	return &PlaceOrderResult{}
}

type PlaceOrderArgs struct {
	Req *order.PlaceOrderReq
}

func (p *PlaceOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.PlaceOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PlaceOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PlaceOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PlaceOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PlaceOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PlaceOrderArgs_Req_DEFAULT *order.PlaceOrderReq

func (p *PlaceOrderArgs) GetReq() *order.PlaceOrderReq {
	if !p.IsSetReq() {
		return PlaceOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PlaceOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PlaceOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PlaceOrderResult struct {
	Success *order.PlaceOrderResp
}

var PlaceOrderResult_Success_DEFAULT *order.PlaceOrderResp

func (p *PlaceOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.PlaceOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PlaceOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PlaceOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PlaceOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PlaceOrderResult) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PlaceOrderResult) GetSuccess() *order.PlaceOrderResp {
	if !p.IsSetSuccess() {
		return PlaceOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PlaceOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.PlaceOrderResp)
}

func (p *PlaceOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PlaceOrderResult) GetResult() interface{} {
	return p.Success
}

func listOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.ListOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).ListOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListOrderArgs:
		success, err := handler.(order.OrderService).ListOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListOrderArgs() interface{} {
	return &ListOrderArgs{}
}

func newListOrderResult() interface{} {
	return &ListOrderResult{}
}

type ListOrderArgs struct {
	Req *order.ListOrderReq
}

func (p *ListOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.ListOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.ListOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListOrderArgs_Req_DEFAULT *order.ListOrderReq

func (p *ListOrderArgs) GetReq() *order.ListOrderReq {
	if !p.IsSetReq() {
		return ListOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListOrderResult struct {
	Success *order.ListOrderResp
}

var ListOrderResult_Success_DEFAULT *order.ListOrderResp

func (p *ListOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.ListOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListOrderResult) Unmarshal(in []byte) error {
	msg := new(order.ListOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListOrderResult) GetSuccess() *order.ListOrderResp {
	if !p.IsSetSuccess() {
		return ListOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.ListOrderResp)
}

func (p *ListOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListOrderResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PlaceOrder(ctx context.Context, Req *order.PlaceOrderReq) (r *order.PlaceOrderResp, err error) {
	var _args PlaceOrderArgs
	_args.Req = Req
	var _result PlaceOrderResult
	if err = p.c.Call(ctx, "PlaceOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListOrder(ctx context.Context, Req *order.ListOrderReq) (r *order.ListOrderResp, err error) {
	var _args ListOrderArgs
	_args.Req = Req
	var _result ListOrderResult
	if err = p.c.Call(ctx, "ListOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
