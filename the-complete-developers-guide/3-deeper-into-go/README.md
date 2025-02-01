# 3. Deeper into GO

touch main.go

#### Basic types

Type	    Description	    Examples
bool	    Represents      boolean values	true, false
string	    Represents      text	"Hi!", "Hows it going?"
int	        Represents      integer numbers	0, -10000, 99999
float64	    Represents      floating-point numbers (double precision)	10.000001, 0.00009, -100.003


#### When to use := vs var?
Go lang - Statically typed langauge

* := (Type inference)
    * When the type is obvious from the context (e.g., card := "Ace of Spades").
    * For concise code, especially in small functions or local scopes.

* var (Explicit typing)
    * When you want to explicitly declare the type for clarity.
    * When the variable is initialized later or with a zero value (e.g., var card string).


#### Run two files form the same package
go run main.go state.go 


#### Array vs Slice
In Go, an array is a fixed-size collection of elements, while a slice is a dynamically-sized, flexible view into an array or a portion of it.


### Go lang
Go is not purely procedural or purely object-oriented. It is a modern language that blends paradigms to provide a simple, efficient, and concurrent programming model. It uses procedural constructs, OOP-like features (without inheritance), and unique concurrency primitives to solve problems effectively.