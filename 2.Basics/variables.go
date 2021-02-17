package main

// If we want to use a functionality implemented by someone else then we do that by importing their packages here
// Go comes with a lot of built-in functionality wrapped in their own packages
// One way to import packages is by wrapping all imports inside parathesis.
import (
	// 'fmt' is the most widely used package provided by Go that allows to format your code 🙂
	// https://golang.org/pkg/fmt/
	"fmt"
	// 💡 If you import a package then you must use it, otherwise your code will not compile
	// Uncomment the line below to see
	// "bytes"
)

// Another way of importing packages is by doing line by line imports although not recommended style
// import "fmt"
// import "someOtherPackage"

/*
	// ✨  VARIABLES OUTSIDE FUNCTION
	We can declare variables here too

	* This is NOT possible here 🥺
	name := "Golang

	* This IS possible here as long as its not declared anywhere else
	var name string = "Golang"

	* When variable name is capitalized, other packages can import them 😉
	In other words you are exporting your variables
	var Name string
*/

// Variables declared here are not required by the compiler to be used by your application, unlike the ones declared inside function
// Uncommenting the line below WILL NOT yield any error
// var Age int

// Basics is capitalized and hence its now exported
// 💡 Its a good practise to have comments above all exported identifiers
// This way `go doc` can help those who import your package
func workWithVariablesConstants() {
	// ✨ VARIABLES INSIDE FUNCTION
	// Value inside variable can change overtime during the lifecycle of your application
	//

	// var varName type
	// Can be used inside and outside function declarations
	var name string
	var age int

	name = "Golang"
	age = 11

	// Once you declare a variable inside a function declaration you must use it, otherwise your code will not compile
	// If you comment the line below then your code will not compile
	fmt.Println(name, age)

	// 1. 🧪
	// You can declare and assign value to a variable at the same time
	// This format can be used both inside and outside a function
	// Since we have already declared the variable above,
	// If you uncomment the lines below then your code will not compile
	// For this to work you will need to comment out the code above
	/*
		var name string = "Golang"
		var age int = 10
	*/

	// 2. 🧪
	// This format can only be used inside function
	/* name := "Golang"
	age := 11
	*/

	// You can create multiple variables in one line
	var path, url string = "/blog", "https://golang.org"

	// 🧪 Alternatively you can create like this
	// path, url := "/blog", "https://golang.org"

	// As always you have to use/consume variables if you create them
	fmt.Println(path, url)

	// You can overwrite the values with similar syntax
	// Notice there is no ':' used,
	path, url = "/doc", "https://golang.org"

	// But what if you dont want to use variables❓
	// For that you can use '_'
	_, url = "/somePage", "https://google.com"
	fmt.Println(url)

	// ✨ CONSTANTS
	// constants are similar to variables except you can only define once during compile time
	// and its value cannot change during runtime/application lifecycle
	const myConst string = "THIS WILL NEVER CHANGE"
}
