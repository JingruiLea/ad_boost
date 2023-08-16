package utils

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"runtime/debug"
)

func SafeGo(ctx context.Context, f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				stackTrace := string(debug.Stack())
				logs.CtxErrorf(ctx, "Recovered a panic: %v\nStack trace: %s", r, stackTrace)
			}
		}()
		f()
	}()
}
