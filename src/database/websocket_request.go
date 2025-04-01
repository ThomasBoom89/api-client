package database

import "gorm.io/gorm"

type WebsocketRequest struct {
	gorm.Model
	Name         string
	CollectionID uint
	Url          string
}

type WebsocketRequestRepository struct {
	database *gorm.DB
}

func NewWebsocketRequestRepository(database *gorm.DB) *WebsocketRequestRepository {
	return &WebsocketRequestRepository{database: database}
}

func (W *WebsocketRequestRepository) GetAll() ([]WebsocketRequest, error) {
	var websocketRequest []WebsocketRequest
	err := W.database.Debug().Find(&websocketRequest).Error
	if err != nil {
		return nil, err
	}

	return websocketRequest, nil
}

func (W *WebsocketRequestRepository) GetById(id uint) (*WebsocketRequest, error) {
	request := &WebsocketRequest{}
	err := W.database.Debug().Find(request, id).Error
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (W *WebsocketRequestRepository) Create(request *WebsocketRequest) (*WebsocketRequest, error) {
	err := W.database.Debug().Create(request).Error
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (W *WebsocketRequestRepository) Update(request *WebsocketRequest) (*WebsocketRequest, error) {
	err := W.database.Debug().Updates(request).Error
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (W *WebsocketRequestRepository) Delete(request *WebsocketRequest) error {
	return W.database.Debug().Delete(request).Error
}
