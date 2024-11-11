package configuration

import (
	"api-client/src/test"
	"reflect"
	"testing"
)

func TestWriteRead(t *testing.T) {
	userDir := test.UserDir{Dir: "./tmp-readwriter_test/"}
	defer userDir.Cleanup()
	readWriter := NewReadWriter(&userDir)
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
