package events

import "gitlab.flyele.vip/server-side/go-zero/v2/tools/sensors"

type DailyActiveEvent struct {
	sensors.BaseEvent
}

func (d DailyActiveEvent) Name() string {
	return "daily_active_event"
}

func (d DailyActiveEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(d)
}
