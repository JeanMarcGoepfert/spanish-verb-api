package repo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"spanish-api/models"
)

func InitDB() (data models.Verbs, err error) {
	jsonFile, err := os.Open("repo/verbs.json")
	defer jsonFile.Close()

	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	var result models.Verbs

	json.Unmarshal([]byte(byteValue), &result)

	return result, nil
}
