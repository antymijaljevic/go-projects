package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// a type can be somewhat analogous to a class in Python, but with some key differences. In Go, you define a type (often using a struct) to group related data, similar to how a class in Python
// method defined only at package level
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"♠", "♥", "♦", "♣"}
	cardVales := []string{"J", "Q", "K"}

	// Create card numbers
	for i := 1; i < 11; i++ {
		cardVales = append(cardVales, strconv.Itoa(i))
	}

	// create cards
	for _, suit := range cardSuits {
		for _, v := range cardVales {
			card := suit + v
			cards = append(cards, card)
		}
	}

	return cards
}

// 'deck' any var with type of deck gets access to this receiver func
// 'd' the actual copy of the deck we're working with. cards==d
// a receiver function in Go is similar to a method in Python.
// no return type is needed because this method does not return any value
func (d deck) print() {
	for _, card := range d {
		// fmt.Printf("%v. %v\n", i+1, card)
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// to one single string with comma separation
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// stat -f "%Sp %OLp" cards_saved | to check permissions
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s) // nedded to use methos later
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		np := r.Intn(len(d) - 1)
		d[1], d[np] = d[np], d[i] //indices swap
	}
}
