package main

import (
	"fmt"
)

func main() {
	fmt.Println("First line")
	defer fmt.Println("Second line")
	fmt.Println("Third line")
}
