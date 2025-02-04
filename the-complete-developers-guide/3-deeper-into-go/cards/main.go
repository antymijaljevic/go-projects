package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// fmt.Println(cards)

	// Assignment
	var ints []int
	for i := 0; i < 11; i++ {
		ints = append(ints, i)
	}

	for num := range ints {
		if num == 0 {
			fmt.Printf("%v is invalid number\n", num)
		} else if num%2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
