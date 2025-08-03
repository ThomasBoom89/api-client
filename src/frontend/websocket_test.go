package frontend

import (
	"api-client/src/app"
	"api-client/src/database"
	"api-client/src/test"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestWebsocket(t *testing.T) {
	userDir := test.UserDir{Dir: "./tmp-websocket_test/"}
	defer userDir.Cleanup()

	mux := http.NewServeMux()
	mux.Handle("/echo", websocket.Handler(func(conn *websocket.Conn) {
		io.Copy(conn, conn)
	}))

	server := httptest.NewServer(mux)
	defer server.Close()
	re := regexp.MustCompile("^http://(.*)$")
	serverUrl := re.FindStringSubmatch(server.URL)[1]

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
	websocketRepository := database.NewWebsocketRequestRepository(databaseClient)
	websocketRequest := NewWebsocketRequests(websocketRepository)
	websocketRequestDto, err := websocketRequest.Create(WebsocketRequestDto{
		Name:         "websocket request 1",
		CollectionID: collection.ID,
		Url:          "ws://" + serverUrl + "/echo",
	})
	appContext := app.NewContext()
	receiver := make(chan []interface{})
	testEvent := test.NewEvent(receiver)
	websocketHandler := NewWebsocket(appContext, testEvent)
	websocketRequestState := websocketHandler.Connect(websocketRequestDto)
	if websocketRequestState.Error != "" || websocketRequestState.Connected == false {
		t.Fatal("could not connect to websocket server")
	}
	websocketRequestState = websocketHandler.Send(websocketRequestDto, "hello world")
	if websocketRequestState.Error != "" || websocketRequestState.Connected == false {
		t.Fatal("could not connect to websocket server")
	}
	message := <-receiver
	websocketRequestState = message[0].(WebsocketStateDto)
	if websocketRequestState.IncomingMessage != "hello world" {
		t.Fatal("did not receive echo message")
	}

	websocketRequestState = websocketHandler.Disconnect(websocketRequestDto)
	if websocketRequestState.Error != "" || websocketRequestState.Connected != false {
		t.Fatal("could not disconnect")
	}
}
