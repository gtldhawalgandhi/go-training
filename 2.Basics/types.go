package main

import (
	"errors"
	"fmt"
)

func workWithTypes() {
	workWithBoolean()
	workWithError()
	workWithInteger()
	workWithString()
	workWithArray()
	workWithSlice()
	workWithMap()
	workWithStruct()
	workWithInterface()
}
func workWithInteger() {
	fmt.Println(" 📚 Integers")
	// Zero value for integer type is 0
	// var zeroIntVal int

	// If you are running Go program on 32 bit machine then int is 32 bit
	// If you are running Go program on 64 bit machine then int is 64 bit
	// Go compiler will adjust the memory allocation accordingly
	// 2^64 = 18446744073709551616 -> -9223372036854775807 to +9223372036854775807
	var i int = 9223372036854775807
	fmt.Println("This is int > ", i)

	// Max value is 2^8 = 256 -> -127 to +127
	var j int8 = 123
	fmt.Println("This is int8 > ", j)

	// Max value is 2^16 = 65536 -> -32767 to +32767
	var k int16 = 32767
	fmt.Println("This is int16 > ", k)

	// Max value is 2^32 = 4294967296 -> -2147483647 to +2147483647
	var l int32 = -2147483647
	fmt.Println("This is int32 > ", l)

	// Max value is 2^64 = 18446744073709551616 -> -9223372036854775808 to +9223372036854775808
	var m int64 = -9223372036854775808
	fmt.Println("This is int64 > ", m)

	// We also have unsigned integer values
	// Its similar to int values except it starts with 0
	// 2 ^ 64 =  -> 0 to 18446744073709551615
	var ui uint = 18446744073709551615
	fmt.Println("This is uint > ", ui)
}

func workWithBoolean() {
	fmt.Println(" 📚 Boolean")
	// Here we create a variable without any initial value and so Go will assign a 'zero' value to it
	// zero values for bool is 'false'
	var zeroValBool bool
	fmt.Println("Zero value for bool > ", zeroValBool)

	var ok bool = true
	fmt.Println("This is bool > ", ok)
}

func workWithError() {
	fmt.Println(" 📚 Error Type")
	// Error type can be a nil type or non-nil type
	var nilError error
	fmt.Println("Zero value error > ", nilError)
	fmt.Println("Is my error nil > ", nilError == nil)
	hr()

	var myError error = errors.New("My Custom error")
	fmt.Println("This is my non nil error > ", myError)
	fmt.Println("Is my error nil > ", myError == nil)
	fmt.Printf("My error type: %T > \n", myError)
	hr()

	fmtErr := fmt.Errorf("%s", "My formatted error")
	fmt.Println("This is my non nil fmt error > ", fmtErr)
	fmt.Println("Is my fmt error nil > ", fmtErr == nil)
	fmt.Printf("My error type: %T > \n", fmtErr)
}

