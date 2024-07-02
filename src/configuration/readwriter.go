package configuration

import (
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
	"path/filepath"
)

const (
	ConfigFileName = "config.toml"
)

// Configuration json annotation is mandatory to export fields in javascript
type Configuration struct {
	Theme string `json:"theme"`
}

type ReadWriter struct {
	userDir UserDir
}

func NewReadWriter(userDir UserDir) *ReadWriter {
	return &ReadWriter{userDir}
}

func (R *ReadWriter) Read() (Configuration, error) {
	log.Println("bla")
	configuration := Configuration{}
	file, err := R.getFile()
	if err != nil {
		return configuration, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return configuration, err
	}
	buffer := make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return configuration, err
	}

	err = toml.Unmarshal(buffer, &configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}

func (R *ReadWriter) Write(configuration Configuration) error {
	log.Println("test write", configuration)
	file, err := R.getFile()
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := toml.Marshal(configuration)
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (R *ReadWriter) getFile() (*os.File, error) {
	path := R.userDir.GetConfigPath() + ConfigFileName
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
