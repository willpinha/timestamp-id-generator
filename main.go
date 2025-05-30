package main

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func main() {
	id := NewId()

	fmt.Println(id.Value())
	fmt.Println(id.ValueWithFormat(PrettyValueFormat))

	if err := clipboard.WriteAll(id.Value()); err != nil {
		fmt.Println("Error copying to clipboard:", err)
		return
	}

	fmt.Println("Copied to clipboard!")
}
