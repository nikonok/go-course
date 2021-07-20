package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// print body of response
func printBody(resp *http.Response) {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBody))
}

// form key for new message
const keyString = "hello-msg"
// our message
const message = "General Kenobi!"

func main() {
	// server port number
	const port = 8080
	// server address
	address := fmt.Sprintf("http://127.0.0.1:%d", port)

	fmt.Printf("Connecting to server %s\n", address)

	// get current message
	resp, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("GET request response:\t")
	fmt.Println(resp)
	printBody(resp)

	// post new message
	resp, err = http.PostForm(address, url.Values{keyString: {message}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("POST request response:\t")
	fmt.Println(resp)
	printBody(resp)

	// get message
	resp, err = http.Get(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("GET request response:\t")
	fmt.Println(resp)
	printBody(resp)
}