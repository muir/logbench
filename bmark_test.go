// go test -v -cpu=4 -run=none -bench=. -benchtime=10s -benchmem bmark_test.go

package xop_test

import (
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
			xopjson.WithEpochTime(time.Nanosecond),
			xopjson.WithDurationFormat(xopjson.AsNanos),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("disable")
	for i := 0; i < b.N; i++ {
		logger.Debug().String("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
	logger.Done()
}

func BenchmarkNormalXop(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Nanosecond),
			xopjson.WithDurationFormat(xopjson.AsNanos),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("disable")
	for i := 0; i < b.N; i++ {
		logger.Info().String("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
	logger.Done()
}

func BenchmarkInterfaceXop(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Nanosecond),
			xopjson.WithDurationFormat(xopjson.AsNanos),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("disable")
	for i := 0; i < b.N; i++ {
		logger.Info().Any("object", &obj).Msg(msg)
	}
	logger.Done()
}

func BenchmarkPrintfXop(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Nanosecond),
			xopjson.WithDurationFormat(xopjson.AsNanos),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("disable")
	for i := 0; i < b.N; i++ {
		logger.Info().Msgf("rate=%s low=%d high=%f msg=%s", "15", 16, 123.2, msg)
	}
	logger.Done()
}

func BenchmarkCallerXop(b *testing.B) {
	logger := xop.NewSeed(xop.WithBase(
		xopjson.New(
			xopbytes.WriteToIOWriter(ioutil.Discard),
			xopjson.WithEpochTime(time.Nanosecond),
			xopjson.WithDurationFormat(xopjson.AsNanos),
			xopjson.WithSpanTags(xopjson.SpanIDTagOption),
			xopjson.WithAttributesObject(false)))).
		Request("disable").
		Sub().
		StackFrames(xopconst.InfoLevel, 1).
		Log()
	for i := 0; i < b.N; i++ {
		logger.Info().String("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
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

func BenchmarkInterfaceZap(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	)).Sugar()
	for i := 0; i < b.N; i++ {
		logger.Infow(msg, "object", &obj)
	}
}

func BenchmarkPrintfZap(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	)).Sugar()
	for i := 0; i < b.N; i++ {
		logger.Infof("rate=%s low=%d high=%f msg=%s", "15", 16, 123.2, msg)
	}
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

func BenchmarkDisableZeroLog(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Debug().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkNormalZeroLog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkInterfaceZeroLog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().Interface("object", &obj).Msg(msg)
	}
}

func BenchmarkPrintfZeroLog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().Msgf("rate=%s low=%d high=%f msg=%s", "15", 16, 123.2, msg)
	}
}

func BenchmarkCallerZeroLog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).With().Caller().Timestamp().Logger()
	for i := 0; i < b.N; i++ {
		logger.Info().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkDisableOneLog(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.DebugWithFields(msg, func(e onelog.Entry) {
			e.String("rate", "15")
			e.Int("low", 16)
			e.Float("high", 123.2)
		})
	}
}

func BenchmarkDisableOneLogChain(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.DebugWith(msg).String("rate", "15").Int("low", 16).Float("high", 123.2).Write()
	}
}

func BenchmarkNormalOneLog(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.InfoWithFields(msg, func(e onelog.Entry) {
			e.String("rate", "15")
			e.Int("low", 16)
			e.Float("high", 123.2)
		})
	}
}

func BenchmarkNormalOneLogChain(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.InfoWith(msg).String("rate", "15").Int("low", 16).Float("high", 123.2).Write()
	}
}

func BenchmarkInterfaceOneLogChain(b *testing.B) {
	logger := onelog.New(ioutil.Discard, onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL)
	for i := 0; i < b.N; i++ {
		logger.InfoWith(msg).Any("object", &obj).Write()
	}
}

func BenchmarkDisablePhusLog(b *testing.B) {
	logger := log.Logger{Level: log.InfoLevel, Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Debug().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkNormalPhusLog(b *testing.B) {
	logger := log.Logger{Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Info().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}

func BenchmarkInterfacePhusLog(b *testing.B) {
	logger := log.Logger{Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Info().Interface("object", &obj).Msg(msg)
	}
}

func BenchmarkPrintfPhusLog(b *testing.B) {
	logger := log.Logger{Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Info().Msgf("rate=%s low=%d high=%f msg=%s", "15", 16, 123.2, msg)
	}
}

func BenchmarkCallerPhusLog(b *testing.B) {
	logger := log.Logger{Caller: 1, Writer: log.IOWriter{ioutil.Discard}}
	for i := 0; i < b.N; i++ {
		logger.Info().Str("rate", "15").Int("low", 16).Float32("high", 123.2).Msg(msg)
	}
}
