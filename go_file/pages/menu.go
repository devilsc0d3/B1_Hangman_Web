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
		if r.FormValue("dif") == "fa" {
			File = "words.txt"
		} else if r.FormValue("dif") == "mo" {
			File = "words2.txt"
		} else if r.FormValue("dif") == "di" {
			File = "words3.txt"
		} else {
			File = "words.txt"
		}

		Bd.Set.Name = r.FormValue("name")
		if r.FormValue("name") == "" {
			Bd.Set.Name = "Gertrude"
		}

		var Word = classic.RandomWord(File)
		Bd.Hangman = game{ClueNbr: 0, File: File, Title: "Good luck " + r.FormValue("name"), Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""), LengthWord: 5, Position: "https://clipground.com/images/html5-logo-2.png"}
		if File == "words3.txt" {
			Bd.Hangman.ClueNbr = 1
		}
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
