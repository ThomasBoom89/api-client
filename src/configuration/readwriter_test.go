package configuration

import (
	"api-client/src/test"
	"reflect"
	"testing"
)

func TestWriteRead(t *testing.T) {
	defer test.Cleanup()
	userDirTest := &test.UserDir{}
	readWriter := NewReadWriter(userDirTest)
	expectedConfiguration := Configuration{Theme: "dark"}
	err := readWriter.Write(expectedConfiguration)
	if err != nil {
		t.Fatal("no error should be returned:", err)
	}
	currentConfiguration, err := readWriter.Read()
	if err != nil {
		t.Fatal("no error should be returned:", err)
	}
	if !reflect.DeepEqual(expectedConfiguration, currentConfiguration) {
		t.Fatal("Expected configuration to be equal", expectedConfiguration, currentConfiguration)
	}
}
