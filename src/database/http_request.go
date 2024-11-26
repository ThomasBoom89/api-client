package database

import "gorm.io/gorm"

type HttpRequest struct {
	gorm.Model
	Name            string
	CollectionID    uint
	Url             string
	Method          string `gorm:"default:GET"`
	HttpRequestBody HttpRequestBody
}

type HttpRequestBody struct {
	gorm.Model
	HttpRequestID uint
	Type          string `gorm:"default:plaintext"`
	Payload       string
}

type HttpRequestRepository struct {
	database *gorm.DB
}

func NewHttpRequestRepository(database *gorm.DB) *HttpRequestRepository {
	return &HttpRequestRepository{database: database}
}

func (H *HttpRequestRepository) GetAll() ([]HttpRequest, error) {
	var httpRequests []HttpRequest
	err := H.database.Joins("HttpRequestBody").Find(&httpRequests).Error
	if err != nil {
		return []HttpRequest{}, err
	}

	return httpRequests, nil
}

func (H *HttpRequestRepository) GetById(id uint) (*HttpRequest, error) {
	httpRequest := &HttpRequest{}
	err := H.database.Model(&HttpRequest{}).Where("id = ?", id).Preload("HttpRequestBody").First(httpRequest).Error
	if err != nil {
		return nil, err
	}

	return httpRequest, nil
}

func (H *HttpRequestRepository) Create(request *HttpRequest) (*HttpRequest, error) {
	err := H.database.Create(request).Error
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (H *HttpRequestRepository) Update(request *HttpRequest) (*HttpRequest, error) {
	err := H.database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(request).Error
	if err != nil {
		return request, err
	}

	return request, nil
}
func (H *HttpRequestRepository) Delete(request *HttpRequest) error {
	err := H.database.Delete(request).Error
	if err != nil {
		return err
	}

	return nil
}
