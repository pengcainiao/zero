package sysx

import (
	"os"

	"gitlab.flyele.vip/server-side/go-zero/v2/core/stringx"
)

var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = stringx.RandId()
	}
}

// Hostname returns the name of the host, if no hostname, a random id is returned.
func Hostname() string {
	return hostname
}
