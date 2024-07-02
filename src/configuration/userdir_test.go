package configuration

import "testing"

func TestGetConfigPath(t *testing.T) {
	xdg := NewXDG()
	// currently only working on linux and if the users name is thomas
	if "/home/thomas/.config/row-api-client/" != xdg.GetConfigPath() {
		t.Fatal("not equal")
	}
}

func TestGetDataPath(t *testing.T) {
	xdg := NewXDG()
	// currently only working on linux and if the users name is thomas
	if "/home/thomas/.local/share/row-api-client/" != xdg.GetDataPath() {
		t.Fatal("not equal")
	}
}
