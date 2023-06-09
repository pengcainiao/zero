// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package datasyncer

import (
	"context"
	"github.com/pengcainiao/zero/rpcx/grpcbase"
	pb "github.com/pengcainiao/zero/rpcx/protos"
	"log"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	"google.golang.org/grpc"
)

type serverBinding struct {
	joinedTasks grpctransport.Handler
}

func (b *serverBinding) JoinedTasks(ctx context.Context, req *pb.JoinedTasksRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：JoinedTasks receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.joinedTasks.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

func (b *serverBinding) RegisterServer(srv *grpc.Server) error {
	pb.RegisterDataSyncerServer(srv, b)
	return nil
}

func (b *serverBinding) GRPCHandler() map[string]grpctransport.Handler {
	return map[string]grpctransport.Handler{
		"joinedTasks": b.joinedTasks,
	}
}

func NewBinding(svc Repository) *serverBinding {
	return &serverBinding{
		joinedTasks: grpcbase.CreateGRPCServer(
			makeJoinedTasksEndpoint(svc),
			decodeJoinedTasksRequest,
			encodeJoinedTasksResponse,
		),
	}
}
