package frontend

import (
	"api-client/src/configuration"
	"api-client/src/database"
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	httpRequestRepository *database.HttpRequestRepository
	configuration         *configuration.ReadWriter
}

type RequestResponseDTO struct {
	Error          string              `json:"error"`
	Url            string              `json:"url"`
	Method         string              `json:"method"`
	ResponseBody   string              `json:"responseBody"`
	SendHeader     map[string][]string `json:"sendHeader"`
	ReceivedHeader map[string][]string `json:"receivedHeader"`
	ElapsedTime    string              `json:"elapsedTime"`
	StatusCode     int                 `json:"statusCode"`
}

func NewRequest(httpRequestRepository *database.HttpRequestRepository, configuration *configuration.ReadWriter) *Request {
	return &Request{
		httpRequestRepository: httpRequestRepository,
		configuration:         configuration,
	}
}

func (R *Request) Submit(requestId uint) (requestResponseDto RequestResponseDTO) {
	httpRequest, err := R.httpRequestRepository.GetById(requestId)
	if err != nil {
		requestResponseDto.Error = err.Error()
		return
	}

	url := R.prepareUrl(httpRequest)

	var request *http.Request
	if httpRequest.HttpRequestBody.Type == "" || httpRequest.HttpRequestBody.Type == "none" {
		request, err = http.NewRequest(httpRequest.Method, url, nil)
	} else {
		request, err = http.NewRequest(httpRequest.Method, url, strings.NewReader(httpRequest.HttpRequestBody.Payload))
	}
	if err != nil {
		requestResponseDto.Error = err.Error()
		return
	}

	switch httpRequest.HttpRequestBody.Type {
	case "json":
		request.Header.Set("Content-Type", "application/json")
	case "plaintext":
		request.Header.Set("Content-Type", "text/plain")
	}

	for _, header := range httpRequest.HttpRequestHeader {
		request.Header.Set(header.Key, header.Value)
	}

	requestResponseDto.Url = request.URL.String()
	requestResponseDto.Method = request.Method
	R.handleHeader(request, &requestResponseDto)

	config, err := R.configuration.Read()
	if err != nil {
		return
	}

	client := http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: config.SkipTLSVerify,
			},
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
	if response.ContentLength > -1 {
		requestResponseDto.ReceivedHeader["Content-Length"] = []string{strconv.Itoa(int(response.ContentLength))}
	}
	requestResponseDto.StatusCode = response.StatusCode

	var buffer []byte
	buffer, err = io.ReadAll(response.Body)
	if err != nil {
		requestResponseDto.Error = err.Error()
		return
	}
	requestResponseDto.ResponseBody = string(buffer)

	return
}

func (R *Request) handleHeader(request *http.Request, requestResponseDto *RequestResponseDTO) {
	request.Header.Set("User-Agent", "")
	buildHeaders := request.Header.Clone()
	buildHeaders.Add("Host", request.Host)
	if buildHeaders.Get("User-Agent") == "" {
		buildHeaders.Del("User-Agent")
	}

	requestResponseDto.SendHeader = buildHeaders
	if request.ContentLength > 0 {
		requestResponseDto.SendHeader["Content-Length"] = []string{strconv.Itoa(int(request.ContentLength))}
	}
}

func (R *Request) prepareUrl(request *database.HttpRequest) string {
	queryParameter := url.Values{}
	for _, parameter := range request.HttpRequestParameter {
		queryParameter.Add(parameter.Key, parameter.Value)
	}
	if queryParameter.Encode() == "" {
		return request.Url
	}

	return request.Url + "?" + queryParameter.Encode()
}
