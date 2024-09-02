package database

import "gorm.io/gorm"

type HttpRequest struct {
	gorm.Model
	Name         string
	CollectionID uint
	Url          string
}

type GrpcRequest struct {
	gorm.Model
	Name         string
	CollectionID uint
}

type WebsocketRequest struct {
	gorm.Model
	Name         string
	CollectionID uint
}
