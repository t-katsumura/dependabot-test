package main

import (
	"fmt"

	"github.com/justinas/alice"
	"gopkg.in/yaml.v3"
)

func marshal(m map[string]any) {
	b, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {
	m := map[string]any{
		"foo": "FFF",
		"bar": "BBB",
	}
	marshal(m)
	alice.New()
}
