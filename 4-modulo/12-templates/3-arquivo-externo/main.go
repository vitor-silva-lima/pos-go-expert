package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

type Persons []Person

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Persons{
		{
			Name: "John Doe",
			Age:  30,
		},
		{
			Name: "Jane Doe",
			Age:  25,
		},
	})
	if err != nil {
		panic(err)
	}

}
