package main

import (
	"classic"
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

type game struct {
	Title      string
	Word       string
	WordUser   []string
	LengthWord int
	Attempts   int
	ToFind     []string
	Position   string
	File       string
}

type character struct {
	Name       string
	clue       int
	difficulty int
	score      int
}

type Settings struct {
	Language  []string
	Scorboard []string
	Pictures  []string
	Sound     []string
}

type base struct {
	Hangman game
	Player  character
	Set     Settings
}

var bd = base{}
var gg = game{Title: "testeeeeeee"}

var Word = classic.RandomWord("words.txt")
var data = game{
	Title: "Hangman by Léo & Nathan", Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""),
	LengthWord: len(Word), Position: "https://clipground.com/images/html5-logo-2.png",
}

func Home(w http.ResponseWriter, r *http.Request) {
	start, _ := template.ParseFiles("./Source/Web/" + "menu" + ".html")

	if r.FormValue("send") == "submit" {
		if r.FormValue("dif") == "fa" {
			bd.Hangman.File = "words.txt"
		} else if r.FormValue("dif") == "mo" {
			bd.Hangman.File = "words2.txt"
		} else if r.FormValue("dif") == "di" {
			bd.Hangman.File = "words3.txt"
		}

		bd.Player = character{Name: r.FormValue("name")}
		bd.Hangman = game{Title: "goodluck " + r.FormValue("name"), Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""), LengthWord: 5, Position: "https://clipground.com/images/html5-logo-2.png"}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	start.ExecuteTemplate(w, "menu.html", nil)
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Source/Web/" + "Hangmanweb.page" + ".tmpl")
	if r.FormValue("loser") == "submit" {
		http.Redirect(w, r, "/loser", http.StatusSeeOther)
	}
	if r.FormValue("reset") == "submit" {
		Word = classic.RandomWord("words.txt")
		data = game{
			Title: "Hangman by Léo & Nathan", Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""),
			LengthWord: len(Word), Position: "https://clipground.com/images/html5-logo-2.png",
		}
	}

	choice := classic.Upper(r.FormValue("wordletter"))
	fmt.Println(choice)

	if len(choice) == 1 {

		index := classic.Verif(data.Word, choice)

		for i := 0; i < len(index); i++ {
			data.WordUser[index[i]] = choice
		}

		if len(index) == 0 {
			data.Attempts -= 1
		} else {
			data.Attempts += doublons(data.ToFind, choice)
		}

	} else {
		if choice == data.Word {
			data.WordUser = classic.StringToList("Congrats !")
			http.Redirect(w, r, "/win", http.StatusSeeOther)

		} else if choice != data.Word && len(choice) > 1 {
			data.Attempts -= 2
		}
	}

	if choice != "" {
		data.ToFind = append(data.ToFind, choice)
	}

	if len(classic.Verif(classic.ListToString(data.WordUser), "_")) == 0 {
		println("\n\nCongrats !")
		http.Redirect(w, r, "/win", http.StatusSeeOther)
		data.WordUser = classic.StringToList("Congrats !")
	}

	if data.Attempts <= 0 {
		http.Redirect(w, r, "/loser", http.StatusSeeOther)
	}
	t.ExecuteTemplate(w, "Hangmanweb.page.tmpl", data)
}

func Loser(w http.ResponseWriter, r *http.Request) {
	ho, _ := template.ParseFiles("./Source/Web/" + "loser" + ".html")
	if r.FormValue("restart") == "submit" {
		print("teste")

		Word := "word"
		bd.Hangman = game{
			Title: "Hangman by Léo & Nathan", Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""),
			LengthWord: len(Word), Position: "https://clipground.com/images/html5-logo-2.png",
		}
		http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
	}
	ho.ExecuteTemplate(w, "loser.html", bd)
}

func Win(w http.ResponseWriter, r *http.Request) {
	start, _ := template.ParseFiles("./Source/Web/" + "win" + ".html")

	start.ExecuteTemplate(w, "win.html", nil)

}

func doublons(liste []string, choice string) int {
	for i := 0; i < len(liste); i++ {
		if liste[i] == choice {
			return -1
		}

	}
	return 0
}

func main() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/loser", Loser)
	http.HandleFunc("/win", Win)
	http.HandleFunc("/", Hangman)

	fmt.Println("http://localhost" + port + "/home")

	fs := http.FileServer(http.Dir("Source"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe("localhost:8080", nil)
}
