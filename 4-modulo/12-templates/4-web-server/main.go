package main

import (
	"net/http"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

type Persons []Person

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Persons{
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
	})
	http.ListenAndServe(":8080", nil)
}
