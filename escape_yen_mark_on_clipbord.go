package main

import (
	"github.com/atotto/clipboard"
	"strings"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	text = strings.Replace(text, "\\", "\\\\", -1)
	clipboard.WriteAll(text)
}
