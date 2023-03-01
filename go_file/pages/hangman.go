package pages

import (
	"classic"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func Hangman(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../Source/Web/" + "hangman" + ".html")

	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
	}

	if r.FormValue("back") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

	if r.FormValue("rules") == "submit" {
		Bd.Set.CurrentPage = "/"
		http.Redirect(w, r, "/rules", http.StatusSeeOther)
	}

	if r.FormValue("clue") == "submit" {
		Clue(Bd.Hangman.Word, Bd.Hangman.WordUser)
		if len(classic.Checking(classic.ListToString(Bd.Hangman.WordUser), "_")) == 0 {
			Bd.Scoreboard.Tab = append(Bd.Scoreboard.Tab, Score{Score: Bd.Hangman.Attempts * 100, Name: Bd.Set.Name, Top: 0})
			Bd.Scoreboard.Tab2 = append(Bd.Scoreboard.Tab2, Score{Score: Bd.Hangman.Attempts * 100, Name: Bd.Set.Name, Top: 0})
			SortingSelection(Bd.Scoreboard.Tab)
			FixTop(Bd.Scoreboard.Tab)
			HighScore()
			http.Redirect(w, r, "/win", http.StatusSeeOther)
		}
	}

	choice := classic.Upper(r.FormValue("Word-letter"))

	if len(choice) == 1 {
		index := classic.Checking(Bd.Hangman.Word, choice)
		for i := 0; i < len(index); i++ {
			Bd.Hangman.WordUser[index[i]] = choice
		}
		if len(index) == 0 {
			Bd.Hangman.Attempts -= 1
		} else {
			Bd.Hangman.Attempts += multiLetterFound(Bd.Hangman.ToFind, choice)
		}
	} else {
		if choice == Bd.Hangman.Word && len(choice) == len(Bd.Hangman.Word) {
			Bd.Scoreboard.Tab = append(Bd.Scoreboard.Tab, Score{Score: Bd.Hangman.Attempts * 100, Name: Bd.Set.Name, Top: 0})
			Bd.Scoreboard.Tab2 = append(Bd.Scoreboard.Tab2, Score{Score: Bd.Hangman.Attempts * 100, Name: Bd.Set.Name, Top: 0})

			SortingSelection(Bd.Scoreboard.Tab)
			FixTop(Bd.Scoreboard.Tab)
			HighScore()
			http.Redirect(w, r, "/win", 303)
		} else if choice != Bd.Hangman.Word && len(choice) > 1 {
			Bd.Hangman.Attempts -= 2
		}
	}

	if choice != "" {
		Bd.Hangman.ToFind = append(Bd.Hangman.ToFind, choice)
	}

	if (len(choice) == 1) && (len(classic.Checking(classic.ListToString(Bd.Hangman.WordUser), "_")) == 0) {
		Bd.Scoreboard.Tab = append(Bd.Scoreboard.Tab, Score{Score: Bd.Hangman.Attempts * 100, Name: Bd.Set.Name, Top: 0})
		Bd.Scoreboard.Tab2 = append(Bd.Scoreboard.Tab2, Score{Score: Bd.Hangman.Attempts * 100, Name: Bd.Set.Name, Top: 0})

		SortingSelection(Bd.Scoreboard.Tab)
		FixTop(Bd.Scoreboard.Tab)
		HighScore()
		http.Redirect(w, r, "/win", http.StatusSeeOther)
	}
	if Bd.Hangman.Attempts <= 0 {
		Bd.Hangman.Attempts = 10
		Bd.Scoreboard.Tab = append(Bd.Scoreboard.Tab, Score{Score: 0, Name: Bd.Set.Name, Top: 0})
		Bd.Scoreboard.Tab2 = append(Bd.Scoreboard.Tab2, Score{Score: 0, Name: Bd.Set.Name, Top: 0})
		SortingSelection(Bd.Scoreboard.Tab)
		FixTop(Bd.Scoreboard.Tab)
		HighScore()
		http.Redirect(w, r, "/loser", http.StatusSeeOther)
	}

	err := t.ExecuteTemplate(w, "hangman.html", Bd)
	if err != nil {
		return
	}
}

func multiLetterFound(array []string, choice string) int {
	for i := 0; i < len(array); i++ {
		if array[i] == choice {
			return -1
		}
	}
	return 0
}

func Clue(word string, word2 []string) {
	rand.Seed(time.Now().UnixNano())
	if Bd.Hangman.ClueNbr >= 1 {
		index := classic.Checking(classic.ListToString(word2), "_")
		nbr := rand.Intn(len(index))
		Bd.Hangman.WordUser[index[nbr]] = string(word[index[nbr]])
		Bd.Hangman.ClueNbr -= 1
	}
}
