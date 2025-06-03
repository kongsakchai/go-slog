package benchmark

import (
	"context"
	"fmt"
	"goslog"
	"io"
	"log/slog"
)

type customSlogBench struct {
	l *slog.Logger
}

func (b *customSlogBench) new(w io.Writer) logBenchmark {
	return &customSlogBench{
		l: goslog.CustomLogger(w),
	}
}

func (b *customSlogBench) newWithCtx(w io.Writer) logBenchmark {
	return &customSlogBench{
		l: goslog.CustomLoggerWithAttrs(w, slogAttrs()),
	}
}

func (b *customSlogBench) name() string {
	return "Custom Slog"
}

func (b *customSlogBench) logEvent(msg string) {
	b.l.Info(msg)
}

func (b *customSlogBench) logEventFmt(msg string, args ...any) {
	b.l.Info(fmt.Sprintf(msg, args...))
}

func (b *customSlogBench) logEventCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		msg,
		slogAttrs()...,
	)
}

func (b *customSlogBench) logEventCtxWeak(msg string) {
	b.l.Info(msg, alternatingKeyValuePairs()...)
}

func (b *customSlogBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *customSlogBench) logDisabledFmt(msg string, args ...any) {
	b.l.Debug(fmt.Sprintf(msg, args...))
}

func (b *customSlogBench) logDisabledCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelDebug,
		msg,
		slogAttrs()...,
	)
}

func (b *customSlogBench) logDisabledCtxWeak(msg string) {
	b.l.Debug(msg, alternatingKeyValuePairs()...)
}
