package frontend

import (
	"api-client/src/database"
	"io"
	"net/http"
	"time"
)

type Request struct {
	requestRepository *database.Repository[database.HttpRequest]
}

type RequestResponseDTO struct {
	Error          string              `json:"error"`
	Url            string              `json:"url"`
	Method         string              `json:"method"`
	ResponseBody   string              `json:"responseBody"`
	SendHeader     map[string][]string `json:"sendHeader"`
	ReceivedHeader map[string][]string `json:"receivedHeader"`
	ElapsedTime    string              `json:"elapsedTime"`
}

func NewRequest(requestRepository *database.Repository[database.HttpRequest]) *Request {
	return &Request{requestRepository: requestRepository}
}

func (R *Request) Submit(requestId uint) RequestResponseDTO {
	httpRequest, err := R.requestRepository.GetById(requestId)
	if err != nil {
		return RequestResponseDTO{Error: err.Error()}
	}
	request, err := http.NewRequest(httpRequest.Method, httpRequest.Url, nil)
	if err != nil {
		return RequestResponseDTO{Error: err.Error()}
	}

	client := new(http.Client)
	startTime := time.Now()
	response, err := client.Do(request)
	endTime := time.Now()
	if err != nil {
		return RequestResponseDTO{Error: err.Error()}
	}
	var buffer []byte
	buffer, err = io.ReadAll(response.Body)
	if err != nil {
		return RequestResponseDTO{Error: err.Error()}
	}
	requestResponseDTO := RequestResponseDTO{
		Error:          "",
		Url:            request.URL.String(),
		Method:         request.Method,
		ResponseBody:   string(buffer),
		SendHeader:     request.Header,
		ReceivedHeader: response.Header,
		ElapsedTime:    endTime.Sub(startTime).String(),
	}

	return requestResponseDTO
}
