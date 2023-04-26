package events

import "gitlab.flyele.vip/server-side/go-zero/v2/tools/sensors"

type PushQrCodeEvent struct {
	sensors.BaseEvent
}

func (e PushQrCodeEvent) Name() string {
	return "push_qr_code"
}

func (e PushQrCodeEvent) Properties() map[string]interface{} {
	return sensors.Struct2Map(e)
}
