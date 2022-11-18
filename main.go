package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

type game struct {
	Title      string
	Word       string
	WordUser   string
	LengthWord int
	Attempts   int
	ToFind     string
}

func main() {
	t, _ := template.ParseFiles("./Web/" + "Hangmanweb.page" + ".tmpl")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Word := "lavabo"
		data := game{
			Title:      "Hangman by Léo & Nathan",
			Word:       Word,
			WordUser:   HiddenWord(Word),
			Attempts:   0,
			ToFind:     "",
			LengthWord: len(Word),
		}

		fmt.Println(r.FormValue("wordletter"))
		t.Execute(w, data)
	})

	fmt.Println("http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func HiddenWord(Word string) string {
	var hiddenword string
	for i := 0; i < len(Word); i++ {
		if i == len(Word)-1 {
			hiddenword += "_"
		} else {
			hiddenword += "_ "
		}
	}
	return hiddenword
}
