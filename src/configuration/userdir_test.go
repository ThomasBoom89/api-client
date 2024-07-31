package configuration

import (
	"runtime"
	"strings"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Log("test not running on different platform than linux")
		return
	}
	xdg := NewXDG()
	// currently only working on linux and if the users name is thomas
	if !strings.Contains(xdg.GetConfigPath(), "/.config/row-api-client") {
		t.Fatal("not equal")
	}
}

func TestGetDataPath(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Log("test not running on different platform than linux")
		return
	}

	xdg := NewXDG()
	// currently only working on linux and if the users name is thomas
	if !strings.Contains(xdg.GetDataPath(), "/share/row-api-client") {
		t.Fatal("not equal")
	}
}
