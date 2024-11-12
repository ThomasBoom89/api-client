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

func (R *Request) Submit(requestId uint) (requestResponseDto RequestResponseDTO) {
	httpRequest, err := R.requestRepository.GetById(requestId)
	if err != nil {
		requestResponseDto.Error = err.Error()
		return
	}
	request, err := http.NewRequest(httpRequest.Method, httpRequest.Url, nil)
	if err != nil {
		requestResponseDto.Error = err.Error()
		return
	}
	R.handleHeader(request, requestResponseDto)

	client := http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}

	startTime := time.Now()
	response, err := client.Do(request)
	endTime := time.Now()
	requestResponseDto.ElapsedTime = endTime.Sub(startTime).String()
	if err != nil {
		requestResponseDto.Error = err.Error()
		return
	}
	requestResponseDto.ReceivedHeader = response.Header

	var buffer []byte
	buffer, err = io.ReadAll(response.Body)
	if err != nil {
		requestResponseDto.Error = err.Error()
		return
	}
	requestResponseDto.ResponseBody = string(buffer)

	return
}

func (R *Request) handleHeader(request *http.Request, requestResponseDto RequestResponseDTO) {
	request.Header.Set("User-Agent", "")
	requestResponseDto.Url = request.URL.String()
	requestResponseDto.Method = request.Method

	buildHeaders := request.Header.Clone()
	buildHeaders.Add("Host", request.Host)
	if buildHeaders.Get("User-Agent") == "" {
		buildHeaders.Del("User-Agent")
	}

	requestResponseDto.SendHeader = buildHeaders
}
