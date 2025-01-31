# 2. Simple Start

* go 
    * build: Compiles a bunch of Go source code files. > go build  main.go > ./main
    * run: Compiles and executes one or two files. > go run main.go
    * fmt: Formats all the code in each file in the current directory.
    * install: Compiles and "installs" a package.
    * get: Downloads the raw source code of someone else's package.
    * test: Runs any tests associated with the current project.

* package main meaning
    * package main > main() = executable code
    * package apple > (no main func) > reusable code

[Standard library](https://pkg.go.dev/std)
* import
    * standard libraries
        * fmt, Formatting and printing output.
        * net/http, Building HTTP servers and clients.
        * encoding/json, Encoding and decoding JSON data.
        * os, Interacting with the operating system.
        * strings, Manipulating and working with strings.
    * reusable package
        * calculator
        * btc