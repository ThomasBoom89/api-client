package configuration

import "github.com/adrg/xdg"

const (
	ApplicationPath = "row-api-client"
	PathSeparator   = "/"
)

type UserDir interface {
	GetConfigPath() string
	GetDataPath() string
}
type XDG struct {
}

func NewXDG() *XDG {
	return &XDG{}
}

func (X *XDG) GetConfigPath() string {
	return xdg.ConfigHome + PathSeparator + ApplicationPath + PathSeparator
}

func (X *XDG) GetDataPath() string {
	return xdg.DataHome + PathSeparator + ApplicationPath + PathSeparator
}
