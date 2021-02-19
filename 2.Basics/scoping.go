// Variable shadowing
// and scoping
package main

import (
	"fmt"
)

func main() {

	var a int = 333

	// This could be a function as well
	{

		a := 555
		fmt.Println("Inside Bracket > ", a)

		// If you want to change the value of 'a' that is in the main scope then use this
		// a = 999

	}

	func() {

		a := 666
		fmt.Println("Inside Function > ", a)

		// If you want to change the value of 'a' that is in the main scope then use this
		// a = 999
	}()

	fmt.Println("Main scope > ", a)

	hr()

	err := fmt.Errorf("%v", "This is my error1")

	if err := fmt.Errorf("%v", "This is my error2"); err != nil {

		fmt.Println("Err is not nil", err)

	} else {

		fmt.Println("Err is nil")

	}

	fmt.Println("We will see error1 here instead of error2 > ", err)

}

func hr() {
	fmt.Println()
	fmt.Println("==================================")
}
