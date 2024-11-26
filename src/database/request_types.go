package database

import "gorm.io/gorm"

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
