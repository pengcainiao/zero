package rescue

import "gitlab.flyele.vip/server-side/go-zero/v2/core/logx"

// Recover is used with defer to do cleanup on panics.
// Use it like:
//  defer Recover(func() {})
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logx.ErrorStack(p)
	}
}
