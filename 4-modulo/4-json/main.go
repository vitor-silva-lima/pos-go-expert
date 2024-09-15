package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{
		Name: "Vitor",
		Age:  20,
	}
	personJson, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(personJson))
	err = json.NewEncoder(os.Stdout).Encode(person)
	if err != nil {
		panic(err)
	}

	personJsonPure := []byte(`{"name":"Gabriel","age":30}`)
	var personPure Person
	err = json.Unmarshal(personJsonPure, &personPure)
	if err != nil {
		panic(err)
	}
	fmt.Println(personPure)
}
