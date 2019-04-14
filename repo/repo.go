package repo

import (
	"spanish-api/models"
)

var DB models.Verbs

func init() {
	DB, _ = InitDB()
}

func GetVerbs() models.Verbs {
	return DB
}

func GetVerb(verb string) *models.Verb {
	return GetVerbs()[verb]
}
