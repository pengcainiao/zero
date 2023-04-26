package handler

import (
	"fmt"
	"net/http"
	"strings"

	"gitlab.flyele.vip/server-side/go-zero/v2/core/breaker"
	"gitlab.flyele.vip/server-side/go-zero/v2/core/logx"
	"gitlab.flyele.vip/server-side/go-zero/v2/core/stat"
	"gitlab.flyele.vip/server-side/go-zero/v2/rest/httpx"
	"gitlab.flyele.vip/server-side/go-zero/v2/rest/internal/response"
)

const breakerSeparator = "://"

// BreakerHandler returns a break circuit middleware.
func BreakerHandler(method, path string, metrics *stat.Metrics) func(http.Handler) http.Handler {
	brk := breaker.NewBreaker(breaker.WithName(strings.Join([]string{method, path}, breakerSeparator)))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			promise, err := brk.Allow()
			if err != nil {
				metrics.AddDrop()
				logx.Errorf("[http] dropped, %s - %s - %s",
					r.RequestURI, httpx.GetRemoteAddr(r), r.UserAgent())
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			cw := &response.WithCodeResponseWriter{Writer: w}
			defer func() {
				if cw.Code < http.StatusInternalServerError {
					promise.Accept()
				} else {
					promise.Reject(fmt.Sprintf("%d %s", cw.Code, http.StatusText(cw.Code)))
				}
			}()
			next.ServeHTTP(cw, r)
		})
	}
}
