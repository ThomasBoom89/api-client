package database

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	Name         string
	CollectionID uint
	Type         string `gorm:"index:idx_type,unique"`
	TypeID       uint   `gorm:"index:idx_type,unique"`
}
