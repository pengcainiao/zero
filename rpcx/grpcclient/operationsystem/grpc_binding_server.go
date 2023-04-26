// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package operationsystem

import (
	"context"
	"log"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
	pb "gitlab.flyele.vip/server-side/go-zero/v2/rpcx/protos"
	"google.golang.org/grpc"
)

type serverBinding struct {
	sendOperationForLoginUser grpctransport.Handler
}

func (b *serverBinding) SendOperationForLoginUser(ctx context.Context, req *pb.SendOperationRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：SendOperationForLoginUser receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.sendOperationForLoginUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

func (b *serverBinding) RegisterServer(srv *grpc.Server) error {
	pb.RegisterOperationSystemServer(srv, b)
	return nil
}

func (b *serverBinding) GRPCHandler() map[string]grpctransport.Handler {
	return map[string]grpctransport.Handler{
		"sendOperationForLoginUser": b.sendOperationForLoginUser,
	}
}

func NewBinding(svc Repository) *serverBinding {
	return &serverBinding{
		sendOperationForLoginUser: grpcbase.CreateGRPCServer(
			makeSendOperationForLoginUserEndpoint(svc),
			decodeSendOperationForLoginUserRequest,
			encodeSendOperationForLoginUserResponse,
		),
	}
}
