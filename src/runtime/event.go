package runtime

import (
	"context"
	runtime2 "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Event interface {
	EventsEmit(ctx context.Context, eventName string, optionalData ...interface{})
}

type Wails struct {
}

func (W *Wails) EventsEmit(ctx context.Context, eventName string, optionalData ...interface{}) {
	runtime2.EventsEmit(ctx, eventName, optionalData...)
}
