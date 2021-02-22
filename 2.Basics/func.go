package main

import "fmt"

func fn1() string {
	return "fn1"
}

// named return
func fn2() (res string) {
	res = "This is a response"

	// notice you dont return with a variable
	return
}

// Multiple return values (notice the return statement is wrapped within brackets *important)
func fn3() (int, string) {

	return 13, "Regular"
}

// Multiple named returns
func fn4() (age int, name string) {
	age = 90
	name = "named value"

	return
}

func fn5(s string) {
	fmt.Println("fn5 > ", s)
}

// 0 or more values (also know as variadic arguments)
func fn6(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// defer in Go is called only when the function is about to return
// Notice the sequence in which these defer functions get executed
func fn7(s string) {
	defer func() {
		fmt.Println("DEFER:1: , Value of s from defer > ", s)
	}()

	fmt.Println("Value of s: 1 > ", s)
	s = "Inside fn7"
	fmt.Println("Value of s: 2 > ", s)

	defer func() { fmt.Println("DEFER:2") }()
	defer func() { fmt.Println("DEFER:3") }()
}

func fn8() {
	defer func() {
		// Try and get the error during crash
		if err := recover(); err != nil {
			fmt.Println("From recover ", err)
			// Log somewhere for your team members to see (centralized place)
		}
	}()

	panic("Some error happened 123")

	// What happens if you uncomment this line?
	// Will this call DEFER:2 and DEFER:3?
	//panic("Lets crash out program")

	defer func() { fmt.Println("DEFER:2") }()
	defer func() { fmt.Println("DEFER:3") }()
}

// Closures in Go
func fn9(arg string) {
	local := "Local to fn9"
	func() {
		fmt.Println("inside fn: arg ", arg)
		fmt.Println("inside fn: local ", local)
		local = "Changed from inside closure"
	}()
	fmt.Println("End of dn9 > ", local)
}

type person struct {
	firstName, lastName string
}

// Methods are just like functions we have seen before except they are bound to a type (person in this case)
func (p *person) fullName() string {
	return p.firstName + p.lastName
}

func main() {
	fmt.Println(fn1())
	fmt.Println(fn2())
	fmt.Println(fn3())
	fmt.Println(fn4())

	fn5("fn5 called")

	fn6()
	fn6("fn6 called with arg1", "fn6 called with arg2")
	fn6([]string{"Str1", "Str2"}...) // converted a slice into variadic arguments (spread operator)

	fn7("Calling fn7")
	fn8()
	fn9("Calling fn9")
}
