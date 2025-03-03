package safego

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"runtime/debug"
)

// Go safe goroutine.
func Go(ctx context.Context, fn func()) {
	go func() {
		defer Recovery(ctx)
		fn()
	}()
}

// Recovery .
func Recovery(ctx context.Context) {
	e := recover()
	if e == nil {
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	logc.Errorw(ctx, "sage goroutine catch panic err", logc.Field("err", e), logc.Field("stack", debug.Stack()))
}
