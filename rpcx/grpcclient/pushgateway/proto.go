// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//
package pushgateway

import (
	"context"

	jsoniter "github.com/json-iterator/go"
	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
)

var (
	WebsocketOnlineHandler func(ctx context.Context, req WebsocketOnlineRequest) grpcbase.Response
)

type Repository interface {
	WebsocketOnline(ctx context.Context, request WebsocketOnlineRequest) grpcbase.Response
}

type Any struct {
	type_url string `json:"type_url,omitempty"`
	value    []byte `json:"value,omitempty"`
}

func (s Any) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type Online struct {
	UserID        string `json:"user_id,omitempty"`
	Platform      string `json:"platform,omitempty"`
	ClientVersion string `json:"client_version,omitempty"`
}

func (s Online) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type Response struct {
	Message string `json:"message,omitempty"`
	Total   int32  `json:"total,omitempty"`
	Data    *Any   `json:"data,omitempty"`
}

func (s Response) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type UserContext struct {
	UserID        string `json:"user_id,omitempty"`
	Platform      string `json:"platform,omitempty"`
	ClientVersion string `json:"client_version,omitempty"`
	Token         string `json:"token,omitempty"`
	ClientIP      string `json:"client_ip,omitempty"`
	RequestID     string `json:"request_id,omitempty"`
}

func (s UserContext) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type WebsocketOnlineRequest struct {
	Users   []string     `json:"users,omitempty"`
	Context *UserContext `json:"context,omitempty"`
}

func (s WebsocketOnlineRequest) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type WebsocketOnlineResponse struct {
	Details []*Online `json:"details,omitempty"`
}

func (s WebsocketOnlineResponse) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type service struct{}

// RpcContextFromHeader 从 httprouter.Context转换为grpc中所需的用户上下文
func RpcContextFromHeader(header string) *UserContext {
	var ctx *UserContext
	_ = jsoniter.UnmarshalFromString(header, ctx)
	return ctx
}

//NewService 新建pushgateway的grpc服务
func NewService() Repository {
	return service{}
}

func (s service) WebsocketOnline(ctx context.Context, req WebsocketOnlineRequest) grpcbase.Response {
	if WebsocketOnlineHandler != nil {
		if req.Context == nil {
			panic("grpc context is nil and requestID must be set")
		}
		return grpcbase.RPCServerSideLogic("/pb.pushgateway/WebsocketOnline", req.Context.RequestID, req, func(ctx context.Context, req interface{}) grpcbase.Response {
			return WebsocketOnlineHandler(ctx, req.(WebsocketOnlineRequest))
		})
	}
	return grpcbase.NotImplErrorResponse
}
