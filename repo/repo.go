package repo

import (
	"sort"
	"spanish-api/lib"
	"spanish-api/models"
)

var DB models.Verbs
var VerbList []string

func init() {
	DB, _ = InitDB()
	for verb := range DB {
		VerbList = append(VerbList, verb)
	}
	sort.Strings(VerbList)
}

func GetVerbs() models.Verbs {
	return DB
}

func GetVerb(verb string) *models.Verb {
	return GetVerbs()[verb]
}

func GetPaginatedVerbs(pageNumber int, pageSize int) []models.Meta {
	//TODO: handle out of bounds errors
	//TODO: move paginate into lib
	start := ((pageNumber - 1) * pageSize)
	end := lib.Min(start+pageSize-1, len(VerbList)-1)
	keys := VerbList[start:end]
	var result []models.Meta

	for _, key := range keys {
		v := GetVerb(key).Meta
		result = append(result, v)
	}

	return result
}
