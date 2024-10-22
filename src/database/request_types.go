package database

import "gorm.io/gorm"

type HttpRequest struct {
	gorm.Model
	Name         string
	CollectionID uint
	Url          string
	Method       string `gorm:"default:GET"`
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
