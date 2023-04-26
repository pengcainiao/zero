package events

import "gitlab.flyele.vip/server-side/go-zero/v2/tools/sensors"

type LogOutEvent struct {
	sensors.BaseEvent
	ActivePassive string `json:"active_passive,omitempty"`  // 主动被动操作
}

func (event LogOutEvent) Name() string {
	return "log_out"
}

func (event LogOutEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(event)
}