package record

import (
	"context"
	"log"
	"os"
	"testing"

	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
)

func TestGetRecord(t *testing.T) {
	client := newClient()
	res := client.GetRecord(newContext(), GetRecordRequest{
		RecordID: "717050075938914",
		Context: &UserContext{
			UserID: os.Getenv("USER_ID"),
		},
	})

	log.Println(res)
}

func TestGetRecordDetails(t *testing.T) {
	client := newClient()
	res := client.GetRecordDetails(newContext(), GetRecordDetailsRequest{
		RecordID: "542678482419980",
		Context: &UserContext{
			UserID: os.Getenv("USER_ID"),
		},
	})

	log.Println(res)
}

func TestGetRelationRecordTotal(t *testing.T) {
	client := newClient()
	res := client.GetRelationRecordTotal(newContext(), GetRelationRecordTotalRequest{
		RefDispatchsID: []string{"408711781418095", "411528000177381", "414095487272120", "427242219898017"},
	})

	log.Println(res)
}

func TestGetRecordTakers(t *testing.T) {
	client := newClient()
	res := client.GetRecordTakers(newContext(), GetRecordTakersRequest{
		RecordID: "576265839706341",
		Context: &UserContext{
			UserID: os.Getenv("USER_ID"),
		},
	})

	log.Println(res)
}

func TestRecordExists(t *testing.T) {
	client := newClient()
	res := client.RecordExists(newContext(), RecordExistsRequest{
		RecordID: "576265839706341",
	})

	log.Println(res)
}

func newClient() Repository {
	c, err := grpcbase.DialClient(grpcbase.ServerAddr(grpcbase.RecordSVC))
	if err != nil {
		log.Fatal(err)
	}
	return c.(Repository)
}

func newContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Authorization", os.Getenv("TOKEN"))      // nolint
	ctx = context.WithValue(ctx, "X-Auth-User", os.Getenv("USER_ID"))      // nolint
	ctx = context.WithValue(ctx, "X-Auth-Platform", os.Getenv("PLATFORM")) // nolint
	return ctx
}
