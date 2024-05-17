package stackTrace

import (
	"context"
	"strings"
)

type (
	Trace struct {
		stack []string
		lock  bool
	}
)

const (
	Key = "request_id"
)

func Init(ctx context.Context) context.Context {
	trace := &Trace{stack: make([]string, 0, 10)}
	return context.WithValue(ctx, Key, trace)
}

func InitWithCap(ctx context.Context, cap int) context.Context {
	trace := &Trace{stack: make([]string, 0, cap)}
	return context.WithValue(ctx, Key, trace)
}

func Lock(ctx context.Context) {
	val, ok := ctx.Value(Key).(*Trace)
	if ok {
		val.lock = true
	}
}

func IsLock(ctx context.Context) bool {
	val, ok := ctx.Value(Key).(*Trace)
	if ok {
		return val.lock
	}
	return false
}

func Add(ctx context.Context, text string) {
	val, ok := ctx.Value(Key).(*Trace)
	if ok {
		val.stack = append(val.stack, text)
	}
}

func Done(ctx context.Context) {
	val, ok := ctx.Value(Key).(*Trace)
	if ok {
		val.stack = val.stack[:len(val.stack)-1]
	}
}

func Get(ctx context.Context) string {
	val, ok := ctx.Value(Key).(*Trace)
	if ok {
		var res strings.Builder
		res.WriteString("[")
		for i, str := range val.stack {
			res.WriteString(str)
			if i != len(val.stack)-1 {
				res.WriteString(" | ")
			}
		}
		res.WriteString("]")
		return res.String()
	}
	return "Not stack trace"
}
