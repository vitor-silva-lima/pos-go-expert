package main

import (
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

type Persons []Person

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{"header.html", "content.html", "footer.html"}
	t := template.Must(template.New("content.html").Funcs(template.FuncMap{
		"ToUpper": ToUpper,
	}).ParseFiles(templates...))
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
