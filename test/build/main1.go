package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
	//	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Hello World!")

	type T struct {
		F int `yaml:"a,omitempty"`
		B int
	}
	var t T
	yaml.Unmarshal([]byte("a: 1\nb: 2"), &t)
	fmt.Println("123 %s\n", t)
}
