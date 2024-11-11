package frontend

import (
	"api-client/src/database"
	"api-client/src/test"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	defer test.Cleanup()

	http.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Yolo"))
	})
	server := &http.Server{
		Addr:                         ":8898",
		Handler:                      nil,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}
	go server.ListenAndServe()

	databaseClient := database.NewClient(&test.UserDir{})
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
	httpRequestRepository := database.NewRepository[database.HttpRequest](databaseClient)
	request := NewRequest(httpRequestRepository)
	httpRequest, err := httpRequestRepository.Create(&database.HttpRequest{
		Name:         "test request",
		CollectionID: collection.ID,
		Url:          "http://localhost:8898",
		Method:       "GET",
	})
	if err != nil {
		t.Fatal("should not fail to create http request")
	}
	requestResponseDTO := request.Submit(httpRequest.ID)
	if requestResponseDTO.Error != "" {
		t.Fatal("should not fail to receive response of http request")
	}
}
