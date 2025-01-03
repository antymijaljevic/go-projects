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

	// age check
	if age >= 10 {
		fmt.Println("Yes, you can play!")
	} else {
		fmt.Println("Nah, you're age below")
		return
	}

	// score tracking
	score := 0
	num_question := 2

	// first question
	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var model string
	fmt.Scan(&answer, &model)

	if answer+" "+model == "RTX 3090" || answer+" "+model == "Rtx 3090" {
		fmt.Println("Correct!")
		score += 1
	} else if answer+" "+model == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	// second question
	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score = score + 1
	} else {
		fmt.Println("Incorrect!")
	}

	// Result
	fmt.Printf("%v, Your score: %v out of %v\n", name, score, num_question)
	percent := (float64(score) / float64(num_question)) * 100
	fmt.Printf("You scored: %v%%.\n", percent)
}
