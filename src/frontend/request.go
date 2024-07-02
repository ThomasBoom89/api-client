package frontend

import (
	"io"
	"net/http"
)

type Request struct {
}

func NewRequest() *Request {
	return &Request{}
}

func (R *Request) Submit(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	var buffer []byte
	buffer, err = io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
