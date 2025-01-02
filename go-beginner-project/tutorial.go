package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	// player name
	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	// player age
	fmt.Printf("What is you age: ")
	var age uint
	fmt.Scan(&age)
	fmt.Println(age >= 10)

	// 37:53
}