func workWithString() {
	fmt.Println(" 📚 String")

	// Strings in Go are all unicode or utf-8 encoded
	// Also strings in Go are immutable sequence of characters or runes

	// Zero value for string is an empty string ""

	// 💡 You CANNOT use single quotes for strings in Go
	// This is invalid
	// var myVar string = 'Single quotes are characters and not string'

	var doubleQuoteStr = "I am inside double quotes"
	var backtickStr = `I am inside backtick`
	var multiLineStr = `Hello,
						I think you
						will like this`
	_ = multiLineStr
	fmt.Printf("doubleQuoteStr %v |  backtickStr %v \n", doubleQuoteStr, backtickStr)
	fmt.Printf("Type of doubleQuoteStr %T |  Type of backtickStr %T \n", doubleQuoteStr, backtickStr)
	fmt.Printf("Length of doubleQuoteStr %v |  Length of backtickStr %v \n", len(doubleQuoteStr), len(backtickStr))

	fmt.Printf("Combine strings > %v \n", doubleQuoteStr+backtickStr)
	concatStr := fmt.Sprintf("%v%v \n", doubleQuoteStr, backtickStr)
	fmt.Printf("More efficient concatenation >  %v \n", concatStr)
	// For more complex and efficient string concatenation we will use 'strings.Builder' in the future

	// I can iterate over my string
	gopherGreeting := "Hello Gophers"

	// We will cover 'for' loop again below
	for i := 0; i < len(gopherGreeting); i++ {
		// %c is for char formatting
		fmt.Printf("%c \n", gopherGreeting[i])

		// What happens if you use %v❓
		// fmt.Printf("%v \n", str[i])
		// fmt.Printf("%T \n", str[i])
	}

	// I can pick individual character like this
	fmt.Printf("Pick a character from string > %c \n", gopherGreeting[0])

	// Ok, but how do I change a single character in a string❓
	// Well, you cannot do this 🤨
	// gopherGreeting[1] = 'a'
	// Instead we convert string to []byte ([]byte will covered below)
	gopherBytes := []byte(gopherGreeting)
	gopherBytes[1] = 'a'
	fmt.Printf("Overwritten string > %s \n", gopherBytes)
	hr()
}

