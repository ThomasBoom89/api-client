package database

import (
	"gorm.io/gorm"
	"strconv"
)

func SeedDatabase(databaseClient *gorm.DB) {
	projects := make([]Project, 100)
	for iter := range 100 {
		projects[iter].Name = "Project " + strconv.Itoa(iter)
		projects[iter].Collections = make([]Collection, 50)
		for jiter := range 50 {
			projects[iter].Collections[jiter].Name = "Collection " + strconv.Itoa(jiter)
			projects[iter].Collections[jiter].HttpRequests = make([]HttpRequest, 50)
			for kiter := range 50 {
				projects[iter].Collections[jiter].HttpRequests[kiter] = HttpRequest{
					Name:         "httpRequest" + strconv.Itoa(kiter),
					CollectionID: uint(jiter),
					Url:          "url" + strconv.Itoa(kiter),
				}
			}
		}
	}
	db := databaseClient.Session(&gorm.Session{CreateBatchSize: 100})

	tx := db.Create(&projects)
	if err := tx.Error; err != nil {
		panic(err)
	}
}
