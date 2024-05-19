package main

import (
	"fmt"
	"os"

	"github.com/pedrolzoliveira/pinhata/token"
)

func main() {
	var content string
	if buffer, error := os.ReadFile("test.pin"); error == nil {
		content = string(buffer)
	} else {
		panic(error)
	}

	tokens, error := token.Tokenize(content)
	if error != nil {
		panic(error)
	}

	fmt.Println(tokens)
}
