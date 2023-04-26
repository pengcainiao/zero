package redistest

import (
	"time"

	"github.com/alicebob/miniredis/v2"
	"gitlab.flyele.vip/server-side/go-zero/v2/core/lang"
	"gitlab.flyele.vip/server-side/go-zero/v2/core/stores/redis"
)

// CreateRedis returns a in process redis.Redis.
func CreateRedis() (r *redis.Redis, clean func(), err error) {
	mr, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}

	return redis.New(), func() {
		ch := make(chan lang.PlaceholderType)

		go func() {
			mr.Close()
			close(ch)
		}()

		select {
		case <-ch:
		case <-time.After(time.Second):
		}
	}, nil
}
