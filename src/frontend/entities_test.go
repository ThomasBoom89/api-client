package frontend

import (
	"api-client/src/database"
	"api-client/src/test"
	"testing"
)

func TestEntities(t *testing.T) {
	userDir := test.UserDir{Dir: "./tmp-entities_test/"}
	defer userDir.Cleanup()
	databaseClient := database.NewClient(&userDir)
	database.AutoMigrate(databaseClient)
	projectRepository := database.NewRepository[database.Project](databaseClient)
	collectionsRepository := database.NewRepository[database.Collection](databaseClient)
	httpRequestRepository := database.NewRepository[database.HttpRequest](databaseClient)
	projects := NewProjects(projectRepository)
	collections := NewCollections(collectionsRepository)
	httpRequests := NewRequests(httpRequestRepository)

	projectDtos, err := projects.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(projectDtos) != 0 {
		t.Fatal("projects should be empty")
	}

	projectDto := ProjectDto{
		Name: "NewProject1",
	}

	projectDto, err = projects.Create(projectDto)
	if err != nil {
		t.Fatal("error creating new project", err)
	}
	if projectDto.ID != 1 {
		t.Fatal("new project id should be 1")
	}

	projectDto.Name = "NewProject2"
	projectDto, err = projects.Update(projectDto)
	if err != nil {
		t.Fatal("error updating project", err)
	}
	if projectDto.Name != "NewProject2" {
		t.Fatal("new project name should be NewProject2")
	}

	projectDtos, err = projects.GetAll()
	if err != nil {
		t.Fatal("could not get projects from database", err)
	}
	if len(projectDtos) != 1 {
		t.Fatal("projects should have one project")
	}

	collectionsDtos, err := collections.GetAll()
	if err != nil {
		t.Fatal("could not get collections from database", err)
	}
	if len(collectionsDtos) != 0 {
		t.Fatal("collections should be empty")
	}
	collectionDto := CollectionDto{
		Name:      "NewCollection1",
		ProjectID: projectDto.ID,
	}

	collectionDto, err = collections.Create(collectionDto)
	if err != nil {
		t.Fatal("error creating new collection", err)
	}
	if collectionDto.ID != 1 {
		t.Fatal("new collection id should be 1")
	}
	collectionDto.Name = "NewCollection2"
	collectionDto, err = collections.Update(collectionDto)
	if err != nil {
		t.Fatal("error updating collection", err)
	}
	if collectionDto.Name != "NewCollection2" {
		t.Fatal("new collection name should be NewCollection2")
	}
	collectionsDtos, err = collections.GetAll()
	if err != nil {
		t.Fatal("could not get collections from database", err)
	}
	if len(collectionsDtos) != 1 {
		t.Fatal("collections should have one collection")
	}

	httpRequestDtos, err := httpRequestRepository.GetAll()
	if err != nil {
		t.Fatal("could not get http requests from database", err)
	}
	if len(httpRequestDtos) != 0 {
		t.Fatal("http requests should be empty")
	}
	httpRequestDto := HttpRequestDto{
		Name:         "httpRequest1",
		Type:         "http",
		CollectionID: collectionDto.ID,
		Url:          "superurl",
	}
	httpRequestDto, err = httpRequests.Create(httpRequestDto)
	if err != nil {
		t.Fatal("error creating http request", err)
	}
	if httpRequestDto.Name != "httpRequest1" {
		t.Fatal("http request name should be httpRequest1")
	}
	httpRequestDto.Url = "superurl2"
	httpRequestDto, err = httpRequests.Update(httpRequestDto)
	if err != nil {
		t.Fatal("error updating http request", err)
	}
	if httpRequestDto.Url != "superurl2" {
		t.Fatal("http request url should be superurl2")
	}
	httpRequestDtos, err = httpRequestRepository.GetAll()
	if err != nil {
		t.Fatal("could not get http requests from database", err)
	}
	if len(httpRequestDtos) != 1 {
		t.Fatal("http requests should have one request")
	}

	err = httpRequests.Delete(httpRequestDto)
	if err != nil {
		t.Fatal("error deleting http request", err)
	}
	err = collections.Delete(collectionDto)
	if err != nil {
		t.Fatal("error deleting collection", err)
	}
	err = projects.Delete(projectDto)
	if err != nil {
		t.Fatal("error deleting project", err)
	}
}
