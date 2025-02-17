package xtime

import (
	"context"
	"time"
)

type key int

const (
	nowFuncKey key = iota + 1
)

func Now(ctx context.Context) time.Time {
	f, ok := ctx.Value(nowFuncKey).(func() time.Time)
	if ok {
		return f()
	}
	return time.Now()
}

func SetNow(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, nowFuncKey, func() time.Time { return t })
}

func SetNowFunc(ctx context.Context, f func() time.Time) context.Context {
	return context.WithValue(ctx, nowFuncKey, f)
}
