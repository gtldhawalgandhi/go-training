package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type data struct {
	Name           string   `json:"name"`
	Age            int      `json:"age,omitempty"` // omitempty is for marshal only
	Address        *Address `json:"address"`
	MyPointerValue *string  `json:"my_pointer_value"` // notice the type is a pointer, helps to figure whether value is present and not a zero value
	Meta           int      `json:"-"`                // Marshal and Unmarshal (ignore this property)
}

func (a *Address) String() string {
	return fmt.Sprintf("%#v", a)
}

func createJSONString() {
	mypointer := "this is from pointer"
	d := &data{
		Name:           "Gopher",
		Meta:           234,
		MyPointerValue: &mypointer,
		Address: &Address{
			Street:  "Street 123",
			State:   "GJ",
			Country: "IN",
		},
	}

	//v, err := json.Marshal(d)
	v, err := json.MarshalIndent(d, "", "  ")

	fmt.Println("JSON Data > ", string(v), err)
}

func getFromJSONString() {
	jsn := `
		 {
				"name": "Benny",
				"address": {
						"street": "benny street 9099",
						"state"	: "MH",
						"country": "IN"
				}				
		 }`

	var a data
	json.Unmarshal([]byte(jsn), &a)

	// When we use pointer values, we no longer get zero value, instead we compare its value against nil
	fmt.Printf("%#v \n", a)
	fmt.Println("Is MyPointerValue nil? > ", a.MyPointerValue == nil)

	// JSON does not have Meta value, but its value in our struct is "0"
	fmt.Println("Meta value from json > ", a.Meta)
}

func main() {
	createJSONString()
	getFromJSONString()

	type data struct {
		Val int `json:"value"`
	}

	d := data{
		Val: 99,
	}

	b, err := json.Marshal(d)

	fmt.Println("This is my data >>  ", string(b), err)

	jsn := `
			{
				"value": 777
			}
	`
	var dt data
	err = json.Unmarshal([]byte(jsn), &dt)

	fmt.Println("Err UNmarshalling", err)

	fmt.Printf("Unmarshalled data %#v \n", dt)
}
