package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"time"
)

type Tracer struct {
	tracerCloser io.Closer
}

func SetGlobalTracer(serviceName string) *Tracer {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	cfg.ServiceName = serviceName
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatal("cannot create tracer", err.Error())
	}
	opentracing.SetGlobalTracer(tracer)

	return &Tracer{tracerCloser: closer}
}

func (t *Tracer) Close() error {
	return t.tracerCloser.Close()
}
