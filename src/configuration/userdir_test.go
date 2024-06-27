package configuration

import "testing"

func TestGetPath(t *testing.T) {
	xdg := NewXDG()
	// currently only work on linux and if the users name is thomas
	if "/home/thomas/.config/row-api-client/" != xdg.GetPath() {
		t.Fatal("not equal")
	}
}
