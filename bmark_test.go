// go test -v -cpu=4 -run=none -bench=. -benchtime=10s -benchmem bmark_test.go

package xop_test

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/francoispqt/onelog"
	"github.com/muir/xop-go"
	"github.com/muir/xop-go/xopbytes"
	"github.com/muir/xop-go/xopconst"
	"github.com/muir/xop-go/xopjson"
	"github.com/phuslu/log"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var msg = "The quick brown fox jumps over the lazy dog"
var obj = struct {
	Rate string
	Low  int
	High float32
}{"15", 16, 123.2}

func BenchmarkDisableXop(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Microsecond),
			xopjson.WithDuration("dur", xopjson.AsString),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("disable")
	for i := 0; i < b.N; i++ {
		logger.Debug().String("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
	logger.Done()
}

func BenchmarkDisableZap(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	))
	for i := 0; i < b.N; i++ {
		logger.Debug(msg, zap.String("rate", "15"), zap.Int("low", 16), zap.Float32("high", 123.2))
	}
}

func BenchmarkDisableZapSugar(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	)).Sugar()
	for i := 0; i < b.N; i++ {
		logger.Debugw(msg, "rate", "15", "low", 16, "high", 123.2)
	}
}

func BenchmarkDisableZeroLog(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Debug().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkDisableOneLogNoTime(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.DebugWithFields(msg, func(e onelog.Entry) {
			e.String("rate", "15")
			e.Int("low", 16)
			e.Float("high", 123.2)
		})
	}
}

func BenchmarkDisableOneLogNTChain(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.DebugWith(msg).String("rate", "15").Int("low", 16).Float("high", 123.2).Write()
	}
}

func BenchmarkDisablePhusLog(b *testing.B) {
	logger := log.Logger{Level: log.InfoLevel, Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Debug().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkNormalXopMilli(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Microsecond),
			xopjson.WithDuration("dur", xopjson.AsString),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("normal")
	for i := 0; i < b.N; i++ {
		logger.Info().String("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
	logger.Done()
}

func BenchmarkNormalXop3339(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithStrftime("%Y-%m-%dT%k:%M:%S %z"),
			xopjson.WithDuration("dur", xopjson.AsString),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("normal-3339")
	for i := 0; i < b.N; i++ {
		logger.Info().String("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
	logger.Done()
}

func BenchmarkNormalZap(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	))
	for i := 0; i < b.N; i++ {
		logger.Info(msg, zap.String("rate", "15"), zap.Int("low", 16), zap.Float32("high", 123.2))
	}
}

func BenchmarkNormalZapSugar(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	)).Sugar()
	for i := 0; i < b.N; i++ {
		logger.Infow(msg, "rate", "15", "low", 16, "high", 123.2)
	}
}

func BenchmarkNormalZeroLog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkNormalOneLogNoTime(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.InfoWithFields(msg, func(e onelog.Entry) {
			e.String("rate", "15")
			e.Int("low", 16)
			e.Float("high", 123.2)
		})
	}
}

func BenchmarkNormalOneLogNTChain(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.InfoWith(msg).String("rate", "15").Int("low", 16).Float("high", 123.2).Write()
	}
}

func BenchmarkNormalOneLogWithTime(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	logger.Hook(func(e onelog.Entry) {
		e.String("time", time.Now().Format(time.RFC3339))
	})
	for i := 0; i < b.N; i++ {
		logger.InfoWithFields(msg, func(e onelog.Entry) {
			e.String("rate", "15")
			e.Int("low", 16)
			e.Float("high", 123.2)
		})
	}
}

func BenchmarkNormalPhusLog(b *testing.B) {
	logger := log.Logger{Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Info().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkInterfaceXop(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Microsecond),
			xopjson.WithDuration("dur", xopjson.AsString),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("interface")
	for i := 0; i < b.N; i++ {
		logger.Info().Any("object", &obj).Msg(msg)
	}
	logger.Done()
}

func BenchmarkInterfaceZap(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	))
	for i := 0; i < b.N; i++ {
		logger.Info(msg, zap.Any("object", &obj))
	}
}