func workWithArray() {
	fmt.Println(" 📚 Arrays")
	var zeroArray [3]int
	fmt.Println("Zero value for array", zeroArray)

	// Array of size 8
	var arr = [8]int{12, 423, 11, 53, 16, 44, 76, 23}
	// Array of size 2
	var strArr = [2]string{"Hello", "Gopher"}
	// Array with no data
	var noDataArr = [100]int{}

	fmt.Println("Arrays > ", arr, strArr, noDataArr)
	fmt.Println("Size of Arrays > ", len(arr), len(strArr), len(noDataArr))
	fmt.Println("Last element of arr > ", arr[len(arr)-1])

	// Copying arrays
	hr()
	arrCopy := arr
	// Change the value at the first index position
	arrCopy[0] = 0
	fmt.Println("Copied array ", arrCopy)
	fmt.Println("Original array ", arr)
	hr()

	// What if you dont know the length of array in advance
	myArr := [...]string{"Hey", "there"}
	fmt.Println("Array with unknown length at declaration: ", myArr, len(myArr))
	hr()

	var twoDArray = [3][1]string{
		{"First"},  // index 0 of the outer array
		{"Second"}, // index 1 of the outer array
		{"Third"},  // index 2 of the outer array
	}

	// I can write the same thing in slightly verbose manner
	/*
		var twoDArray = [3][1]string{
			[1]string{"Hello"}, // index 0 of the outer array
			[1]string{"Second"}, // index 1 of the outer array
			[1]string{"Third"}, // index 2 of the outer array
		}
	*/
	fmt.Println("2 dimentional array > ", twoDArray)

	/*
		There are items in the outer array and one item in the inside array
	*/
	arr3 := twoDArray[2][0]
	fmt.Println("get the first element of inside array from the last outer array > ", arr3)
	hr()
}
func workWithSlice() {
	fmt.Println(" 📚 Slices")
	// Slices are higher abstraction of arrays
	// We usually use slices more than arrays because it is more convenient
	// Infact slices are just pointers to an underslying array only

	// When you declare a slice variable then it gets initialized with zero values like so
	var zeroValueSlice []int
	fmt.Println("Zero value for slice", zeroValueSlice)
	fmt.Println("Zero value length for slice", len(zeroValueSlice))
	fmt.Println("Zero value capacity for slice", cap(zeroValueSlice))

	var mySlice []string
	// What happens if you run this code below?
	// mySlice[0] = "My new slice"
	// It will panic because our newly created slice has length 0 and we are trying to access index position in a slice that has now data

	mySlice = append(mySlice, "SliceItemOne")
	fmt.Println("Slice items > ", mySlice)
	fmt.Println("Slice length > ", len(mySlice))
	fmt.Println("Slice capacity > ", cap(mySlice))
	// This is now possible, because we now have some data in our slice at index 0
	mySlice[0] = "Replace contents at index position 1"

	premadeSlice := []string{
		"SliceItemOne",
		"SliceItemTwo", //💡 notice you need a comma after the last item
	}
	fmt.Println("Premade Slice items > ", premadeSlice)
	fmt.Println("Premade Slice length > ", len(premadeSlice))
	fmt.Println("Premade Slice capacity > ", cap(premadeSlice))
	hr()

	var sliceWithLength = make([]int, 5)
	fmt.Println("sliceWithLength items > ", sliceWithLength)
	fmt.Println("sliceWithLength length > ", len(sliceWithLength))
	fmt.Println("sliceWithLength capacity > ", cap(sliceWithLength))
	hr()

	var sliceWithLengthAndCapacity = make([]int, 5, 8)
	fmt.Println("sliceWithLengthAndCapacity items > ", sliceWithLengthAndCapacity)
	fmt.Println("sliceWithLengthAndCapacity length > ", len(sliceWithLengthAndCapacity))
	fmt.Println("sliceWithLengthAndCapacity capacity > ", cap(sliceWithLengthAndCapacity))
	hr()

	// Append items to the slice
	// If a slice capacity is greater than its length then appending items is a fast operation
	sliceWithLengthAndCapacity = append(sliceWithLengthAndCapacity, 55)
	fmt.Println("sliceWithLengthAndCapacity single item appended > ", sliceWithLengthAndCapacity)
	hr()

	// Replace an item in a slice
	index := 0
	sliceWithLengthAndCapacity[index] = 77
	fmt.Println("sliceWithLengthAndCapacity first item replaced > ", sliceWithLengthAndCapacity)
	fmt.Println("sliceWithLengthAndCapacity second item removed from slice > ", sliceWithLengthAndCapacity)
	fmt.Println("sliceWithLengthAndCapacity length slice > ", len(sliceWithLengthAndCapacity))
	fmt.Println("sliceWithLengthAndCapacity capacity slice > ", cap(sliceWithLengthAndCapacity))
	hr()

	// Remove an item from a slice
	index = 1
	// '...' is called variadic operator, it converts a slice into individual items
	// Since the second argument in the 'append' can only be individual items we use '...' to achieve that
	sliceWithLengthAndCapacity = append(sliceWithLengthAndCapacity[:index], sliceWithLengthAndCapacity[index+1:]...)
	fmt.Println("sliceWithLengthAndCapacity second item removed from slice > ", sliceWithLengthAndCapacity)
	fmt.Println("sliceWithLengthAndCapacity length slice > ", len(sliceWithLengthAndCapacity))
	fmt.Println("sliceWithLengthAndCapacity capacity slice > ", cap(sliceWithLengthAndCapacity))
	hr()

	var newItems = []int{44, 66, 88, 99}
	sliceWithLengthAndCapacity = append(sliceWithLengthAndCapacity, newItems...)
	fmt.Println("sliceWithLengthAndCapacity 4 items appended > ", sliceWithLengthAndCapacity)
	fmt.Println("sliceWithLengthAndCapacity length slice > ", len(sliceWithLengthAndCapacity))

	// Bear in mind that every time Go has to double the capacity it has to copy all items from old underlying array into a new one with double the capacity, its an expensive process if you have large slice and if you perform this type of action too frequently
	fmt.Println("sliceWithLengthAndCapacity capacity slice > ", cap(sliceWithLengthAndCapacity), " < notice the capacity has been doubled by Go runtime (😲)")
	hr()

	// Loop through slice
	// Use whatever variables names you like, I am going to use 'i' and 'v'
	for i, v := range sliceWithLengthAndCapacity {
		if v == 0 {
			v = i
		}
		fmt.Printf("Slice item# %v in a loop > %v \n", i, v)
	}
	hr()

	// Loop through slice without 'v'
	for i := range sliceWithLengthAndCapacity {
		fmt.Printf("Slice index only in a loop > %v\n", i)
	}
	hr()

	// Loop through slice with only 'v'
	for _, v := range sliceWithLengthAndCapacity {
		fmt.Printf("Slice value only in a loop > %v\n", v)
	}
	hr()

	// 🙌  Copying slices 🙌
	// Unlike arrays this will NOT copy your slice.
	copiedSlice := sliceWithLengthAndCapacity
	copiedSlice[1] = 11
	fmt.Println("sliceWithLengthAndCapacity original slice > ", sliceWithLengthAndCapacity)
	fmt.Println("- - - - - - - - - improperly copied slice > ", copiedSlice)
	fmt.Println("copiedSlice length slice > ", len(copiedSlice))
	fmt.Println("copiedSlice capacity slice > ", cap(copiedSlice))

	var copiedSlice2 = make([]int, len(sliceWithLengthAndCapacity), cap(sliceWithLengthAndCapacity))
	n := copy(copiedSlice2, sliceWithLengthAndCapacity)
	copiedSlice2[2] = 22
	fmt.Println("sliceWithLengthAndCapacity original slice > ", sliceWithLengthAndCapacity)
	fmt.Println("- - - - - - - - - - properly copied slice > ", copiedSlice2)
	fmt.Println("copiedSlice2 length slice > ", len(copiedSlice2))
	fmt.Println("copiedSlice2 capacity slice > ", cap(copiedSlice2))
	fmt.Println("- - - - - - - - - - number of items copied > ", n)
	hr()
}

