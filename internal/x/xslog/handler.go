package xslog

import (
	"context"
	"log/slog"
)

type Handler struct {
	Handler slog.Handler
}

var _ slog.Handler = &Handler{}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		Handler: h.Handler.WithAttrs(attrs),
	}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		Handler: h.Handler.WithGroup(name),
	}
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	ctxArgs, ok := ctx.Value(attrKey).(contextArgs)
	if ok {
		for _, args := range ctxArgs {
			r.Add(args...)
		}
	}
	return h.Handler.Handle(ctx, r)
}
