package benchmark

import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

type slogBench struct {
	l *slog.Logger
}

func DefaultLogger(w io.Writer) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	return logger
}

func DefaultLoggerWithAttrs(w io.Writer, attrs []slog.Attr) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}).WithAttrs(attrs))

	return logger
}

func (b *slogBench) new(w io.Writer) logBenchmark {
	return &slogBench{
		l: DefaultLogger(w),
	}
}

func (b *slogBench) newWithCtx(w io.Writer) logBenchmark {
	return &slogBench{
		l: DefaultLoggerWithAttrs(w, slogAttrs()),
	}
}

func (b *slogBench) name() string {
	return "Slog"
}

func (b *slogBench) logEvent(msg string) {
	b.l.Info(msg)
}

func (b *slogBench) logEventFmt(msg string, args ...any) {
	b.l.Info(fmt.Sprintf(msg, args...))
}

func (b *slogBench) logEventCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		msg,
		slogAttrs()...,
	)
}

func (b *slogBench) logEventCtxWeak(msg string) {
	b.l.Info(msg, alternatingKeyValuePairs()...)
}

func (b *slogBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *slogBench) logDisabledFmt(msg string, args ...any) {
	b.l.Debug(fmt.Sprintf(msg, args...))
}

func (b *slogBench) logDisabledCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelDebug,
		msg,
		slogAttrs()...,
	)
}

func (b *slogBench) logDisabledCtxWeak(msg string) {
	b.l.Debug(msg, alternatingKeyValuePairs()...)
}
