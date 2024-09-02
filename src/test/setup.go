package test

import "os"

const Dir = "./tmp/"

type UserDir struct {
}

func (U *UserDir) GetConfigPath() string {
	return Dir
}
func (U *UserDir) GetDataPath() string {
	return Dir
}

func Cleanup() {
	err := os.RemoveAll(Dir)
	if err != nil {
		panic(err)
	}
}
