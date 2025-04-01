package test

import (
	"context"
	"os"
)

type UserDir struct {
	Dir string
}

func (U *UserDir) GetConfigPath() string {
	return U.Dir
}
func (U *UserDir) GetDataPath() string {
	return U.Dir
}

func (U *UserDir) Cleanup() {
	err := os.RemoveAll(U.Dir)
	if err != nil {
		panic(err)
	}
}

type Event struct {
	receive chan []interface{}
}

func NewEvent(receive chan []interface{}) *Event {
	return &Event{receive: receive}
}

func (E *Event) EventsEmit(ctx context.Context, eventName string, optionalData ...interface{}) {
	E.receive <- optionalData
}
