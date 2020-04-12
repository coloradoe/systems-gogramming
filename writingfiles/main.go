package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	WriteToFile()
}

func WriteToFile() {

	if len(os.Args) != 3 {
		fmt.Println("Please Specify a package")
	}
	// The second argument, the content, needs to be casted to a byte slice
	if err := ioutil.WriteFile(os.Args[1], []byte(os.Args[2]), 0644); err != nil {
		fmt.Println("Error:", err)
	}
}
