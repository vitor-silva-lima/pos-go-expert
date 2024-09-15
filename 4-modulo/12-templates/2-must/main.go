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
	t := template.Must(template.New("person").Parse("Name: {{.Name}}\nAge: {{.Age}}"))
	err := t.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}

}
