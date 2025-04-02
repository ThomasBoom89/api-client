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

type WebsocketRequestDto struct {
	ID           uint      `json:"id"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	CollectionID uint      `json:"collectionId"`
	Url          string    `json:"url"`
}

type HttpRequestDto struct {
	ID           uint                      `json:"id"`
	UpdatedAt    time.Time                 `json:"updatedAt"`
	Name         string                    `json:"name"`
	Type         string                    `json:"type"`
	CollectionID uint                      `json:"collectionId"`
	Url          string                    `json:"url"`
	Method       string                    `json:"method"`
	Body         HttpRequestBodyDto        `json:"body"`
	Parameter    []HttpRequestParameterDto `json:"parameter"`
	Header       []HttpRequestHeaderDto    `json:"header"`
}

type HttpRequestBodyDto struct {
	ID            uint      `json:"id"`
	UpdatedAt     time.Time `json:"updatedAt"`
	HttpRequestID uint      `json:"httpRequestID"`
	Type          string    `json:"type"`
	Payload       string    `json:"payload"`
}

type HttpRequestParameterDto struct {
	ID            uint      `json:"id"`
	UpdatedAt     time.Time `json:"updatedAt"`
	HttpRequestID uint      `json:"httpRequestID"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
}

type HttpRequestHeaderDto struct {
	ID            uint      `json:"id"`
	UpdatedAt     time.Time `json:"updatedAt"`
	HttpRequestID uint      `json:"httpRequestID"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
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

type WebsocketRequests struct {
	websocketRequestRepository *database.WebsocketRequestRepository
}

type Requests struct {
	httpRequests      *HttpRequests
	websocketRequests *WebsocketRequests
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

func NewWebsocketRequests(websocketRequestRepository *database.WebsocketRequestRepository) *WebsocketRequests {
	return &WebsocketRequests{websocketRequestRepository}
}

func NewRequests(httpRequests *HttpRequests, websocketRequests *WebsocketRequests) *Requests {
	return &Requests{httpRequests, websocketRequests}
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

func (R *Requests) GetAll() ([]interface{}, error) {
	httpRequestDtos, err := R.httpRequests.GetAll()
	if err != nil {
		return nil, err
	}
	websocketRequestDtos, err := R.websocketRequests.GetAll()
	if err != nil {
		return nil, err
	}
	requestDtos := make([]interface{}, 0)
	for _, httpRequestDto := range httpRequestDtos {
		requestDtos = append(requestDtos, httpRequestDto)
	}
	for _, websocketRequestDto := range websocketRequestDtos {
		requestDtos = append(requestDtos, websocketRequestDto)
	}

	return requestDtos, err
}

func (H *HttpRequests) GetAll() ([]HttpRequestDto, error) {
	httpRequests, err := H.httpRequestRepository.GetAll()
	if err != nil {
		return nil, err
	}
	httpRequestDtos := make([]HttpRequestDto, len(httpRequests))
	for iter, request := range httpRequests {
		httpRequestDtos[iter] = H.buildDtoFromDatabase(request)
	}

	return httpRequestDtos, nil
}

func (W *WebsocketRequests) GetAll() ([]WebsocketRequestDto, error) {
	websocketRequests, err := W.websocketRequestRepository.GetAll()
	if err != nil {
		return nil, err
	}
	websocketRequestDtos := make([]WebsocketRequestDto, len(websocketRequests))
	for iter, websocketRequest := range websocketRequests {
		websocketRequestDtos[iter] = WebsocketRequestDto{
			ID:           websocketRequest.ID,
			UpdatedAt:    websocketRequest.UpdatedAt,
			Name:         websocketRequest.Name,
			Type:         "websocket",
			CollectionID: websocketRequest.CollectionID,
			Url:          websocketRequest.Url,
		}
	}

	return websocketRequestDtos, nil
}

func (W *WebsocketRequests) Create(websocketRequestDto WebsocketRequestDto) (WebsocketRequestDto, error) {
	websocketRequest := &database.WebsocketRequest{
		Name:         websocketRequestDto.Name,
		CollectionID: websocketRequestDto.CollectionID,
		Url:          websocketRequestDto.Url,
	}
	websocketRequest, err := W.websocketRequestRepository.Create(websocketRequest)
	if err != nil {
		return websocketRequestDto, err
	}
	websocketRequestDto.ID = websocketRequest.ID
	websocketRequestDto.UpdatedAt = websocketRequest.UpdatedAt
	websocketRequestDto.Type = "websocket"

	return websocketRequestDto, nil
}

func (W *WebsocketRequests) Update(websocketRequestDto WebsocketRequestDto) (WebsocketRequestDto, error) {
	websocketRequest, err := W.websocketRequestRepository.GetById(websocketRequestDto.ID)
	if err != nil {
		return websocketRequestDto, err
	}
	websocketRequest.Name = websocketRequestDto.Name
	websocketRequest.Url = websocketRequestDto.Url
	websocketRequest, err = W.websocketRequestRepository.Update(websocketRequest)
	if err != nil {
		return websocketRequestDto, err
	}

	return websocketRequestDto, nil
}

func (W *WebsocketRequests) Delete(websocketRequestDto WebsocketRequestDto) error {
	websocketRequest, err := W.websocketRequestRepository.GetById(websocketRequestDto.ID)
	if err != nil {
		return err
	}

	return W.websocketRequestRepository.Delete(websocketRequest)
}

func (H *HttpRequests) buildDtoFromDatabase(httpRequest database.HttpRequest) HttpRequestDto {
	parameterDtos := make([]HttpRequestParameterDto, len(httpRequest.HttpRequestParameter))
	for iter, parameter := range httpRequest.HttpRequestParameter {
		parameterDtos[iter] = HttpRequestParameterDto{
			ID:            parameter.ID,
			UpdatedAt:     parameter.UpdatedAt,
			HttpRequestID: parameter.HttpRequestID,
			Key:           parameter.Key,
			Value:         parameter.Value,
		}
	}

	headerDtos := make([]HttpRequestHeaderDto, len(httpRequest.HttpRequestHeader))
	for iter, header := range httpRequest.HttpRequestHeader {
		headerDtos[iter] = HttpRequestHeaderDto{
			ID:            header.ID,
			UpdatedAt:     header.UpdatedAt,
			HttpRequestID: header.HttpRequestID,
			Key:           header.Key,
			Value:         header.Value,
		}
	}

	return HttpRequestDto{
		ID:           httpRequest.ID,
		UpdatedAt:    httpRequest.UpdatedAt,
		Name:         httpRequest.Name,
		CollectionID: httpRequest.CollectionID,
		Url:          httpRequest.Url,
		Type:         "http",
		Method:       httpRequest.Method,
		Body: HttpRequestBodyDto{
			ID:            httpRequest.HttpRequestBody.ID,
			UpdatedAt:     httpRequest.HttpRequestBody.UpdatedAt,
			HttpRequestID: httpRequest.HttpRequestBody.HttpRequestID,
			Type:          httpRequest.HttpRequestBody.Type,
			Payload:       httpRequest.HttpRequestBody.Payload,
		},
		Parameter: parameterDtos,
		Header:    headerDtos,
	}

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
	httpRequestDto.Method = httpRequest.Method
	httpRequestDto.Body.HttpRequestID = httpRequest.ID
	httpRequestDto.Body.ID = httpRequest.HttpRequestBody.ID
	httpRequestDto.Body.Type = httpRequest.HttpRequestBody.Type
	httpRequestDto.Body.Payload = httpRequest.HttpRequestBody.Payload
	httpRequestDto.Body.UpdatedAt = httpRequest.HttpRequestBody.UpdatedAt

	return httpRequestDto, nil
}

func (H *HttpRequests) Update(httpRequestDto HttpRequestDto) (HttpRequestDto, error) {

	currentHttpRequest, err := H.httpRequestRepository.GetById(httpRequestDto.ID)
	if err != nil {
		return httpRequestDto, err
	}
	httpRequestParameter, err := H.prepareParameter(httpRequestDto, currentHttpRequest)
	if err != nil {
		return httpRequestDto, err
	}
	httpRequestHeader, err := H.prepareHeader(httpRequestDto, currentHttpRequest)
	if err != nil {
		return httpRequestDto, err
	}

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
		HttpRequestParameter: httpRequestParameter,
		HttpRequestHeader:    httpRequestHeader,
	}
	newHttpRequest, err := H.httpRequestRepository.Update(httpRequest)
	if err != nil {
		return httpRequestDto, err
	}

	return H.buildDtoFromDatabase(*newHttpRequest), nil
}

func (H *HttpRequests) Delete(httpRequestDto HttpRequestDto) error {
	httpRequest := &database.HttpRequest{
		Model: gorm.Model{
			ID: httpRequestDto.ID,
		},
	}

	return H.httpRequestRepository.Delete(httpRequest)
}

func (H *HttpRequests) prepareParameter(httpRequestDto HttpRequestDto, currentHttpRequest *database.HttpRequest) ([]database.HttpRequestParameter, error) {
	var httpRequestParameter []database.HttpRequestParameter
	notDelete := make(map[uint]bool)
	for _, parameter := range httpRequestDto.Parameter {
		if parameter.ID == 0 {
			newParameter, err := H.createParameter(parameter)
			if err != nil {
				return nil, err
			}
			httpRequestParameter = append(httpRequestParameter, newParameter)
			parameter.ID = newParameter.ID
			parameter.UpdatedAt = newParameter.UpdatedAt
			parameter.HttpRequestID = newParameter.HttpRequestID

			continue
		}
		for _, currentParameter := range currentHttpRequest.HttpRequestParameter {
			if parameter.ID != currentParameter.ID {
				continue
			}
			if parameter.Key != currentParameter.Key {
				currentParameter.Key = parameter.Key
			}
			if parameter.Value != currentParameter.Value {
				currentParameter.Value = parameter.Value
			}
			httpRequestParameter = append(httpRequestParameter, currentParameter)
			notDelete[currentParameter.ID] = true
		}
	}
	for _, parameter := range currentHttpRequest.HttpRequestParameter {
		if _, ok := notDelete[parameter.ID]; !ok {
			err := H.httpRequestRepository.DeleteParameter(&parameter)
			if err != nil {
				return nil, err
			}
		}
	}
	return httpRequestParameter, nil
}

func (H *HttpRequests) createParameter(httpRequestParameterDto HttpRequestParameterDto) (database.HttpRequestParameter, error) {
	httpRequestParameter := &database.HttpRequestParameter{
		HttpRequestID: httpRequestParameterDto.HttpRequestID,
		Key:           httpRequestParameterDto.Key,
		Value:         httpRequestParameterDto.Value,
	}

	httpRequestParameter, err := H.httpRequestRepository.CreateParameter(httpRequestParameter)
	if err != nil {
		return database.HttpRequestParameter{}, err
	}

	return *httpRequestParameter, nil
}

func (H *HttpRequests) prepareHeader(httpRequestDto HttpRequestDto, currentHttpRequest *database.HttpRequest) ([]database.HttpRequestHeader, error) {
	var httpRequestHeaders []database.HttpRequestHeader
	notDelete := make(map[uint]bool)
	for _, header := range httpRequestDto.Header {
		if header.ID == 0 {
			newParameter, err := H.createHeader(header)
			if err != nil {
				return nil, err
			}
			httpRequestHeaders = append(httpRequestHeaders, newParameter)
			header.ID = newParameter.ID
			header.UpdatedAt = newParameter.UpdatedAt
			header.HttpRequestID = newParameter.HttpRequestID

			continue
		}
		for _, currentHeader := range currentHttpRequest.HttpRequestHeader {
			if header.ID != currentHeader.ID {
				continue
			}
			if header.Key != currentHeader.Key {
				currentHeader.Key = header.Key
			}
			if header.Value != currentHeader.Value {
				currentHeader.Value = header.Value
			}
			httpRequestHeaders = append(httpRequestHeaders, currentHeader)
			notDelete[currentHeader.ID] = true
		}
	}
	for _, header := range currentHttpRequest.HttpRequestHeader {
		if _, ok := notDelete[header.ID]; !ok {
			err := H.httpRequestRepository.DeleteHeader(&header)
			if err != nil {
				return nil, err
			}
		}
	}
	return httpRequestHeaders, nil
}

func (H *HttpRequests) createHeader(httpRequestHeaderDto HttpRequestHeaderDto) (database.HttpRequestHeader, error) {
	httpRequestHeader := &database.HttpRequestHeader{
		HttpRequestID: httpRequestHeaderDto.HttpRequestID,
		Key:           httpRequestHeaderDto.Key,
		Value:         httpRequestHeaderDto.Value,
	}

	httpRequestHeader, err := H.httpRequestRepository.CreateHeader(httpRequestHeader)
	if err != nil {
		return database.HttpRequestHeader{}, err
	}

	return *httpRequestHeader, nil
}
