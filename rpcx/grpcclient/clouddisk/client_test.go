package clouddisk

import (
	"context"
	"log"
	"testing"
	"time"

	"gitlab.flyele.vip/server-side/go-zero/v2/rest/httprouter"

	"gitlab.flyele.vip/server-side/go-zero/v2/core/trace"

	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
)

func TestClientCall(t *testing.T) {
	trace.StartAgent(trace.Config{
		Name:     "xx",
		Endpoint: "https://jaeger-collector.flyele.vip/api/traces",
		Sampler:  1,
		Batcher:  "jaeger",
	})
	FileDetailHandler = func(context context.Context, request FileDetailRequest) grpcbase.Response {
		//var v = context.Value(grpcbase.RequestGrpcContext)
		//fmt.Println(v)
		//m:=grpcbase.GetGrpcContext(context)
		//m.InjectContext()
		var resp = make([]*FileDetail, 0)
		resp = append(resp, &FileDetail{
			FileID:   "xxx",
			FileName: "2222.jpg",
		})
		return grpcbase.Response{Data: FileDetailResponse{Detail: resp}}
	}
	AssociatedFilesHandler = func(context context.Context, request AssociatedFilesRequest) grpcbase.Response {
		return grpcbase.Response{
			Data: AssociatedFilesResponse{Count: 1000000},
		}
	}
	UploadFileHandler = func(ctx context.Context, req UploadFileRequest) grpcbase.Response {
		return grpcbase.Response{Message: "xxxxxx"}
	}
	go func() {
		if err := grpcbase.RegisterServer(NewBinding(NewService())); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(time.Second * 5)
	for i := 0; i < 1; i++ {
		go func(index int) {
			c := NewClient()
			//c.HandleFileDetail()
			//c, err := grpcbase.DialClient(grpcbase.ServerAddr(grpcbase.CloudDiskSVC))
			//if err != nil {
			//	log.Fatal(err)
			//}
			//client := c.(Repository)
			ctx := httprouter.NewContextData(context.Background(), &httprouter.HeaderData{RequestID: "00-9669923f709fa745488109f692a8ff11-2242da0f508476e0-01"})

			resp := c.HandleFileDetail(ctx, FileDetailRequest{
				Files:   []string{"354714419724419", "354711185653891"},
				Context: &UserContext{RequestID: "00-9669923f709fa745488109f692a8ff11-2242da0f508476e0-01"},
			})
			log.Println(resp.String())
			//
			//resp = client.AssociatedFiles(ctx, AssociatedFilesRequest{FlyeleID: "xxxxxxxxxxxx"})
			//log.Println(resp.Data)
			//
			//resp = client.UploadFile(ctx, UploadFileRequest{
			//	FilePath:      "xx",
			//	Content:       nil,
			//	OssBucketName: "",
			//})
			//log.Println(resp.Data)
		}(i)
	}
	select {}
}
