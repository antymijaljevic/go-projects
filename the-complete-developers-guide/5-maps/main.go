package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b7888",
		"white": "#ffffff",
	}

	printMap(colors)

	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	// colors := make(map[int]string)
	// colors[10] = "#ffffff"
	// delete(colors, 10)

	// fmt.Println(colors)
}
