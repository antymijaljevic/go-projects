package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func AnimalSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	AnimalSound(dog)
	AnimalSound(cat)
}
