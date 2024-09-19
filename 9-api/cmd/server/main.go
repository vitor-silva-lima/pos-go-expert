package main

import (
	"first-api/configs"
	"fmt"
)

func main() {
	config := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
