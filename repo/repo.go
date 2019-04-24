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

func GetVerbCount() int {
	return len(VerbList)
}

func GetPaginatedVerbs(pageNumber int, pageSize int) []models.Meta {
	result := []models.Meta{}

	max := len(VerbList)
	start := (pageNumber - 1) * pageSize

	if start >= max || start < 0 {
		return result
	}

	end := lib.Min(start+pageSize, max-1)

	keys := VerbList[start:end]

	for _, key := range keys {
		v := GetVerb(key).Meta
		result = append(result, v)
	}

	return result
}
