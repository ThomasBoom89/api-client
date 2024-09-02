package database

import (
	"gorm.io/gorm"
)

type Repository[T Project | Collection | HttpRequest | GrpcRequest | WebsocketRequest] struct {
	database *gorm.DB
}

func NewRepository[T Project | Collection | HttpRequest | GrpcRequest | WebsocketRequest](database *gorm.DB) *Repository[T] {
	return &Repository[T]{database}
}

func (R *Repository[T]) Create(value *T) (*T, error) {
	err := R.database.Create(value)
	if err.Error != nil {
		return nil, err.Error
	}

	return value, nil
}

func (R *Repository[T]) Update(value *T) (*T, error) {
	err := R.database.Save(value)
	if err.Error != nil {
		return value, err.Error
	}

	return value, nil
}

func (R *Repository[T]) Delete(value *T) error {
	err := R.database.Delete(value)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (R *Repository[T]) GetAll() ([]T, error) {
	var values []T
	err := R.database.Find(&values)
	if err.Error != nil {
		return nil, err.Error
	}

	return values, nil
}