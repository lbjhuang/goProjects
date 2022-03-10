package study

var json_data = []byte(`
	{
		"name" : "Golang",
		"say" : ["Hello", "World!"]
	}
`)

type Data struct {
	Name string `json:"name"`
	Say []string `json:"say"`
}
