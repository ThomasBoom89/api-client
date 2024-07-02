package configuration

import (
	"os"
	"reflect"
	"testing"
)

type UserDirTest struct {
}

const (
	TestDir = "./tmp/"
)

func (U *UserDirTest) GetConfigPath() string {
	return TestDir
}
func (U *UserDirTest) GetDataPath() string {
	return TestDir
}

func TestWriteRead(t *testing.T) {
	userDirTest := &UserDirTest{}
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
	cleanup()
}

func cleanup() {
	err := os.RemoveAll(TestDir)
	if err != nil {
		panic(err)
	}
}
