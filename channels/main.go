package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	// Creating a channel of type string
	c := make(chan string)

	for _, link := range links {

		// Enabling Go Routine
		go checkLink(link, c)

	}

	/*
		// Receiving msgs from channel
		for {
			// Call the Go Routine over the link channel has sent along with a channel
			// <-c is ultimately a string as the link passed to checkLink() is a string
			go checkLink(<-c, c)
		}

	*/

	// Same as above
	for l := range c {
		// go checkLink(l, c)

		// Function literals

		// Here, 'l' is passed from an outer loop and thus the Child and Main Routine are pointing to
		// same value

		// Receives link of type string
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)  // When 'l' is passed as an argument to function literal, child routine creates a copy
		     // of 'l' and thus both Main and Child Routines have different values

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}

	fmt.Println(link, " is up!")

	// To continuously check the status of the link
	c <- link
}
