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
	t, _ := template.ParseFiles("./Source/Web/" + "hangman" + ".tmpl")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := game{
			Title:    "Hangman by Léo & Nathan",
			MotBase:  "lavabo",
			Mot:      "_ _ _ _ _ _",
			Attempts: 0,
			ToFind:   "",
		}
		fmt.Println(r.FormValue("wordletter"))
		t.Execute(w, data)
	})

	fmt.Println("http://localhost" + port)
	http.ListenAndServe(port, nil)
}
