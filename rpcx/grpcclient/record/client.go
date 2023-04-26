// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//
package record

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log"

	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
)

//GrpcClient grpc客户端，各服务无需重复开发客户端。
// 如果对response有特殊处理，请在该文件夹下新增 response.go 处理后续逻辑
type GrpcClient struct {
	client Repository
}

//NewClient record GRPC客户端实例
func NewClient() GrpcClient {
	c, err := grpcbase.DialClient(grpcbase.ServerAddr(grpcbase.RecordSVC))
	if err != nil {
		log.Fatal(err)
	}
	client := c.(Repository)
	return GrpcClient{client}
}

// HandleGetRecord 客户端处理逻辑
func (c *GrpcClient) HandleGetRecord(ctx context.Context, req GetRecordRequest) *GetRecordResponse {
	if req.Context == nil {
		panic("grpc context is nil and requestID must be set")
	}
	ctx1, span, requestID := grpcbase.NewTraceSpanFromRequestID("client", "/pb.record/GetRecord", req.Context.RequestID)
	defer span.End()

	span.SetAttributes(attribute.String("message.id", req.Context.RequestID))
	span.SetAttributes(attribute.String("user.id", req.Context.UserID))

	span.AddEvent("request", trace.WithAttributes(attribute.String("message.data", req.String())))
	req.Context.RequestID = requestID

	resp := c.client.GetRecord(ctx1, req)
	if resp.Message != "" {
		log.Println("ERROR：请求 GetRecord 失败，error：" + resp.Message)
		return nil
	}
	if resp.Data != nil {
		var m = resp.Data.(GetRecordResponse)
		return &m
	}
	return nil
}

// HandleGetRecordDetails 客户端处理逻辑
func (c *GrpcClient) HandleGetRecordDetails(ctx context.Context, req GetRecordDetailsRequest) *GetRecordDetailsResponse {
	if req.Context == nil {
		panic("grpc context is nil and requestID must be set")
	}
	ctx1, span, requestID := grpcbase.NewTraceSpanFromRequestID("client", "/pb.record/GetRecordDetails", req.Context.RequestID)
	defer span.End()

	span.SetAttributes(attribute.String("message.id", req.Context.RequestID))
	span.SetAttributes(attribute.String("user.id", req.Context.UserID))

	span.AddEvent("request", trace.WithAttributes(attribute.String("message.data", req.String())))
	req.Context.RequestID = requestID

	resp := c.client.GetRecordDetails(ctx1, req)
	if resp.Message != "" {
		log.Println("ERROR：请求 GetRecordDetails 失败，error：" + resp.Message)
		return nil
	}
	if resp.Data != nil {
		var m = resp.Data.(GetRecordDetailsResponse)
		return &m
	}
	return nil
}

// HandleGetRelationRecordTotal 客户端处理逻辑
func (c *GrpcClient) HandleGetRelationRecordTotal(ctx context.Context, req GetRelationRecordTotalRequest) *GetRelationRecordTotalResponse {
	if req.Context == nil {
		panic("grpc context is nil and requestID must be set")
	}
	ctx1, span, requestID := grpcbase.NewTraceSpanFromRequestID("client", "/pb.record/GetRelationRecordTotal", req.Context.RequestID)
	defer span.End()

	span.SetAttributes(attribute.String("message.id", req.Context.RequestID))
	span.SetAttributes(attribute.String("user.id", req.Context.UserID))

	span.AddEvent("request", trace.WithAttributes(attribute.String("message.data", req.String())))
	req.Context.RequestID = requestID

	resp := c.client.GetRelationRecordTotal(ctx1, req)
	if resp.Message != "" {
		log.Println("ERROR：请求 GetRelationRecordTotal 失败，error：" + resp.Message)
		return nil
	}
	if resp.Data != nil {
		var m = resp.Data.(GetRelationRecordTotalResponse)
		return &m
	}
	return nil
}

// HandleGetRecordTakers 客户端处理逻辑
func (c *GrpcClient) HandleGetRecordTakers(ctx context.Context, req GetRecordTakersRequest) *GetRecordTakersResponse {
	if req.Context == nil {
		panic("grpc context is nil and requestID must be set")
	}
	ctx1, span, requestID := grpcbase.NewTraceSpanFromRequestID("client", "/pb.record/GetRecordTakers", req.Context.RequestID)
	defer span.End()

	span.SetAttributes(attribute.String("message.id", req.Context.RequestID))
	span.SetAttributes(attribute.String("user.id", req.Context.UserID))

	span.AddEvent("request", trace.WithAttributes(attribute.String("message.data", req.String())))
	req.Context.RequestID = requestID

	resp := c.client.GetRecordTakers(ctx1, req)
	if resp.Message != "" {
		log.Println("ERROR：请求 GetRecordTakers 失败，error：" + resp.Message)
		return nil
	}
	if resp.Data != nil {
		var m = resp.Data.(GetRecordTakersResponse)
		return &m
	}
	return nil
}

// HandleRecordExists 客户端处理逻辑
func (c *GrpcClient) HandleRecordExists(ctx context.Context, req RecordExistsRequest) *RecordExistsResponse {
	if req.Context == nil {
		panic("grpc context is nil and requestID must be set")
	}
	ctx1, span, requestID := grpcbase.NewTraceSpanFromRequestID("client", "/pb.record/RecordExists", req.Context.RequestID)
	defer span.End()

	span.SetAttributes(attribute.String("message.id", req.Context.RequestID))
	span.SetAttributes(attribute.String("user.id", req.Context.UserID))

	span.AddEvent("request", trace.WithAttributes(attribute.String("message.data", req.String())))
	req.Context.RequestID = requestID

	resp := c.client.RecordExists(ctx1, req)
	if resp.Message != "" {
		log.Println("ERROR：请求 RecordExists 失败，error：" + resp.Message)
		return nil
	}
	if resp.Data != nil {
		var m = resp.Data.(RecordExistsResponse)
		return &m
	}
	return nil
}

// HandleBatchUnBindTask 客户端处理逻辑
func (c *GrpcClient) HandleBatchUnBindTask(ctx context.Context, req BatchUnBindTaskRequest) error {
	resp := c.client.BatchUnBindTask(ctx, req)
	if resp.Message != "" {
		return errors.New(resp.Message)
	}
	return nil
}

// HandleIncrRecordCommentNum 客户端处理逻辑
func (c *GrpcClient) HandleIncrRecordCommentNum(ctx context.Context, req IncrRecordCommentNumRequest) error {
	resp := c.client.IncrRecordCommentNum(ctx, req)
	if resp.Message != "" {
		return errors.New(resp.Message)
	}
	return nil
}

// HandleGenerateRecordForNewUsers 客户端处理逻辑
func (c *GrpcClient) HandleGenerateRecordForNewUsers(ctx context.Context, req GenerateRecordForNewUsersRequest) error {
	resp := c.client.GenerateRecordForNewUsers(ctx, req)
	if resp.Message != "" {
		return errors.New(resp.Message)
	}
	return nil
}
