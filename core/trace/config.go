package trace

import (
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// HttpTraceName represents the tracing name.
const (
	HttpTraceName    = "http-server"
	RedisTraceName   = "redis-client"
	MysqlTraceName   = "mysql-client"
	ElasticTraceName = "es-client"
	RpcTraceName     = "grpc"
	NsqTraceName     = "nsq"
	ZincTraceName    = "zinc"
)

// A Config is a opentelemetry config.
type Config struct {
	Name     string  `json:",optional"`
	Endpoint string  `json:",optional"`
	Sampler  float64 `json:",default=1.0"`
	Batcher  string  `json:",default=jaeger,options=jaeger|zipkin"`
}

func PatchSpanData(sp oteltrace.Span) {
	if sp == nil {
		return
	}
	sp.SetAttributes(
		attribute.String("sls.otel.project", "flyele-tracing"),
		attribute.String("sls.otel.instanceid", "app-all-routers"),
		attribute.String("sls.otel.akid", "LTAI5tPUswtBgkYTVDNehTxN"),
		attribute.String("sls.otel.aksecret", "IdHucDClmhHhsKO5LgV4fdD6E6KEcT"),
	)
}
