package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// fmt.Println(cards)

	// Assignment
	var numbers []int
	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
