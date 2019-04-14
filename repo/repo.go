package repo

import (
	"spanish-api/models"
)

var DB models.Verbs

func init() {
	DB, _ = InitDB()
}

func GetVerbs() (data models.Verbs) {
	return DB
}
