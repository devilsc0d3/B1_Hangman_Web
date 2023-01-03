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
	X          string
}

type menu struct {
	Title string
	Word  string
}

func Home(w http.ResponseWriter, r *http.Request) {
	men := menu{Title: "testeu", Word: ""}
	start, _ := template.ParseFiles("./Source/Web/" + "menu" + ".html")
	if r.FormValue("dif") == "fa" {
		print("facile")
		men.Word = "words.txt"

	} else if r.FormValue("dif") == "mo" {
		print("moyen")
		men.Word = "words2.txt"

	} else if r.FormValue("dif") == "di" {
		print("difficile")
		men.Word = "words3.txt"

	}
	if r.FormValue("dif") != "" {
		Hangman(&men)
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

	start.ExecuteTemplate(w, "menu.html", men)
}

func Hangman(men *menu) {
	t, _ := template.ParseFiles("./Source/Web/" + "Hangmanweb.page" + ".tmpl")

	Word := classic.RandomWord(men.Word)

	data := game{
		Title:      "Hangman de :",
		Word:       classic.Upper(Word),
		WordUser:   classic.WordChoice(Word),
		Attempts:   10,
		ToFind:     classic.StringToList(""),
		LengthWord: len(Word),
		X:          "./static/pictures/imgte.png",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println(*&men.Word)

		if r.FormValue("reset") == "submit" {
			Word = classic.RandomWord("Words.txt")
			data = game{
				Title: "Hangman by Léo & Nathan", Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""),
				LengthWord: len(Word), X: "https://clipground.com/images/html5-logo-2.png",
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
			data.WordUser = classic.StringToList("Congrats !")
		}

		if data.Attempts <= 0 {
			http.Redirect(w, r, "/loser", http.StatusSeeOther)
		}
		t.ExecuteTemplate(w, "Hangmanweb.page.tmpl", data)
	})
}

func Loser() {
	start, _ := template.ParseFiles("./Source/Web/" + "loser" + ".html")
	me := menu{Title: "testeu"}
	http.HandleFunc("/loser", func(w http.ResponseWriter, r *http.Request) {
		start.ExecuteTemplate(w, "loser.html", me)
	})

}

func Win() {
	start, _ := template.ParseFiles("./Source/Web/" + "win" + ".html")
	ma := menu{Title: "testeu"}
	http.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		start.ExecuteTemplate(w, "win.html", ma)
	})

}

func doublons(liste []string, choice string) int {
	for i := 0; i < len(liste); i++ {
		if liste[i] == choice {
			fmt.Println("conno")
			return -1
		}

	}
	return 0
}

func main() {
	http.HandleFunc("/home", Home)
	Loser()
	Win()

	fmt.Println("http://localhost" + port + "/home")

	fs := http.FileServer(http.Dir("Source"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe("localhost:8080", nil)
}