func BenchmarkInterfaceZapSugar(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	)).Sugar()
	for i := 0; i < b.N; i++ {
		logger.Infow(msg, "object", &obj)
	}
}

func BenchmarkInterfaceZeroLog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().Interface("object", &obj).Msg(msg)
	}
}

func BenchmarkInterfaceOneLogNTChain(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.InfoWith(msg).Any("object", &obj).Write()
	}
}

func BenchmarkInterfacePhusLog(b *testing.B) {
	logger := log.Logger{Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Info().Interface("object", &obj).Msg(msg)
	}
}

func BenchmarkPrintfXop(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Microsecond),
			xopjson.WithDuration("dur", xopjson.AsString),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("printf")
	for i := 0; i < b.N; i++ {
		logger.Info().Msgf("rate=%s low=%d high=%f msg=%s", "15", 16, 123.2, msg)
	}
	logger.Done()
}

func BenchmarkPrintfZapSugar(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	)).Sugar()
	for i := 0; i < b.N; i++ {
		logger.Infof("rate=%s low=%d high=%f msg=%s", "15", 16, 123.2, msg)
	}
}

func BenchmarkPrintfZeroLog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().Msgf("rate=%s low=%d high=%f msg=%s", "15", 16, 123.2, msg)
	}
}

func BenchmarkPrintfPhusLog(b *testing.B) {
	logger := log.Logger{Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Info().Msgf("rate=%s low=%d high=%f msg=%s", "15", 16, 123.2, msg)
	}
}

func BenchmarkCallerXop(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Microsecond),
			xopjson.WithDuration("dur", xopjson.AsString),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("caller").
		Sub().
		StackFrames(xopconst.InfoLevel, 1).
		Log()
	for i := 0; i < b.N; i++ {
		logger.Info().String("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
	logger.Done()
}

func BenchmarkCallerZap(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel),
		zap.AddCaller(),
	)
	for i := 0; i < b.N; i++ {
		logger.Info(msg, zap.String("rate", "15"), zap.Int("low", 16), zap.Float32("high", 123.2))
	}
}

func BenchmarkCallerZeroLog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).With().Caller().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkCallerPhusLog(b *testing.B) {
	logger := log.Logger{Caller: 1, Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Info().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkEmptyXop(b *testing.B) {
	seed := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Microsecond),
			xopjson.WithDuration("dur", xopjson.AsString),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false))))
	for i := 0; i < b.N; i++ {
		seed.Request("empty").Done()
	}
}

func BenchmarkEmptyOTEL(b *testing.B) {
	exp, err := stdouttrace.New(
		stdouttrace.WithWriter(ioutil.Discard),
	)
	if err != nil {
		b.FailNow()
	}
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(r),
	)
	defer func() {
		_ = tp.Shutdown(context.Background())
	}()
	otel.SetTracerProvider(tp)

	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		_, span := otel.Tracer("name").Start(ctx, "Run")
		span.End()
	}
}

func BenchmarkTenspanXop(b *testing.B) {
	seed := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Microsecond),
			xopjson.WithDuration("dur", xopjson.AsString),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false))))
	for i := 0; i < b.N; i++ {
		request := seed.Request("empty")
		for j := 0; j < 10; j++ {
			request.Sub().Step("subspan").Wait().Done()
		}
		request.Done()
	}
}

func BenchmarkTenspanOTEL(b *testing.B) {
	exp, err := stdouttrace.New(
		stdouttrace.WithWriter(ioutil.Discard),
	)
	if err != nil {
		b.FailNow()
	}
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(r),
	)
	defer func() {
		_ = tp.Shutdown(context.Background())
	}()
	otel.SetTracerProvider(tp)

	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		ctx, span := otel.Tracer("name").Start(ctx, "Run")
		for j := 0; j < 10; j++ {
			_, span := otel.Tracer("inner").Start(ctx, "Run")
			span.End()
		}
		span.End()
	}
}
