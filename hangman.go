package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

type game struct {
	Title    string
	MotBase  string
	Mot      string
	Attempts int
	ToFind   string
}

func main() {
	server := http.NewServeMux()
	t, _ := template.ParseFiles("./templates/" + "hangman" + ".tmpl")

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := game{
			Title:    "Hangman by Léo & Nathan",
			MotBase:  "lavabo",
			Mot:      "_ _ _ _ _ _",
			Attempts: 0,
			ToFind:   "",
		}
		t.Execute(w, data)
	})

	fmt.Println("http://localhost" + port)
	http.ListenAndServe(port, nil)
}
