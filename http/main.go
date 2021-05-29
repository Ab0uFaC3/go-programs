package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)





func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	/*
		// make a byte slice with size 99999
		bs := make([]byte, 99999)

		// Read() will read the body from resp and will store it in bs
		// If byte slice is initialized to 0 size then Read() assumes that the slice is full
		resp.Body.Read(bs)
		fmt.Println(bs)
	*/

	// This does the same work as above lines
	// os.Stdout implemets Writer Interface which writes values to the output medium
	// resp.Body implements Reader Interface which reads the response from the http url

	// Writer interface has Write() which takes []byte and outputs the contents 
	// Reader interface has Read() which copies the response to []byte
	io.Copy(os.Stdout, resp.Body)

}
