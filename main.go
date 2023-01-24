package main

import (
	"classic"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

const port = ":8080"

type game struct {
	Title        string
	Word         string
	WordUser     []string
	Last_Attemps string
	LengthWord   int
	Attempts     int
	ToFind       string
}

type UserInfo struct {
	Difficulty string
	Pseudo     string
	Score      int
}

type Board struct {
	Easy Facile

	Medium Moyen

	Hard Difficile
}

type Facile struct {
	Pseudo1 string
	Score1  int
	Pseudo2 string
	Score2  int
	Pseudo3 string
	Score3  int
}
type Moyen struct {
	Pseudo1 string
	Score1  int
	Pseudo2 string
	Score2  int
	Pseudo3 string
	Score3  int
}
type Difficile struct {
	Pseudo1 string
	Score1  int
	Pseudo2 string
	Score2  int
	Pseudo3 string
	Score3  int
}

func init_board() {
	Sb.Easy = Facile{Pseudo1: "N/A", Score1: 3, Pseudo2: "N/A", Score2: 2, Pseudo3: "N/A", Score3: 1}
	Sb.Medium = Moyen{Pseudo1: "N/A", Score1: 3, Pseudo2: "N/A", Score2: 2, Pseudo3: "N/A", Score3: 1}
	Sb.Hard = Difficile{Pseudo1: "N/A", Score1: 3, Pseudo2: "N/A", Score2: 2, Pseudo3: "N/A", Score3: 1}
}

var Sb = Board{}

func Scoreb(w http.ResponseWriter, r *http.Request) {
	start, _ := template.ParseFiles("./Web/" + "ScoreBoard.page.tmpl" + ".html")
	start.ExecuteTemplate(w, "ScoreBoard.page.tmpl.html", Sb)
}

func main() {
	init_board()
	var Joueur = UserInfo{Score: 2, Difficulty: "Hard", Pseudo: "mwa gem bi1"}
	fis := http.FileServer(http.Dir("Source"))
	http.Handle("/static/", http.StripPrefix("/static/", fis))
	Word := classic.RandomWord(os.Args[1])
	t, _ := template.ParseFiles("./Web/" + "Hangmanweb.page" + ".tmpl")
	data := game{
		Title:      "Hangman by Léo & Nathan",
		Word:       Word,
		WordUser:   classic.WordChoice(Word),
		Attempts:   10,
		ToFind:     "",
		LengthWord: len(Word),
	}
	scoreboard(&Joueur, &Sb)
	http.HandleFunc("/scoreboard", Scoreb)
	handler := func(w http.ResponseWriter, r *http.Request) {
		Input := (r.FormValue("wordletter"))
		var Nbr []int
		if Input != "" {
			Nbr = classic.Verif(data.Word, Input)
			if len(Nbr) > 0 {
				for i := 0; i < len(Nbr); i++ {
					data.WordUser[Nbr[i]] = Input
					print(data.WordUser)
				}
			}
		}
		if len(r.FormValue("wordletter")) == 1 {

			index := classic.Verif(data.Word, r.FormValue("wordletter"))

			for i := 0; i < len(index); i++ {
				data.WordUser[index[i]] = r.FormValue("wordletter")
			}
			for i := 0; i < len(index); i++ {
				data.WordUser[index[i]] = r.FormValue("wordletter")
			}

			if len(index) == 0 {
				data.Attempts -= 1
			}

		} else {
			if r.FormValue("wordletter") == data.Word {
				data.WordUser = classic.Sti("Congrats !")
			} else if r.FormValue("wordletter") != data.Word && len(r.FormValue("wordletter")) > 1 {
				data.Attempts -= 2
			}
		}

		if len(classic.Verif(classic.ListToString(data.WordUser), "_")) == 0 {
			println("\n\nCongrats !")
			data.WordUser = classic.Sti("Congrats !")
		}
		if data.Attempts <= 0 {
			data.Attempts = 0
			println("\n\nYou loose !")
			data.WordUser = classic.Sti("You loose !")
		}

		t.Execute(w, data)
	}
	http.HandleFunc("/", handler)
	fmt.Println("http://localhost" + port)
	/*	fs := http.FileServer(http.Dir("Web"))
		http.Handle("/Web/", http.StripPrefix("/Web", fs)) */
	http.ListenAndServe(port, nil)
}

func scoreboard(User *UserInfo, Scoreboard *Board) {
	switch User.Difficulty {
	case "Hard":
		if User.Score > Scoreboard.Hard.Score1 {
			Scoreboard.Hard.Score3 = Scoreboard.Hard.Score2
			Scoreboard.Hard.Pseudo3 = Scoreboard.Hard.Pseudo2
			Scoreboard.Hard.Score2 = Scoreboard.Hard.Score1
			Scoreboard.Hard.Pseudo2 = Scoreboard.Hard.Pseudo1
			Scoreboard.Hard.Score1 = User.Score
			Scoreboard.Hard.Pseudo1 = User.Pseudo
			break
		} else if User.Score < Scoreboard.Hard.Score1 && User.Score > Scoreboard.Hard.Score2 {
			Scoreboard.Hard.Pseudo3 = Scoreboard.Hard.Pseudo2
			Scoreboard.Hard.Score3 = Scoreboard.Hard.Score2
			Scoreboard.Hard.Pseudo2 = User.Pseudo
			Scoreboard.Hard.Score2 = User.Score
			break
		} else if User.Score > Scoreboard.Hard.Score3 {
			Scoreboard.Hard.Score3 = User.Score
			Scoreboard.Hard.Pseudo3 = User.Pseudo
		} else {
			break
		}
	case "Medium":
		if User.Score > Scoreboard.Medium.Score1 {
			Scoreboard.Medium.Score3 = Scoreboard.Medium.Score2
			Scoreboard.Medium.Pseudo3 = Scoreboard.Medium.Pseudo2
			Scoreboard.Medium.Score2 = Scoreboard.Medium.Score1
			Scoreboard.Medium.Pseudo2 = Scoreboard.Medium.Pseudo1
			Scoreboard.Medium.Score1 = User.Score
			Scoreboard.Medium.Pseudo1 = User.Pseudo
			break
		} else if User.Score < Scoreboard.Medium.Score2 && User.Score > Scoreboard.Medium.Score3 {
			Scoreboard.Medium.Pseudo3 = Scoreboard.Medium.Pseudo2
			Scoreboard.Medium.Score3 = Scoreboard.Medium.Score2
			Scoreboard.Medium.Pseudo2 = User.Pseudo
			Scoreboard.Medium.Score2 = User.Score
			break
		} else if User.Score > Scoreboard.Medium.Score3 {
			Scoreboard.Medium.Score3 = User.Score
			Scoreboard.Medium.Pseudo3 = User.Pseudo
		} else {
			break
		}
	case "Easy":
		if User.Score > Scoreboard.Easy.Score1 {
			Scoreboard.Easy.Score3 = Scoreboard.Easy.Score2
			Scoreboard.Easy.Pseudo3 = Scoreboard.Easy.Pseudo2
			Scoreboard.Easy.Score2 = Scoreboard.Easy.Score1
			Scoreboard.Easy.Pseudo2 = Scoreboard.Easy.Pseudo1
			Scoreboard.Easy.Score1 = User.Score
			Scoreboard.Easy.Pseudo1 = User.Pseudo
			break
		} else if User.Score < Scoreboard.Easy.Score2 && User.Score > Scoreboard.Easy.Score3 {
			Scoreboard.Easy.Pseudo3 = Scoreboard.Easy.Pseudo2
			Scoreboard.Easy.Score3 = Scoreboard.Easy.Score2
			Scoreboard.Easy.Pseudo2 = User.Pseudo
			Scoreboard.Easy.Score2 = User.Score
			break
		} else if User.Score > Scoreboard.Easy.Score3 {
			Scoreboard.Easy.Score3 = User.Score
			Scoreboard.Easy.Pseudo3 = User.Pseudo
		} else {
			break
		}
	}
	User.Score = 0
	User.Pseudo = "N/A"
	return
}
