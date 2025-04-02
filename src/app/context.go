package app

import (
	"context"
	"time"
)

type Context struct {
	Ctx    context.Context
	cancel context.CancelFunc
}

func NewContext() *Context {
	return &Context{}
}

func (C *Context) SetContext(ctx context.Context) {
	C.Ctx, C.cancel = context.WithCancel(ctx)
	time.Sleep(5 * time.Second)
}

func (C *Context) Cancel(ctx context.Context) {
	C.cancel()
}
