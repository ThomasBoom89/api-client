package database

import "gorm.io/gorm"

type Http struct {
	gorm.Model
}

type Grpc struct {
	gorm.Model
}

type Websocket struct {
	gorm.Model
}

func (Http) TableName() string {
	return "http_requests"
}

func (Grpc) TableName() string {
	return "grpc_requests"
}

func (Websocket) TableName() string {
	return "websocket_requests"
}
