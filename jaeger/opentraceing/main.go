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

/*docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14250:14250 \
-p 14268:14268 \
-p 14269:14269 \
-p 9411:9411 \
jaegertracing/all-in-one:1.32*/
func InitJaeger() (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		ServiceName: "kademo1",
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "192.168.1.30:6831",
			LogSpans:           true,
		},
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory
	closer, err := cfg.InitGlobalTracer("kademo1", config.Logger(jLogger), config.Metrics(jMetricsFactory))
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
	sp := tr.StartSpan("test3")
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
