package database

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name              string
	ProjectID         uint
	HttpRequests      []HttpRequest
	GrpcRequests      []GrpcRequest
	WebsocketRequests []WebsocketRequest
}
