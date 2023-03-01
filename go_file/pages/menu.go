package pages

import (
	"classic"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("../Source/Web/" + "menu" + ".html")

	File := ""
	if r.FormValue("send") == "submit" {
		if r.FormValue("difficulty") == "easy" {
			File = "words" + Bd.Set.CurrentLanguage[23] + ".txt"
			Bd.Hangman.ClueNbr = 0
		} else if r.FormValue("difficulty") == "medium" {
			File = "words" + Bd.Set.CurrentLanguage[23] + "2" + ".txt"
			Bd.Hangman.ClueNbr = 0
		} else if r.FormValue("difficulty") == "hard" {
			File = "words" + Bd.Set.CurrentLanguage[23] + "3" + ".txt"
			Bd.Hangman.ClueNbr = 1
		} else {
			File = "words.txt"
			Bd.Hangman.ClueNbr = 0
		}

		Bd.Set.Name = r.FormValue("name")
		if r.FormValue("name") == "" {
			Bd.Set.Name = "R0B1"
		}

		//var Word = classic.RandomWord("../Source/txt/" + File)
		//Bd.Hangman = game{ClueNbr: 0, File: File, Title: "Good luck " + r.FormValue("name"), Word: classic.Upper(Word),
		//	WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList("")}

		var Word = classic.RandomWord("../Source/txt/" + File)
		Bd.Hangman.Attempts = 10
		Bd.Hangman.Title = "Good luck " + r.FormValue("name")
		Bd.Hangman.Word = classic.Upper(Word)
		Bd.Hangman.WordUser = classic.WordChoice(Word)
		Bd.Hangman.ToFind = classic.StringToList("")

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.FormValue("setting") == "submit" {
		http.Redirect(w, r, "/setting", http.StatusSeeOther)
	}

	if r.FormValue("rules") == "submit" {
		Bd.Set.CurrentPage = "/home"
		http.Redirect(w, r, "/rules", http.StatusSeeOther)
	}

	if r.FormValue("scores") == "submit" {
		Bd.Set.CurrentPage = "/home"
		http.Redirect(w, r, "/scoreboard", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "menu.html", Bd)
	if err != nil {
		return
	}

}
