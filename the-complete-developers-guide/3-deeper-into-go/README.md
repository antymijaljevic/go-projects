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