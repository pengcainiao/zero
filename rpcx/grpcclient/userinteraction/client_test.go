package userinteraction

import (
	"context"
	"log"
	"os"
	"testing"

	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
)

func TestCreateCommentInternal(t *testing.T) {
	client := newClient()
	res := client.CommonCreateComment(newContext(), CommonCreateCommentRequest{
		RefType:     "task",            // 类型 task,record
		RefID:       "617620681523324", // 关联id task_id, record_id
		MsgType:     1,                 // 消息类型,https://api.flyele.vip/markdown?filename=14、协作信息类型.md
		SystemType:  1,                 // 系统消息类型,https://api.flyele.vip/markdown?filename=14、协作信息类型.md
		Content:     "{\"content\":\"创建了事项\",\"title\":\"skr哭呢\"}",
		AffectedUID: "",                // 受影响用户id，多个用逗号隔开
		NotifyUID:   "542321150263504", // 通知用户id，多个用逗号隔开
	})

	log.Println(res)
}

func TestGetTaskNewComment(t *testing.T) {
	client := newClient()
	res := client.GetTaskNewComment(newContext(), GetTaskNewCommentRequest{
		CommentMap: nil,
		Context:    nil,
	})

	log.Println(res)
}

func newClient() Repository {
	c, err := grpcbase.DialClient(grpcbase.ServerAddr(grpcbase.UserInteractionSVC))
	if err != nil {
		log.Fatal(err)
	}
	return c.(Repository)
}

func newContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, grpcbase.RequestAuthorization, os.Getenv("TOKEN")) // nolint
	ctx = context.WithValue(ctx, grpcbase.RequestUserID, os.Getenv("USER_ID"))      // nolint
	ctx = context.WithValue(ctx, grpcbase.RequestPlatform, os.Getenv("PLATFORM"))   // nolint
	return ctx
}
