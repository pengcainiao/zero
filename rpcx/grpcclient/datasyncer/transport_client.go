// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package datasyncer

import (
	"context"
	"errors"
	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
	pb "gitlab.flyele.vip/server-side/go-zero/v2/rpcx/protos"
)

func encodeJoinedTasksRequest(ctx context.Context, req interface{}) (request interface{}, err error) {
	r := req.(JoinedTasksRequest)
	var xpaging *pb.Paging
	if r.Paging != nil {
		xpaging = &pb.Paging{
			PageNumber: r.Paging.PageNumber,
			PageRecord: r.Paging.PageRecord,
		}
	}
	var xcontext *pb.UserContext
	if r.Context != nil {
		xcontext = &pb.UserContext{
			UserID:        r.Context.UserID,
			Platform:      r.Context.Platform,
			ClientVersion: r.Context.ClientVersion,
			Token:         r.Context.Token,
			ClientIP:      r.Context.ClientIP,
			RequestID:     r.Context.RequestID,
		}
	}
	return &pb.JoinedTasksRequest{
		Keyword: r.Keyword,
		Paging:  xpaging,
		Context: xcontext,
		Type:    r.Type,
	}, nil
}

func decodeJoinedTasksResponse(ctx context.Context, resp interface{}) (response interface{}, err error) {
	r := resp.(*pb.Response)
	if r.Message != "" {
		var err = errors.New(r.Message)
		return nil, err
	}
	//在所有类型中匹配名称相同的消息名称
	var pbresp pb.JoinedTasksResponse
	if err := r.Data.UnmarshalTo(&pbresp); err != nil {
		return nil, err
	}

	return grpcbase.Response{
		Data: JoinedTasksResponse{
			RefID: pbresp.RefID,
		},
	}, nil
}