func workWithMap() {
	fmt.Println(" 📚 Maps")
	// Map is like a dictionary
	// It holds key value pair
	var myMap map[string]string

	fmt.Println("myMap zeroValue > ", myMap)
	fmt.Println("myMap zeroValue length > ", len(myMap))
	fmt.Println("myMap zeroValue is nil > ", myMap == nil)

	// Map does not have capacity like slices
	// So you cannot do this
	// fmt.Println("myMap zeroValue capcaity > ", cap(myMap))

	// What happens if you run this code now?
	/*
		myMap["Hello"] = "World"
		fmt.Println("myMap > ", myMap)
		fmt.Println("myMap length > ", len(myMap))
	*/
	// It panics because our newly created map above is nil. So now what?
	hr()
	// Well, to solve this problem you will always need to initialize your map with `make` keyword
	var mapWithMake = make(map[int]string)
	mapWithMake[12] = "Value for key 12"
	mapWithMake[13] = "Value for key 13"
	fmt.Println("mapWithMake > ", mapWithMake)
	fmt.Println("mapWithMake length > ", len(mapWithMake))
	fmt.Println("mapWithMake is nil > ", mapWithMake == nil)

	fmt.Println("deleting key from a map")
	delete(mapWithMake, 13)
	fmt.Println("mapWithMake after delete > ", mapWithMake)
	fmt.Println("mapWithMake after delete length > ", len(mapWithMake))
	fmt.Println("mapWithMake after delete is nil > ", mapWithMake == nil)

	hr()

	fmt.Println("Copying a map")
	/* DONT DO THIS: This will change both the maps
	   myCopiedMap := mapWithMake
	   myCopiedMap[12] = "New Value for key 12"
	   fmt.Println("myCopiedMap > ", myCopiedMap)
	   fmt.Println("myCopiedMap length > ", len(myCopiedMap))
	   fmt.Println("myCopiedMap is nil > ", myCopiedMap == nil)
	   fmt.Println("mapWithMake > ", mapWithMake)
	   fmt.Println("mapWithMake length > ", len(mapWithMake))
	   fmt.Println("mapWithMake is nil > ", mapWithMake == nil)
	*/

	myCopiedMap := make(map[int]string, len(mapWithMake))

	// You CANNOT do this. This is for slices only
	// _ = copy(myCopiedMap, mapWithMake)

	// For maps you will need help from 'for' loop
	// where k=key for the map and v=value for the map
	for k, v := range mapWithMake {
		myCopiedMap[k] = v
	}
	myCopiedMap[12] = "New Value for key 12"
	myCopiedMap[22] = "New Value for key 22"
	fmt.Println("myCopiedMap > ", myCopiedMap)
	fmt.Println("myCopiedMap length > ", len(myCopiedMap))
	fmt.Println("myCopiedMap is nil > ", myCopiedMap == nil)
	fmt.Println("mapWithMake > ", mapWithMake)
	fmt.Println("mapWithMake length > ", len(mapWithMake))
	fmt.Println("mapWithMake is nil > ", mapWithMake == nil)

	hr()
	fmt.Printf("Retrieve value from map > %+v \n", myCopiedMap[22])

	// What happens if you try to get value for key that does not exist?
	fmt.Printf("Retrieve value from a map for non existent key > %v \n", myCopiedMap[33])
	myCopiedMap[33] = ""
	fmt.Printf("Retrieve empty string value from a map for key > %v \n", myCopiedMap[33])
	fmt.Println(`
	Above two look exactly the same. How can you distinguish between the two?
	I cannot tell whether key has empty value or whether the key itself does not exist`)

	v, ok := myCopiedMap[44]
	if ok {
		fmt.Println("Key 44 exists in map > ", v)
	} else {
		fmt.Println("Key 44 DOES NOT exist in map ")
	}

	// You can write the same code above in a slightly succinct manner
	if v, ok := myCopiedMap[44]; ok {
		fmt.Println("Key 44 exists in map > ", v)
	} else {
		fmt.Println("Key 44 DOES NOT exist in map ")
	}

	hr()
}

