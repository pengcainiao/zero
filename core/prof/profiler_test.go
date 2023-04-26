package prof

import (
	"testing"

	"gitlab.flyele.vip/server-side/go-zero/v2/core/utils"
)

func TestProfiler(t *testing.T) {
	EnableProfiling()
	Start()
	Report("foo", ProfilePoint{
		ElapsedTimer: utils.NewElapsedTimer(),
	})
}

func TestNullProfiler(t *testing.T) {
	p := newNullProfiler()
	p.Start()
	p.Report("foo", ProfilePoint{
		ElapsedTimer: utils.NewElapsedTimer(),
	})
}
