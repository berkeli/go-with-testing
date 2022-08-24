package main

import "fmt"

var prefixes = map[string]string{
	"English": "Hello",
	"Spanish": "Hola",
	"French":  "Bonjour",
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix := prefixes["English"]
	if _, ok := prefixes[lang]; ok {
		prefix = prefixes[lang]
	}
	return prefix + ", " + name + "!"
}

func main() {
	fmt.Println(Hello("World", "English"))
}
