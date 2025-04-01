package database

import "gorm.io/gorm"

type RequestType interface {
}

type GrpcRequest struct {
	gorm.Model
	Name         string
	CollectionID uint
}
