package events

import "gitlab.flyele.vip/server-side/go-zero/v2/tools/sensors"

type UserOrigin struct {
	sensors.BaseEvent
	UserOrigin string `json:"user_origin"`
	IsNewUser  bool   `json:"is_new_user"`
}

func (e UserOrigin) Name() string {
	return "user_origin"
}

func (e UserOrigin) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
