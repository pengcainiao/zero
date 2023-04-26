package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.flyele.vip/server-side/go-zero/v2/core/prometheus"
)

func TestPromMetricHandler_Disabled(t *testing.T) {
	promMetricHandler := PrometheusHandler()
	handler := promMetricHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestPromMetricHandler_Enabled(t *testing.T) {
	prometheus.StartAgent(prometheus.Config{
		Host: "localhost",
		Path: "/",
	})
	promMetricHandler := PrometheusHandler()
	handler := promMetricHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
