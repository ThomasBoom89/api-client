package frontend

import (
	"api-client/src/database"
	"gorm.io/gorm"
	"time"
)

type ProjectDto struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

type CollectionDto struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ProjectID uint      `json:"projectId"`
}

type HttpRequestDto struct {
	ID           uint               `json:"id"`
	UpdatedAt    time.Time          `json:"updatedAt"`
	Name         string             `json:"name"`
	Type         string             `json:"type"`
	CollectionID uint               `json:"collectionId"`
	Url          string             `json:"url"`
	Method       string             `json:"method"`
	Body         HttpRequestBodyDto `json:"body"`
}

type HttpRequestBodyDto struct {
	ID            uint      `json:"id"`
	UpdatedAt     time.Time `json:"updatedAt"`
	HttpRequestID uint      `json:"httpRequestID"`
	Type          string    `json:"type"`
	Payload       string    `json:"payload"`
}

type Projects struct {
	projectRepository *database.Repository[database.Project]
}

type Collections struct {
	collectionRepository *database.Repository[database.Collection]
}

type HttpRequests struct {
	httpRequestRepository *database.HttpRequestRepository
}

func NewProjects(projectRepository *database.Repository[database.Project]) *Projects {
	return &Projects{projectRepository}
}

func NewCollections(collectionRepository *database.Repository[database.Collection]) *Collections {
	return &Collections{collectionRepository}
}

func NewHttpRequests(httpRequestRepository *database.HttpRequestRepository) *HttpRequests {
	return &HttpRequests{httpRequestRepository}
}

func (P *Projects) Create(projectDto ProjectDto) (ProjectDto, error) {
	project := &database.Project{
		Name: projectDto.Name,
	}

	project, err := P.projectRepository.Create(project)
	if err != nil {
		return projectDto, err
	}

	projectDto = ProjectDto{
		ID:        project.ID,
		UpdatedAt: project.UpdatedAt,
		Name:      project.Name,
	}

	return projectDto, nil
}

func (P *Projects) Update(projectDto ProjectDto) (ProjectDto, error) {
	project := &database.Project{
		Model: gorm.Model{
			ID: projectDto.ID,
		},
		Name: projectDto.Name,
	}

	project, err := P.projectRepository.Update(project)
	if err != nil {
		return projectDto, err
	}

	projectDto = ProjectDto{
		ID:        project.ID,
		UpdatedAt: project.UpdatedAt,
		Name:      project.Name,
	}

	return projectDto, nil
}

func (P *Projects) Delete(projectDto ProjectDto) error {
	project := &database.Project{
		Model: gorm.Model{
			ID: projectDto.ID,
		},
	}

	return P.projectRepository.Delete(project)
}

func (P *Projects) GetAll() ([]ProjectDto, error) {
	projects, err := P.projectRepository.GetAll()
	if err != nil {
		return nil, err
	}
	projectDtos := make([]ProjectDto, len(projects))
	for iter, project := range projects {
		projectDto := ProjectDto{
			ID:        project.ID,
			UpdatedAt: project.UpdatedAt,
			Name:      project.Name,
		}

		projectDtos[iter] = projectDto
	}

	return projectDtos, nil
}

func (C *Collections) Create(collectionDto CollectionDto) (CollectionDto, error) {
	collection := &database.Collection{
		Name:      collectionDto.Name,
		ProjectID: collectionDto.ProjectID,
	}

	collection, err := C.collectionRepository.Create(collection)
	if err != nil {
		return collectionDto, err
	}
	collectionDto = CollectionDto{
		ID:        collection.ID,
		UpdatedAt: collection.UpdatedAt,
		Name:      collection.Name,
		ProjectID: collection.ProjectID,
	}

	return collectionDto, nil
}

func (C *Collections) Update(collectionDto CollectionDto) (CollectionDto, error) {
	collection := &database.Collection{
		Model: gorm.Model{
			ID: collectionDto.ID,
		},
		Name:      collectionDto.Name,
		ProjectID: collectionDto.ProjectID,
	}

	collection, err := C.collectionRepository.Update(collection)
	if err != nil {
		return collectionDto, err
	}
	collectionDto = CollectionDto{
		ID:        collection.ID,
		UpdatedAt: collection.UpdatedAt,
		Name:      collection.Name,
		ProjectID: collection.ProjectID,
	}

	return collectionDto, nil
}

