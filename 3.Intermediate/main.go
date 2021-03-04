// Package main provides ...
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	l "github.com/gtldhawalgandhi/go-training/3.Intermediate/logger"
)

func helo() {
	l.D("Helo debug")
	l.F("My Fatal example")
}

func simpleGETClient() {

	myClient := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := myClient.Get("https://httpbin.org/get")
	if err != nil {
		// Try figure what error type it is and take decisions based on that
		// Dont throw fatal errors if not required
		l.F(err)
	}

	// var nB = 20 (read 20 bytes)
	// var nB = res.ContentLength
	// var p = make([]byte, nB)
	// n, err := res.Body.Read(p)

	// if err != nil {
	// 	l.F(err)
	// }
	// defer res.Body.Close()
	// fmt.Println("Bytes read", n)
	// fmt.Println("Bytes read data", string(p))

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		l.F(err)
	}
	defer res.Body.Close()
	// fmt.Println(string(data))

	var mp = make(map[string]interface{})
	json.Unmarshal(data, &mp)
	fmt.Println(mp)
}

func simplePOSTRequest() {

	myClient := http.Client{
		Timeout: 5 * time.Second,
	}
	data := map[string]interface{}{
		"Gopher": "Message from someone",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		l.F(err)
	}

	payload := bytes.NewReader(jsonData)

	req, err := http.NewRequest("POST", "https://httpbin.org/post", payload)
	// jsonData2, err := json.Marshal(data)
	if err != nil {
		l.F(err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := myClient.Do(req)
	if err != nil {
		l.F(err)
	}

	readerData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		l.F(err)
	}
	defer res.Body.Close()
	fmt.Println(string(readerData))
}

// -trimpath will cut short file name everywhere in our code when displaying
// go build -trimpath -o app && ./app
// OR
// go run -trimpath .
func main() {
	// l.SetFileLogger("myLog.txt", l.TRACE)
	// defer l.CleanUp()
	// l.I("Entering main file")

	// l.T("My Trace data")
	// l.E("My error")
	// helo()

	simpleGETClient()
	simplePOSTRequest()

}
