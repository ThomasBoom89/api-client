package database

import (
	"api-client/src/test"
	"testing"
)

func TestDatabase(t *testing.T) {
	defer test.Cleanup()
	databaseClient := NewClient(&test.UserDir{})
	AutoMigrate(databaseClient)

	projectRepository := NewRepository[Project](databaseClient)
	collectionRepository := NewRepository[Collection](databaseClient)
	httpRequestRepository := NewRepository[HttpRequest](databaseClient)

	project := &Project{
		Name:        "NewProject1",
		Collections: nil,
	}
	_, err := projectRepository.Create(project)
	if err != nil {
		t.Fatal("project was not created")
	}
	project.Name = "NewProject2"
	_, err = projectRepository.Update(project)
	if err != nil || project.Name != "NewProject2" {
		t.Fatal("project was not updated")
	}
	projects, err := projectRepository.GetAll()
	if err != nil {
		t.Fatal("projects was not retrieved")
	}
	if len(projects) != 1 {
		t.Fatal("expected 1 project")
	}

	collection := &Collection{
		Name:              "TestCollection1",
		HttpRequests:      nil,
		GrpcRequests:      nil,
		WebsocketRequests: nil,
		ProjectID:         project.ID,
	}
	_, err = collectionRepository.Create(collection)
	if err != nil {
		t.Fatal("collection was not created")
	}
	collection.Name = "NewCollection2"
	_, err = collectionRepository.Update(collection)
	if err != nil {
		t.Fatal("collection was not updated")
	}
	collections, err := collectionRepository.GetAll()
	if err != nil {
		t.Fatal("collections was not retrieved")
	}
	if len(collections) != 1 {
		t.Fatal("expected 1 collection")
	}

	httpRequest := &HttpRequest{
		Name:         "httpRequest1",
		CollectionID: collection.ID,
		Url:          "superurl",
	}
	_, err = httpRequestRepository.Create(httpRequest)
	if err != nil {
		t.Fatal("http request was not created")
	}
	httpRequest.Url = "superurl2"
	_, err = httpRequestRepository.Update(httpRequest)
	if err != nil {
		t.Fatal("http request was not updated")
	}
	httpRequests, err := httpRequestRepository.GetAll()
	if err != nil {
		t.Fatal("http request was not retrieved")
	}
	if len(httpRequests) != 1 {
		t.Fatal("expected 1 http request")
	}

	err = projectRepository.Delete(project)
	if err != nil {
		t.Fatal("project was not deleted")
	}
	err = collectionRepository.Delete(collection)
	if err != nil {
		t.Fatal("collection was not deleted")
	}
	err = httpRequestRepository.Delete(httpRequest)
	if err != nil {
		t.Fatal("http request was not deleted")
	}
}
