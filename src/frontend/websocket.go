package frontend

import (
	"api-client/src/app"
	runtime2 "api-client/src/runtime"
	"golang.org/x/net/websocket"
	"log"
	"runtime"
)

type Websocket struct {
	connections map[uint]*websocket.Conn
	ctx         *app.Context
	event       runtime2.Event
}

type WebsocketStateDto struct {
	ID              uint   `json:"id"`
	Error           string `json:"error"`
	Connected       bool   `json:"connected"`
	IncomingMessage string `json:"incomingMessage"`
}

func NewWebsocket(ctx *app.Context, event runtime2.Event) *Websocket {
	return &Websocket{ctx: ctx, connections: make(map[uint]*websocket.Conn), event: event}
}

func (W *Websocket) Connect(websocketRequestDto WebsocketRequestDto) WebsocketStateDto {
	websocketStateDto := WebsocketStateDto{
		ID:              websocketRequestDto.ID,
		Error:           "",
		Connected:       false,
		IncomingMessage: "",
	}
	connection, err := websocket.Dial(websocketRequestDto.Url, "", "http://localhost")
	if err != nil {
		websocketStateDto.Error = err.Error()

		return websocketStateDto
	}
	websocketStateDto.Connected = true
	// how to shutdown go routine when disconnecting
	go W.receive(websocketRequestDto, connection)
	W.connections[websocketRequestDto.ID] = connection

	return websocketStateDto
}

func (W *Websocket) Disconnect(websocketRequestDto WebsocketRequestDto) WebsocketStateDto {
	websocketStateDto := WebsocketStateDto{
		ID:              websocketRequestDto.ID,
		Error:           "",
		Connected:       false,
		IncomingMessage: "",
	}
	connection, ok := W.connections[websocketRequestDto.ID]
	if ok == false {
		return websocketStateDto
	}
	connection.Close()

	return websocketStateDto
}

func (W *Websocket) Send(websocketRequestDto WebsocketRequestDto, message string) WebsocketStateDto {
	websocketStateDto := WebsocketStateDto{
		ID:              websocketRequestDto.ID,
		Error:           "",
		Connected:       true,
		IncomingMessage: "",
	}

	connection, ok := W.connections[websocketRequestDto.ID]
	if ok == false {
		websocketStateDto.Connected = false

		return websocketStateDto
	}

	_, err := connection.Write([]byte(message))
	if err != nil {
		W.Disconnect(websocketRequestDto)
		websocketStateDto.Error = err.Error()
		websocketStateDto.Connected = false

		return websocketStateDto
	}

	return websocketStateDto
}

func (W *Websocket) receive(websocketRequestDto WebsocketRequestDto, connection *websocket.Conn) {
	defer func() {
		W.Disconnect(websocketRequestDto)
	}()
	for {
		runtime.Gosched()
		websocketStateDto := WebsocketStateDto{
			ID:              websocketRequestDto.ID,
			Error:           "",
			Connected:       true,
			IncomingMessage: "",
		}
		var msg = make([]byte, 512)
		length, err := connection.Read(msg)
		if err != nil {
			websocketStateDto.Error = err.Error()
			websocketStateDto.Connected = false
			W.event.EventsEmit(W.ctx.Ctx, "websocket", websocketStateDto)
			log.Println(err)
			return
		}
		websocketStateDto.IncomingMessage = string(msg[:length])
		W.event.EventsEmit(W.ctx.Ctx, "websocket", websocketStateDto)

	}
}
