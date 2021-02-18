package main

import (
	"fmt"
)

func hr() {
	fmt.Println("================================")
}

func main() {

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

	// You cannot do this
	// We will try something else
	//for i, v := range arr {
	//v[i] = 12
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
	var val int = 123

	fmt.Println(&val)

	err := pr(&val)
	fmt.Println("Is my error nil > ", err == nil)
	fmt.Println("My error > ", err)

	hr()
	fmt.Println("Mutating variable val")
	fmt.Println(val)
}

func pr(a *int) error {
	// Complex business logic to mutate a variable
	*a = 456
	fmt.Println(a == nil)
	//return nil

	//return errors.New("My custom error")

	return fmt.Errorf("The error is : %v", "Bad error")
}
