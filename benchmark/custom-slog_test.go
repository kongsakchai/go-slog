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

func CustomLogger(w io.Writer) *slog.Logger {
	logger := slog.New(goslog.NewTextHandler(w, &goslog.HandlerOptions{
		Level:      slog.LevelInfo,
		TimeFormat: "[2006-01-02 15:04:05]",
	}))

	return logger
}

func CustomLoggerWithAttrs(w io.Writer, attrs []slog.Attr) *slog.Logger {
	logger := slog.New(goslog.NewTextHandler(w, &goslog.HandlerOptions{
		Level:      slog.LevelInfo,
		TimeFormat: "[2006-01-02 15:04:05]",
	}).WithAttrs(attrs))

	return logger
}

func (b *customSlogBench) new(w io.Writer) logBenchmark {
	return &customSlogBench{
		l: CustomLogger(w),
	}
}

func (b *customSlogBench) newWithCtx(w io.Writer) logBenchmark {
	return &customSlogBench{
		l: CustomLoggerWithAttrs(w, slogAttrs()),
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
