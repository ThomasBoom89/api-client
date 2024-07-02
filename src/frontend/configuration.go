package frontend

import "api-client/src/configuration"

type Configuration struct {
	readwriter *configuration.ReadWriter
}

func NewConfiguration(readwriter *configuration.ReadWriter) *Configuration {
	return &Configuration{readwriter}
}

func (C *Configuration) Get() (configuration.Configuration, error) {
	return C.readwriter.Read()
}

func (C *Configuration) Set(configuration configuration.Configuration) error {
	return C.readwriter.Write(configuration)
}
