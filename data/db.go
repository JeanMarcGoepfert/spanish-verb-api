package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Forms struct {
	Form1s string `json:"form_1s"`
	Form2s string `json:"form_2s"`
	Form3s string `json:"form_3s"`
	Form1p string `json:"form_1p"`
	Form2p string `json:"form_2p"`
	Form3p string `json:"form_3p"`
}

type Tenses struct {
	Imperfect          *Forms `json:"imperfect,omitempty"`
	PresentPerfect     *Forms `json:"present_perfect,omitempty"`
	Conditional        *Forms `json:"conditional,omitempty"`
	PastPerfect        *Forms `json:"past_perfect,omitempty"`
	Preterite          *Forms `json:"preterite,omitempty"`
	FuturePerfect      *Forms `json:"future_perfect,omitempty"`
	ConditionalPerfect *Forms `json:"conditional_perfect,omitempty"`
	Present            *Forms `json:"present,omitempty"`
	Future             *Forms `json:"future,omitempty"`
}

type Moods struct {
	Indicative            Tenses `json:"indicative"`
	Subjunctive           Tenses `json:"subjunctive"`
	ImperativeAffirmative Tenses `json:"imperative_affirmative"`
	ImperativeNegative    Tenses `json:"imperative_negative"`
}

type Meta struct {
	Infinitive            string `json:"infinitive"`
	InfinitiveEnglish     string `json:"infinitive_english"`
	VerbEnglish           string `json:"verb_english"`
	Gerund                string `json:"gerund"`
	GerundEnglish         string `json:"gerund_english"`
	PastParticible        string `json:"pastparticiple"`
	PastParticibleEnglish string `json:"pastparticiple_english"`
}

type Usage struct {
	Usage Moods `json:"useage"`
	Meta  Meta  `json:"meta"`
}

type Data map[string]*Usage

func GetVerbs() Data {
	jsonFile, err := os.Open("data/verbs.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result Data

	json.Unmarshal([]byte(byteValue), &result)

	return result
}
