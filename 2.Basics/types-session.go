package main

import (
	"fmt"
)

// ENUMS (iota)
// iota always starts with value 0
const (
	Jan int = iota + 1 // 1
	Feb                // value 2
	//_                  // skipped
	Mar // 4
)

// Use iota to represent file size like so
const (
	// Use _ to skip
	_  = 1 << (iota * 10) // 1 * (2^0)
	KB                    // 1 * (2^10)
	MB                    // 1 * (2^20)
	GB                    // 1 * (2^30)
)

func typesSession() {

	fmt.Printf("Type: %T, March >  %[1]v \n", Mar)
	fmt.Printf("Type: %T, GB >  %[1]v \n", GB)

	var hello string = "Hello" + " World"

	// strings in Go are utf-8 encoded => lay out bytes
	for i := 0; i < len(hello); i++ {
		fmt.Printf("%T %v \n", hello[i], string(hello[i]))
	}

	fmt.Println(len(hello))

	hello2 := hello

	hr()

	fmt.Println(hello, hello2)
	hello = "My Hello " + "World"

	hr()

	fmt.Println(hello, " >>> ", hello2)

	hr()

	// Zero value

	var aa bool

	fmt.Println(aa)

	var bb int

	fmt.Println(bb)

	var cc string

	fmt.Println(cc)

	// Performant to use arrays
	var arr [2]int

	fmt.Println(arr)

	// Copying in array
	arr2 := arr

	arr[0] = 34
	arr[1] = 44
	fmt.Println(arr, arr2)

	// You can loop through the array and modify its values like so
	//for i := range arr {
	//  arr[i] = 12
	//}

	//hr()
	//fmt.Println(arr)

	var slc []string = []string{"Hello", "Gopher", "Slices"}

	fmt.Println(slc, len(slc), cap(slc))

	var slc2 []string

	fmt.Println(slc2, len(slc2), cap(slc2))

	slc2 = append(slc2, "Appended")
	fmt.Println(slc2, len(slc2), cap(slc2))

	slc3 := slc2

	hr()
	fmt.Println(slc2, slc3)
	hr()
	slc2[0] = "Changed"
	fmt.Println(slc2, slc3)

	// Copying a slice
	var slc4 []string = make([]string, len(slc))
	copy(slc4, slc)

	hr()
	fmt.Println(slc, slc4)

	hr()
	slc[0] = "Original"
	fmt.Println(slc, slc4)

	hr()
	slc4 = append(slc4, "Slice4")
	fmt.Println(slc, slc4)

	hr()
	var mp map[string]int = make(map[string]int, 3)
	mp["hello"] = 123
	mp["removethis"] = 444

	fmt.Printf("%#v \n", mp)

	hr()
	delete(mp, "removethis")
	fmt.Printf("%#v \n", mp)

	hr()

	// usr slice and dont pass its address if possible
	// If you are using an array then you can pass its address
	var val []int = []int{1}

	fmt.Println(val)

	err := pr(val)
	fmt.Println("Is my error nil > ", err == nil)
	fmt.Println("My error > ", err)

	hr()
	fmt.Println("Mutating variable val")
	fmt.Println(val)
}

//
func pr(a []int) error {
	// Complex business logic to mutate a variable
	// De referencing a variable
	a = []int{456}
	fmt.Println(a == nil)
	//return nil

	//return errors.New("My custom error")

	return fmt.Errorf("The error is : %v", "Bad error")
}
