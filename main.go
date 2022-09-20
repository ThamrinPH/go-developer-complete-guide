package main

import (
	"fmt"
	"io"
	"os"
)

/*
Create a program that reads the contents of a text file then prints
its contents to the terminal.

The file to open should be provided as an argument to the program
when it is executed at the terminal. For example, 'go run main.go myfile.txt'
should open up the myfile.txt file

to read in the arguments provided to a program,
you can reference the variable 'os.Args', which is a slice of type string

To open a file, check out the documentation for the "Open" function
in the "os" package - https://golang.org/pkg/os/#Open

What interfaces does the "File" type implement?

if the "File" type implements the "Reader" interface, you might be able
to reuse the io.Copy function
*/
func main() {
	filePath := os.Args[1]
	data, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, data)
}
