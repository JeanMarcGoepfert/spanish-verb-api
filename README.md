### Spanish Verb API

### Install dependencies

`go get -u ./...`

### Useage

Get information about a specific spanish verb
```
/verb/:verb
```


Get paginated list of verbs (sorted alphabetically)
```
/verbs?page_number=x&page_number=y
```

### Development

Running tests:
`go test ./...`

Fetch all dependencies:
`go get -u ./...`

Adding a new dependency:
`go get -u <repo url>`
