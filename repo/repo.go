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
	//TODO: move paginate into lib
	//TODO: handle negative numbers
	//TODO: handle last page better (empty array if out of range)

	max := (len(VerbList) - 1)
	start := lib.Min(max-1, ((pageNumber - 1) * pageSize))
	end := lib.Min(start+pageSize, len(VerbList)-1)
	keys := VerbList[start:end]

	var result []models.Meta

	for _, key := range keys {
		v := GetVerb(key).Meta
		result = append(result, v)
	}

	return result
}
