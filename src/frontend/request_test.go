package frontend

import (
	"api-client/src/configuration"
	"api-client/src/database"
	"api-client/src/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequest(t *testing.T) {
	userDir := test.UserDir{Dir: "./tmp-request_test/"}
	defer userDir.Cleanup()
	readWriter := configuration.NewReadWriter(&userDir)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		header := request.Header.Get("tom")
		if header != "riddle" {
			t.Errorf("expecting riddle got %s", header)
		}
		query := request.URL.Query().Get("hera")
		if query != "aax" {
			t.Errorf("expecting aax got %s", query)
		}
		writer.Write([]byte("Yolo"))
	})
	mux.HandleFunc("GET /nobody", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Yolo"))
	})

	server := httptest.NewServer(mux)
	tlsServer := httptest.NewTLSServer(mux)
	defer server.Close()
	defer tlsServer.Close()
	t.Log(server.URL)
	t.Log(tlsServer.URL)

	databaseClient := database.NewClient(&userDir)
	database.AutoMigrate(databaseClient)
	projectRepository := database.NewRepository[database.Project](databaseClient)
	project, err := projectRepository.Create(&database.Project{
		Name:        "test project",
		Collections: nil,
	})
	if err != nil {
		t.Fatal("should not fail to create project")
	}
	collectionsRepository := database.NewRepository[database.Collection](databaseClient)
	collection, err := collectionsRepository.Create(&database.Collection{
		Name:              "test collection",
		ProjectID:         project.ID,
		HttpRequests:      nil,
		GrpcRequests:      nil,
		WebsocketRequests: nil,
	})
	if err != nil {
		t.Fatal("should not fail to create collection")
	}
	httpRequestRepository := database.NewHttpRequestRepository(databaseClient)
	defaultConfiguration := configuration.Configuration{}
	readWriter.Write(defaultConfiguration)
	request := NewRequest(httpRequestRepository, readWriter)
	httpRequest, err := httpRequestRepository.Create(&database.HttpRequest{
		Name:         "test request",
		CollectionID: collection.ID,
		Url:          server.URL + "/nobody",
		Method:       "GET",
	})
	if err != nil {
		t.Fatal("should not fail to create http get request")
	}
	requestResponseDTO := request.Submit(httpRequest.ID)
	if requestResponseDTO.Error != "" {
		t.Fatal("should not fail to receive response of http get request")
	}

	httpRequest, err = httpRequestRepository.Create(&database.HttpRequest{
		Name:         "test request",
		CollectionID: collection.ID,
		Url:          server.URL,
		Method:       "POST",
		HttpRequestBody: database.HttpRequestBody{
			Type:    "json",
			Payload: `{"key": "value"}`,
		},
	})
	if err != nil {
		t.Fatal("should not fail to create http post request")
	}
	httpRequestHeader := &database.HttpRequestHeader{
		HttpRequestID: httpRequest.ID,
		Key:           "tom",
		Value:         "riddle",
	}

	httpRequestParameter := &database.HttpRequestParameter{
		HttpRequestID: httpRequest.ID,
		Key:           "hero",
		Value:         "aax",
	}
	_, _ = httpRequestRepository.CreateHeader(httpRequestHeader)
	_, _ = httpRequestRepository.CreateParameter(httpRequestParameter)

	httpRequest, _ = httpRequestRepository.GetById(httpRequest.ID)

	requestResponseDTO = request.Submit(httpRequest.ID)
	if requestResponseDTO.Error != "" {
		t.Fatal("should not fail to receive response of http post request")
	}

	skipTlsConfiguration := configuration.Configuration{
		Theme:         "dark",
		SkipTLSVerify: true,
	}
	readWriter.Write(skipTlsConfiguration)
	httpRequest, err = httpRequestRepository.Create(&database.HttpRequest{
		Name:         "test request",
		CollectionID: collection.ID,
		Url:          tlsServer.URL + "/nobody",
		Method:       "GET",
	})
	if err != nil {
		t.Fatal("should not fail to create http request")
	}
	requestResponseDTO = request.Submit(httpRequest.ID)
	if requestResponseDTO.Error != "" {
		t.Fatal("should not create error: ", requestResponseDTO.Error)
	}
	if requestResponseDTO.TlsSkipped == false {
		t.Fatal("tls verify should be skipped")
	}

	noSkipTlsConfiguration := configuration.Configuration{
		Theme:         "dark",
		SkipTLSVerify: false,
	}
	readWriter.Write(noSkipTlsConfiguration)
	requestResponseDTO = request.Submit(httpRequest.ID)
	if requestResponseDTO.Error == "" {
		t.Fatal("should create error")
	}
	if requestResponseDTO.TlsSkipped == true {
		t.Fatal("tls verify should not be skipped")
	}
}
