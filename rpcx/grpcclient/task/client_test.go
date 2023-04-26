package task

import (
	"context"
	"log"
	"os"
	"testing"

	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
)

func TestGetTask(t *testing.T) {
	client := newClient()
	res := client.GetTask(newContext(), GetTaskRequest{
		TaskID: "542321582801158",
	})

	log.Println(res)
}

func TestGetTasksName(t *testing.T) {
	client := newClient()
	res := client.GetTasksName(newContext(), GetTasksNameRequest{
		RefTaskID: "542321582801158",
	})

	log.Println(res)
}

func TestGetTaskDispatch(t *testing.T) {
	client := newClient()
	res := client.GetTaskDispatch(newContext(), GetTaskDispatchRequest{
		RefTaskID: "408711781417071",
	})

	log.Println(res)
}

func TestGetTaskTakers(t *testing.T) {
	client := newClient()
	res := client.GetTaskTakers(newContext(), GetTaskTakersRequest{
		TaskID:  "542427745616136",
		TakerID: "542427745878280",
	})

	log.Println(res)
}

func TestTaskExists(t *testing.T) {
	client := newClient()
	res := client.TaskExists(newContext(), TaskExistsRequest{
		TaskID: "574001421484097",
	})

	log.Println(res)
}

func TestCreateGuideTask(t *testing.T) {
	client := newClient()
	res := client.CreateGuideTask(newContext(), CreateGuideTaskRequest{
		TakerID: "551262062379281",
		Context: &UserContext{
			UserID:   "999999999",
			Platform: "wechat",
		},
	})

	log.Println(res)
}

func TestBatchQueryTask(t *testing.T) {
	client := newClient()
	res := client.BatchQueryTask(newContext(), BatchQueryTaskRequest{
		RefTaskID: []string{"542321582801158", "542706743378015"},
	})

	log.Println(res)
}

func TestGetTaskRelevantTakers(t *testing.T) {
	client := newClient()
	res := client.GetTaskRelevantTakers(newContext(), GetTaskRelevantTakersRequest{
		TaskID: "814779717124191",
		Context: &UserContext{
			UserID: os.Getenv("USER_ID"),
		},
	})

	log.Println(res)
}

func newClient() Repository {
	c, err := grpcbase.DialClient(grpcbase.ServerAddr(grpcbase.TaskSVC))
	if err != nil {
		log.Fatal(err)
	}
	return c.(Repository)
}

func newContext() context.Context {
	return context.Background()
}
