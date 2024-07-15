package database

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string
	Collections []Collection
}
