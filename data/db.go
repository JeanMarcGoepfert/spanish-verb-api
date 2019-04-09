package data

type Verb struct {
	Base string `json:"base"`
}

var DB = map[string]Verb{
	"hablar":   {Base: "hablar"},
	"aprender": {Base: "aprender"},
}
