package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"io"
	"reflect"
	"time"
	"unsafe"
)

func InitJaeger() (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		ServiceName: "kademo",
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "192.168.1.30:6831",
			LogSpans:           true,
		},
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory
	closer, err := cfg.InitGlobalTracer("kademo", config.Logger(jLogger), config.Metrics(jMetricsFactory))
	if err != nil {
		panic(err)
	}
	return opentracing.GlobalTracer(), closer
}
func main() {
	tr, closer := InitJaeger()
	defer closer.Close()
	opentracing.InitGlobalTracer(tr)

	//traceID TraceID, spanID, parentID SpanID, sampled bool, baggage map[string]string
	traceId, _ := jaeger.TraceIDFromString("752aa5d597d7c7e1")
	spanId, _ := jaeger.SpanIDFromString("56fad8b70e01af11")
	pid, _ := jaeger.SpanIDFromString("752aa5d597d7c7e0")
	spanCtx := jaeger.NewSpanContext(traceId, spanId, pid, true, nil)
	sp := tr.StartSpan("test2")
	spp := sp.(*jaeger.Span)
	ri := reflect.ValueOf(sp).Elem()

	va := ri.FieldByName("context")
	ml := reflect.NewAt(va.Type(), unsafe.Pointer(va.UnsafeAddr())).Elem()
	ml.Set(reflect.ValueOf(spanCtx))
	fmt.Println(traceId, spp)
	time.Sleep(time.Second)
	sp.SetTag("a", "hahaha")
	sp.Finish()

}