func (C *Collections) Delete(collectionDto CollectionDto) error {
	collection := &database.Collection{
		Model: gorm.Model{
			ID: collectionDto.ID,
		},
	}

	return C.collectionRepository.Delete(collection)
}

func (C *Collections) GetAll() ([]CollectionDto, error) {
	collections, err := C.collectionRepository.GetAll()
	if err != nil {
		return nil, err
	}
	collectionDtos := make([]CollectionDto, len(collections))
	for iter, collection := range collections {
		collectionDtos[iter] = CollectionDto{
			ID:        collection.ID,
			UpdatedAt: collection.UpdatedAt,
			Name:      collection.Name,
			ProjectID: collection.ProjectID,
		}
	}

	return collectionDtos, nil
}

func (H *HttpRequests) GetAll() ([]HttpRequestDto, error) {
	httpRequests, err := H.httpRequestRepository.GetAll()
	if err != nil {
		return nil, err
	}
	httpRequestDtos := make([]HttpRequestDto, len(httpRequests))
	for iter, request := range httpRequests {
		httpRequestDtos[iter] = HttpRequestDto{
			ID:           request.ID,
			UpdatedAt:    request.UpdatedAt,
			Name:         request.Name,
			CollectionID: request.CollectionID,
			Url:          request.Url,
			Type:         "http",
			Method:       request.Method,
			Body: HttpRequestBodyDto{
				ID:            request.HttpRequestBody.ID,
				UpdatedAt:     request.HttpRequestBody.UpdatedAt,
				HttpRequestID: request.HttpRequestBody.HttpRequestID,
				Type:          request.HttpRequestBody.Type,
				Payload:       request.HttpRequestBody.Payload,
			},
		}
	}

	return httpRequestDtos, nil
}

func (H *HttpRequests) Create(httpRequestDto HttpRequestDto) (HttpRequestDto, error) {
	httpRequest := &database.HttpRequest{
		Name:         httpRequestDto.Name,
		CollectionID: httpRequestDto.CollectionID,
		Url:          httpRequestDto.Url,
		HttpRequestBody: database.HttpRequestBody{
			Type:    "none",
			Payload: "",
		},
	}
	httpRequest, err := H.httpRequestRepository.Create(httpRequest)
	if err != nil {
		return httpRequestDto, err
	}
	httpRequestDto.ID = httpRequest.ID
	httpRequestDto.UpdatedAt = httpRequest.UpdatedAt
	httpRequestDto.Type = "http"
	httpRequestDto.Method = "GET"
	httpRequestDto.Body.HttpRequestID = httpRequest.ID
	httpRequestDto.Body.ID = httpRequest.HttpRequestBody.ID
	httpRequestDto.Body.Type = httpRequest.HttpRequestBody.Type
	httpRequestDto.Body.Payload = httpRequest.HttpRequestBody.Payload
	httpRequestDto.Body.UpdatedAt = httpRequest.HttpRequestBody.UpdatedAt

	return httpRequestDto, nil
}

func (H *HttpRequests) Update(httpRequestDto HttpRequestDto) (HttpRequestDto, error) {

	httpRequest := &database.HttpRequest{
		Model: gorm.Model{
			ID: httpRequestDto.ID,
		},
		Name:         httpRequestDto.Name,
		CollectionID: httpRequestDto.CollectionID,
		Url:          httpRequestDto.Url,
		Method:       httpRequestDto.Method,
		HttpRequestBody: database.HttpRequestBody{
			Model: gorm.Model{
				ID: httpRequestDto.Body.ID,
			},
			HttpRequestID: httpRequestDto.ID,
			Type:          httpRequestDto.Body.Type,
			Payload:       httpRequestDto.Body.Payload,
		},
	}
	httpRequest, err := H.httpRequestRepository.Update(httpRequest)
	if err != nil {
		return httpRequestDto, err
	}
	httpRequestDto.UpdatedAt = httpRequest.UpdatedAt

	return httpRequestDto, nil
}

func (H *HttpRequests) Delete(httpRequestDto HttpRequestDto) error {
	httpRequest := &database.HttpRequest{
		Model: gorm.Model{
			ID: httpRequestDto.ID,
		},
	}

	return H.httpRequestRepository.Delete(httpRequest)
}
