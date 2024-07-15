package database

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name      string
	Requests  []Request
	ProjectID uint
}
