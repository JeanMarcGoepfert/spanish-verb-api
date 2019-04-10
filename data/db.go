package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetVerbs() map[string]interface{} {
	jsonFile, err := os.Open("data/verbs.json")
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	return result
}
