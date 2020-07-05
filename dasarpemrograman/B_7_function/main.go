package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Superhero{
			Name:    "Bruce Wayne",
			Alias:   "Batman",
			Friends: []string{"CatWoman", "Robin", "Superman"},
		}
		var tmp1 = template.Must(template.ParseFiles("view.html"))
		if err := tmp1.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
