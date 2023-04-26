package cronjob

import (
	"context"
	"log"
	"testing"
	"time"

	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
)

func TestClientCall(t *testing.T) {
	CreateCronJobHandler = func(ctx context.Context, request CreateCronJobRequest) grpcbase.Response {
		return grpcbase.Response{
			Data: CreateCronJobResponse{JobID: "xxxx"},
		}
	}
	go func() {
		if err := grpcbase.RegisterServer(NewBinding(NewService())); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(time.Second * 5)

	c, err := grpcbase.DialClient(grpcbase.ServerAddr(grpcbase.CronJobSVC))
	if err != nil {
		log.Fatal(err)
	}
	client := c.(Repository)
	ctx := context.Background()

	//resp := client.FileDetail(ctx, FileDetailRequest{Files: []string{"354714419724419", "354711185653891"}})
	//log.Println(resp.Data)

	resp := client.CreateCronJob(ctx, CreateCronJobRequest{
		Schedule:        "",
		Code:            0,
		RefTaskID:       "",
		ExecuteMetadata: nil,
		Context:         nil,
		//Context: &UserContext{
		//	UserID:        "xxxx",
		//	Platform:      "xxx",
		//	ClientVersion: "",
		//	Token:         "",
		//	ClientIP:      "",
		//	TraceID:       "",
		//},
	})
	log.Println(resp.Data)

}
