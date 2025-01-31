package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("This is just a sample!"))
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
