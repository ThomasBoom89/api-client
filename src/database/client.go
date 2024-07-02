package database

import (
	"api-client/src/configuration"
	"api-client/src/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	Filename      = "superdatabase.db"
	SqliteOptions = "?cache=shared&_fk=1"
)

func NewClient(userDir configuration.UserDir) *ent.Client {
	err := os.MkdirAll(userDir.GetDataPath(), 0755)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create data directory")
	}
	client, err := ent.Open("sqlite3", "file:"+userDir.GetDataPath()+Filename+SqliteOptions)
	if err != nil {
		log.Fatal().Msgf("failed opening connection to sqlite: %v", err)
	}

	return client
}
