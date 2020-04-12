package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

func main() {

	ReadAllAtOnce()
	UsingBuffers()
	MoreBuffers()
	EmulateWCL()
}

func ReadAllAtOnce() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path.")
	}
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(b))
}

func UsingBuffers() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a file.")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer f.Close() // We ensure close to avoid leaks

	var (
		b = make([]byte, 16)
	)
	for n := 0; err == nil; {
		n, err = f.Read(b)
		if err == nil {
			fmt.Println(string(b[:n])) // Only print what's been read
		}
	}
	if err != nil && err != io.EOF { // We expect an EOF
		fmt.Println("\n\nError:", err)
	}
}
func MoreBuffers() {
	var b = bytes.NewBuffer(make([]byte, 26))
	var texts = []string{
		`Lorem ipsum dolor sit amet,`,
		`consectetur adipiscing elit.`,
		`Cras non placerat ex, et placerat leo.`,
	}
	for i := range texts {
		b.Reset()
		b.WriteString(texts[i])
		fmt.Println("Length:", b.Len(), "\tCapacity:", b.Cap())
	}
}

func EmulateWCL() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path.")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer f.Close()
	r := bufio.NewReader(f) // Wrapping the reader with a buffered one
	var rowCount int
	for err == nil {
		var b []byte
		for moar := true; err == nil && moar; {
			b, moar, err = r.ReadLine()
			if err == nil {
				fmt.Println(string(b))
			}
		}
		// each time moar is false, a line is completely read
		if err == nil {
			fmt.Println()
			rowCount++
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println("\nError:", err)
	}
	fmt.Println("\nRow count:", rowCount)
}
