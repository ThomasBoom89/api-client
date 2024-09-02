package database

import (
	"api-client/src/configuration"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

const (
	Filename      = "superdatabase.db"
	SqliteOptions = "?cache=shared&_fk=1"
)

func NewClient(userDir configuration.UserDir) *gorm.DB {
	err := os.MkdirAll(userDir.GetDataPath(), 0755)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create data directory")
	}
	database, err := gorm.Open(sqlite.Open("file:"+userDir.GetDataPath()+Filename+SqliteOptions), &gorm.Config{})
	if err != nil {
		log.Fatal().Msgf("failed opening connection to sqlite: %v", err)
	}

	return database
}

func AutoMigrate(databaseClient *gorm.DB) {
	err := databaseClient.AutoMigrate(
		&HttpRequest{},
		&GrpcRequest{},
		&WebsocketRequest{},
		&Collection{},
		&Project{},
	)
	if err != nil {
		log.Fatal().Msg("migration was not successful")
	}
}
