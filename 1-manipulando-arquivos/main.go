package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create file
	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}

	// Write file
	// file.WriteString("Hello, World!")
	file.Write([]byte("Hi, my name is Vitor Gabriel Silva Lima!"))
	file.Close()

	// Read file
	readFile, err := os.ReadFile("./file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(readFile))

	// Read File chunks (buffer)
	file, err = os.Open("./file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 4)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	file.Close()

	// Remove file
	err = os.Remove("./file.txt")
	if err != nil {
		panic(err)
	}
}
