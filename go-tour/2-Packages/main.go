package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// import "fmt"
// import "math"

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7)) // %g format specifier is used for floating-point formatting.
	fmt.Println(math.Pi)                                    // a name is exported if it begins with a capital letter

	fmt.Println(add(133, 22))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!" // If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	fmt.Println(i, j, c, python, java)

	var u, p int = 1, 2
	k := 3 // Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	golang, js, r := true, false, "no!"
	fmt.Println(u, p, k, r, golang, js, r)

	// go basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var (
		l int
		f float64
		e bool
		s string
	)
	fmt.Printf("%v %v %v %q\n", l, f, e, s) // Variables declared without an explicit initial value are given their zero value.

	// Type conversions
	var i1 int = 42
	var f2 float64 = float64(i1)
	var u1 uint = uint(f2)

	fmt.Println(i1, f2, u1)

	// or
	// i := 42
	// f := float64(i)
	// u := uint(f)

	var x3, y3 int = 3, 4
	var f3 float64 = math.Sqrt(float64(x3*x3 + y3*y3))
	var z3 uint = uint(f3)
	fmt.Println(x3, y3, z3)

	// https: //go.dev/tour/basics/14
}

func add(x int, y int) int { // or func add(x, y int)
	return x + y
}

func swap(x, y string) (string, string) { // A function can return any number of results.
	return x, y
}

func split(sum int) (x, y int) { // Named return values. This is known as a "naked" return. Only used with short functions
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i, j int = 1, 2

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

/*
When to Use Each

	Use var:
	•	When declaring variables at the package level.
	•	When you need to specify the type explicitly.
	•	When you want to declare variables without initializing them.

	Use :=:
	•	For local variable declarations within functions.
	•	When you want concise and readable code.
	•	When the type can be easily inferred from the initial value.
*/

/*

Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
