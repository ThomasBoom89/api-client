package database

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	Name         string
	CollectionID uint
}
