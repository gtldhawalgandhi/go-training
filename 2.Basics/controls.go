package main

import (
	"fmt"
)

func someFn() bool {
	return true
}

func main() {

	// 1 True Brace Style syntax

	if 3 > 22 {
		fmt.Println("condition1 met")
	} else if 3 > 1 {
		fmt.Println("Else if 3 > 1")
	} else {
		fmt.Println("Printing else")
	}

	// short cut

	if a := true; a == true {
		fmt.Println("A is true")
	}

	// Short cut with a function call
	if isOk := someFn(); isOk {
		fmt.Println("is actually OK")
	}

	// For loop style 1

	for i := 0; i < 5; i++ {
		fmt.Printf("%v \t", i)
	}

	hr()

	// While loop

	var cond int
	for cond < 2 {
		cond++
		fmt.Println("My while loop in Go")
	}

	hr()

	// For looping an array or slice

	var slice []string = []string{"Loop", "Me", "Over"}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%v \t", slice[i])
	}

	hr()

	// For loop with Go style

	for index, value := range slice {
		fmt.Printf("%d::%s \t", index, value)
	}
	hr()

	// For loop with Go style without value

	for index := range slice {
		fmt.Printf("%v \t", slice[index])
	}
	hr()

	// For loop without index

	for _, value := range slice {
		fmt.Printf("%s \t", value)
	}
	hr()

	// infinite loop (kinda)

	y := 0
	for {

		if y > 0 {
			break
		}
		y++
		fmt.Println("I could be inside an infinite loop if I dont increment y and break", y)
	}

	// loop with a label

	type obj struct {
		key string
	}

	outerKeys := []string{"a", "b"}
	innerObjects := []obj{
		{key: "a"},
		{key: "b"}, // note there is a comma at the end
	}

findMyStuff:
	for _, v1 := range outerKeys {

		for _, v2 := range innerObjects {

			if v1 == v2.key {

				fmt.Println("Found the key", v2.key)
				continue findMyStuff

			}
		}

		fmt.Println("Key not found")

	}

	hr()
	// Switch Statements

	var a interface{} = [3]int{12, 42, 11}
	fmt.Printf("Switch in Go:: Type interface :%T: %[1]v \n", a)

	// Here I am switching on a.(type) and not 'a'. I just store the result in a for me to use inside 'case'
	switch a := a.(type) {

	case string:
		fmt.Println("String value is ", a)

	// no 'break' statement is required as its the default behavior
	// If you want to go to the next use fallthrough like so

	// fallthrough

	case [3]int:
		fmt.Println("Array value is [3]int", a)

	default:
		fmt.Println("sorry wrong type")
	}

	// Use multiple case values to switch on
	hr()
	var myVal int = 5

	// Here I am switching on a variable myVal but you can switch even on a function call
	// Ex: switch response := myApiCall() { .... }
	switch myVal {

	case 1, 2, 3:
		fmt.Println("Value found is in first case: ", myVal)

		// You cannot repeat values in other cases
		// Ex: you cannot write 3,4,5 because '3' is already used in above case
	case 4, 5, 6:
		fmt.Println("Value found is in second case : ", myVal)

	default:
		fmt.Println("No matching value found")
	}

	// Use switch like any other if else statement
	switch {
	case myVal < 2:
	case myVal > 2:
	default:
		fmt.Println("No match found")
	}

}

func hr() {
	fmt.Println()
	fmt.Println("==================================")
}
