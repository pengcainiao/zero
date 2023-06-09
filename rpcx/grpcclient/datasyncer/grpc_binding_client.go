// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package datasyncer

import (
	"context"
	"github.com/pengcainiao/zero/rpcx/grpcbase"
	"log"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
)

func init() {
	grpcbase.RegisterClients(grpcbase.ServerAddr(grpcbase.DataSyncerSVC), &clientBinding{})
}

type clientBinding struct {
	joinedTasks endpoint.Endpoint
}

func (c *clientBinding) JoinedTasks(ctx context.Context, params JoinedTasksRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：JoinedTasks request context is nil，trace span将无法生效")
	}
	response, err := c.joinedTasks(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}

func (c *clientBinding) GRPCClient(cc *grpc.ClientConn) interface{} {
	c.newClient(cc)
	return c
}

func (c *clientBinding) newClient(cc *grpc.ClientConn) {

	c.joinedTasks = grpcbase.CreateGRPCClientEndpoint(cc, "pb.DataSyncer",
		"JoinedTasks",
		encodeJoinedTasksRequest,
		decodeJoinedTasksResponse)
}
