package stackTrace

import (
	"context"
	"strings"
	"sync"
)

type (
	Trace struct {
		stack []string
		lock  bool
		mu    sync.Mutex
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
		val.mu.Lock()
		val.lock = true
		val.mu.Unlock()
	}
}

func IsLock(ctx context.Context) bool {
	val, ok := ctx.Value(Key).(*Trace)
	if ok {
		val.mu.Lock()
		defer val.mu.Unlock()
		return val.lock
	}
	return false
}

func Add(ctx context.Context, text string) {
	val, ok := ctx.Value(Key).(*Trace)
	if ok {
		val.mu.Lock()
		val.stack = append(val.stack, text)
		val.mu.Unlock()
	}
}

func Done(ctx context.Context) {
	val, ok := ctx.Value(Key).(*Trace)
	if ok {
		val.mu.Lock()
		if len(val.stack) > 0 {
			val.stack = val.stack[:len(val.stack)-1]
		}
		val.mu.Unlock()
	}

}

func Get(ctx context.Context) string {
	val, ok := ctx.Value(Key).(*Trace)
	if ok {
		val.mu.Lock()
		defer val.mu.Unlock()
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
	return "no stack trace found in context"
}
