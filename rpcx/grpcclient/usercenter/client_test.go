package usercenter

import (
	"context"
	"log"
	"testing"
	"time"

	"gitlab.flyele.vip/server-side/go-zero/v2/rpcx/grpcbase"
)

func TestClientCall(t *testing.T) {
	BatchQueryUsersHandler = func(ctx context.Context, request BatchQueryUsersRequest) grpcbase.Response {
		return grpcbase.Response{
			Data: BatchQueryUsersResponse{
				Response: []*SimpleUserInfo{{
					Avatar:   "fafsdaf",
					NickName: "xxxx",
					Sex:      "nv",
					Pinyin:   "xxxx",
					UserID:   "408635168784547"}},
			},
		}
	}
	//SetUserSettingsHandler = func(ctx context.Context, req SetUserSettingsRequest) grpcbase.Response {
	//	return grpcbase.Response{}
	//}
	UpdateUserHandler = func(ctx context.Context, req UpdateUserRequest) grpcbase.Response {
		return grpcbase.Response{}
	}
	CreateUserByOfficialAccountHandler = func(ctx context.Context, req CreateUserByOfficialAccountRequest) grpcbase.Response {
		return grpcbase.Response{}
	}
	TryGrantAccessTokenHandler = func(ctx context.Context, req TryGrantAccessTokenRequest) grpcbase.Response {
		return grpcbase.Response{Data: TryGrantAccessTokenResponse{
			UserID: "xxx",
			Token:  "121212-xxxx",
		}}
	}
	SetUserLoginStateHandler = func(ctx context.Context, req SetUserLoginStateRequest) grpcbase.Response {
		t.Log(req)
		return grpcbase.Response{Data: TryGrantAccessTokenResponse{
			UserID: "xxx",
			Token:  "121212-xxxx",
		}}
	}
	go func() {
		if err := grpcbase.RegisterServer(NewBinding(NewService())); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(time.Second * 15)

	c, err := grpcbase.DialClient(grpcbase.ServerAddr(grpcbase.UserCenterSVC))
	if err != nil {
		log.Fatal(err)
	}
	client := c.(Repository)
	ctx := context.Background()

	//resp := client.FileDetail(ctx, FileDetailRequest{Files: []string{"354714419724419", "354711185653891"}})
	//log.Println(resp.Data)

	resp := client.BatchQueryUsers(ctx, BatchQueryUsersRequest{UserID: []string{"408635168784547"}})
	log.Println(resp.Data)
	resp = client.TryGrantAccessToken(ctx, TryGrantAccessTokenRequest{UnionID: "xxxxxx"})
	log.Println(resp.Data)

	//resp = client.SetUserSettings(ctx, SetUserSettingsRequest{UserID: "xxx", Data: map[string]string{}})
	//log.Println(resp.Data)

	resp = client.UpdateUser(ctx, UpdateUserRequest{UnionID: "xxx"})
	log.Println(resp.Data)
	resp = client.CreateUserByOfficialAccount(ctx, CreateUserByOfficialAccountRequest{NickName: "名字"})
	log.Println(resp.Data)
	resp = client.SetUserLoginState(ctx, SetUserLoginStateRequest{
		UserID:   "xx",
		DeviceID: "xxx",
		Platform: "pc",
		Online:   1,
	})
	log.Println(resp.Data)

}

func TestServerRun(t *testing.T) {
	SetUserLoginStateHandler = func(ctx context.Context, req SetUserLoginStateRequest) grpcbase.Response {
		t.Log(req)
		return grpcbase.Response{Data: TryGrantAccessTokenResponse{
			UserID: "xxx",
			Token:  "121212-xxxx",
		}}
	}
	if err := grpcbase.RegisterServer(NewBinding(NewService())); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Minute)
}

func TestSetUserLoginState(t *testing.T) {
	c, err := grpcbase.DialClient(grpcbase.ServerAddr(grpcbase.UserCenterSVC))
	if err != nil {
		log.Fatal(err)
	}
	client := c.(Repository)
	ctx := context.Background()
	resp := client.BatchGetUserAccount(ctx, BatchGetUserAccountRequest{
		UserAccountInfoRequests: []*UserAccountInfoRequest{
			{
				Account:         "onOcf5DDihXOcMIgGsbPnBjhPlGM",
				AccountProvider: "wechat-openid",
			},
			{
				OfficialOpenID:  "string",
				AccountProvider: "wechat-openid",
			},
		},
	})
	//resp := client.SetUserLoginState(ctx, SetUserLoginStateRequest{
	//	UserID:   "xx",
	//	DeviceID: "xxx",
	//	Platform: "pc",
	//	Online:   1,
	//})
	t.Log(resp)
}
