package main

import (
	"fmt"
	"os"
)

/*
	go doc <package-name> :: go doc fmt (see fmt package document on your terminal)

	go get golang.org/x/tools/cmd/godoc to run locally
	godoc -http=:4444 ( now go to localhost:4444 on your browser to see docs)
*/

var s string = "My var value"

// Cant use this syntax outside your function defintion
// Only inside your function
//s := "Outside"

func main() {

	//var a, b string = "A", "B"
	a, b, c := "A", "B", "C"
	// logic

	fmt.Println(a, b, c)

	fmt.Println("Hello")

	//s := "Some string"
	//s = "My other value"
	fmt.Printf("%v \n", s)

	fmt.Printf("%d \n", 23)
	fmt.Printf("Go Syntax: %#v \n", s)
	fmt.Printf("Type: %T \n", s)
	fmt.Printf("Quotes: %q \n", s)

	r := [3]rune{'a', 'b', 'c'}
	fmt.Printf("Rune Go Syntax: %#v \n", r)
	fmt.Printf("Rune value: %v \n", r)
	fmt.Printf("Rune quote: %q \n", r)

	st := struct{ name string }{name: "Hello"}
	fmt.Printf("%#v \n", st)
	fmt.Printf("Struct value:%v \n", st)

	// F stands for writer. A writer (from io.Writer package) is a built in Go interface that is satisfied by os.Stdout and os.Stderr
	fmt.Fprintln(os.Stderr, "Lets print an error on OS ERROR file descriptor")

	str := fmt.Sprintln("Also a formatting tool but ... it returns a string instead of printing on terminal")

	fmt.Println(str)

	dec := 2.345
	fmt.Printf("%f %.2f \n", dec, dec)

	fmt.Printf("|%5d|%5d| \n", 23, 44)

}
