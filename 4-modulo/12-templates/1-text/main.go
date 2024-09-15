package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "John Doe",
		Age:  30,
	}
	templateString := "Name: {{.Name}}\nAge: {{.Age}}"
	tmpl, err := template.New("person").Parse(templateString)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}
}
