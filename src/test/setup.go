package test

import "os"

type UserDir struct {
	Dir string
}

func (U *UserDir) GetConfigPath() string {
	return U.Dir
}
func (U *UserDir) GetDataPath() string {
	return U.Dir
}

func (U *UserDir) Cleanup() {
	err := os.RemoveAll(U.Dir)
	if err != nil {
		panic(err)
	}
}
