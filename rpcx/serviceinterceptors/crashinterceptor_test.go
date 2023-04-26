package serverinterceptors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.flyele.vip/server-side/go-zero/v2/core/logx"
	"google.golang.org/grpc"
)

func init() {
	logx.Disable()
}

func TestStreamCrashInterceptor(t *testing.T) {
	err := StreamCrashInterceptor(nil, nil, nil, func(
		srv interface{}, stream grpc.ServerStream) error {
		panic("mock panic")
	})
	assert.NotNil(t, err)
}

func TestUnaryCrashInterceptor(t *testing.T) {
	_, err := UnaryCrashInterceptor(context.Background(), nil, nil,
		func(ctx context.Context, req interface{}) (interface{}, error) {
			panic("mock panic")
		})
	assert.NotNil(t, err)
}
