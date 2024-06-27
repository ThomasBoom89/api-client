package configuration

import "github.com/adrg/xdg"

const (
	ApplicationPath = "row-api-client"
	PathSeparator   = "/"
)

type UserDir interface {
	GetPath() string
}
type XDG struct {
}

func NewXDG() *XDG {
	return &XDG{}
}

func (X *XDG) GetPath() string {
	return xdg.ConfigHome + PathSeparator + ApplicationPath + PathSeparator
}
