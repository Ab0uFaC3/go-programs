package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Read the arguments provided to the program
	arguments := os.Args[1]
	//fmt.Println(arguments)

	// Open a file
	file, err := os.Open(arguments)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Read the contents of the file
	io.Copy(os.Stdout, file)

}