func workWithStruct() {
	fmt.Println(" 📚 Struct")

	// Struct allows to put together all basic types we have learned into a single entity

	// I could have created 3 different variables to store the data and thats just fine, but struct is a much better way to organize your code
	type person struct {
		name    string
		age     int
		hobbies []string
	}
	gopher := person{
		name: "Gopher",
		age:  11,
		hobbies: []string{
			"Make developers' lives happy",
		},
	}

	fmt.Printf("I am a Gopher > %+v \n", gopher)
	fmt.Printf("Person Type> %T \n", gopher)

	hr()
}

func workWithInterface() {

	// Interface is a type in Go, that can hold value of any primitive type
	// When you are not sure what type to handle you can use interfaces

	//var a interface{} = "Hello Interface"
	// Empty interface
	var a interface{} = [3]int{12, 42, 11}
	//var a interface{} = 23.33
	fmt.Printf("Type interface :%T: %[1]v \n", a)

	switch a := a.(type) {

	case string:
		fmt.Println("String value is ", a)

	case [3]int:
		fmt.Println("Array value is [3]int", a)

	default:
		fmt.Println("sorry wrong type")
	}

	hr()
	// What I want is a value that holds a human
	// How do I do that?
	// By creating a value of myType in this example
	// Notice we doing type conversion from int > myType
	var h human = myType(89)

	fmt.Printf("Type: %T, Value: %[1]v \n", h)

	h.walk()
	h.talk()
	h.hear()

}

// Interface methods

// I have 3 methods for 'human' interface
type human interface {
	walk()
	talk()
	hear()
}

// This is my custom type (its nothing but int in this case) that is going to fulfill all 3 methods of type 'human'
type myType int

func (t myType) walk() {
	fmt.Printf("My type %T can walk with value: %[1]v \n", t)
}
func (t myType) talk() {
	fmt.Printf("My type %T can talk with value: %[1]v \n", t)
}

// if you remove hear method then Go compiler will throw error on line where we declared our 'human' variable above
func (t myType) hear() {
	fmt.Printf("My type %T can hear with value: %[1]v \n", t)
}
