package main

import (
	"classic"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
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
	ClueNbr    int
}

type Language struct {
	Fr []string
	En []string
	Es []string
	Ge []string
}
type Settings struct {
	Language        Language
	CurrentLanguage []string
	Name            string
	CurrentPage     string
}

type Score struct {
	Top   int
	Name  string
	Score int
}

type Board struct {
	Tab     []Score
	Tab2    []Score
	Player1 Score
	Player2 Score
	Player3 Score
	Player4 Score
	Player5 Score
	Player6 Score
	Player7 Score
	Player8 Score
	Player9 Score
}

type base struct {
	Hangman    game
	Set        Settings
	Scoreboard Board
}

var bd = base{}

func variable() {
	InitScoreboard()
	FixTop(bd.Scoreboard.Tab)
	HighScore()

	bd.Set.Language.Fr = []string{"Pac-Hangman Adventure", "facile", "moyen", "difficile", "entre un nom", "lancer",
		"Bonne chance ", "Vous avez", "essaies", "entrez une lettre ou un mot", "envoye", "lettre deja essayer : ", "rejouer",
		"tu as Gagne", "tableaux des scores", "rejoué",
		"Retour", "Parametre", "Regles", "classement", "Indices ?", "Plus", "le mot etait",
	}
	bd.Set.Language.En = []string{"Pac-Hangman Adventure", "easy", "medium", "hard", "enter a name", "start",
		"good luck", "you have", "tries", "enter a letter or word", "sent", "letter already tried:", "replay",
		"you won", "scoreboard", "replay",
		"Back", "Setting", "Rules", "Hight Score", "Clues ?", "More", "The word was",
	}
	bd.Set.Language.Es = []string{"Pac-Hangman Adventure", "facil", "medio", "dificil", "Introduce un apodo", "iniciar",
		"buena suerte", "tienes", "intentos", "introduce una letra o palabra", "enviado", "letra ya intentada : ", "reproducir",
		"has ganado", "marcador", "reproducir",
		"Atras", "parametro", "Reglas", "clasificacion", "Pistas ?", "Mas", "La palabra era",
	}
	bd.Set.Language.Ge = []string{"Pac-Hangman Adventure", "leicht", "mittel", "schwer", "einen Namen eingeben", "starten",
		"Viel Gluck", "Sie haben", "Versuche", "Geben Sie einen Buchstaben oder ein Wort ein", "Gesendet", "Buchstabe bereits versucht:", "Wiederholen",
		"Sie haben gewonnen", "Anzeigetafel", "Wiederholung",
		"Zuruck", "Einstellung", "Regeln", "Hohe Punktzahl", "Hinweise?", "Mehr", "Das Wort war",
	}
	bd.Set.CurrentLanguage = bd.Set.Language.En

	var Word = classic.RandomWord("words.txt")
	var data = game{
		Title: "...", Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""),
		LengthWord: len(Word), Position: "https://clipground.com/images/html5-logo-2.png", File: "word.txt", ClueNbr: 0,
	}
	bd.Hangman = data
}

func Home(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "menuv1" + ".html")

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

		bd.Set.Name = r.FormValue("name")
		if r.FormValue("name") == "" {
			bd.Set.Name = "Gertrude"
		}

		var Word = classic.RandomWord(File)
		bd.Hangman = game{ClueNbr: 0, File: File, Title: "goodluck " + r.FormValue("name"), Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""), LengthWord: 5, Position: "https://clipground.com/images/html5-logo-2.png"}
		if File == "words3.txt" {
			bd.Hangman.ClueNbr = 1
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.FormValue("setting") == "submit" {
		http.Redirect(w, r, "/setting", http.StatusSeeOther)
	}

	if r.FormValue("rules") == "submit" {
		bd.Set.CurrentPage = "/home"
		http.Redirect(w, r, "/rules", http.StatusSeeOther)
	}

	if r.FormValue("scores") == "submit" {
		bd.Set.CurrentPage = "/home"
		http.Redirect(w, r, "/scoreboard", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "menuv1.html", bd)
	if err != nil {
		return
	}
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Source/Web/" + "hangman1.0" + ".html")

	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", 303)
	}

	if r.FormValue("back") == "submit" {
		http.Redirect(w, r, "/home", 303)
	}

	if r.FormValue("rules") == "submit" {
		bd.Set.CurrentPage = "/"
		http.Redirect(w, r, "/rules", 303)
	}

	if r.FormValue("clue") == "submit" {
		Clue(bd.Hangman.File, bd.Hangman.Word, bd.Hangman.WordUser)
		if len(classic.Verif(classic.ListToString(bd.Hangman.WordUser), "_")) == 0 {
			bd.Scoreboard.Tab = append(bd.Scoreboard.Tab, Score{Score: bd.Hangman.Attempts * 100, Name: bd.Set.Name, Top: 0})
			bd.Scoreboard.Tab2 = append(bd.Scoreboard.Tab2, Score{Score: bd.Hangman.Attempts * 100, Name: bd.Set.Name, Top: 0})
			SortingSelection(bd.Scoreboard.Tab)
			FixTop(bd.Scoreboard.Tab)
			HighScore()
			http.Redirect(w, r, "/win", 303)
		}
	}

	choice := classic.Upper(r.FormValue("wordletter"))

	if len(choice) == 1 {
		index := classic.Verif(bd.Hangman.Word, choice)
		for i := 0; i < len(index); i++ {
			bd.Hangman.WordUser[index[i]] = choice
		}
		if len(index) == 0 {
			bd.Hangman.Attempts -= 1
		} else {
			bd.Hangman.Attempts += doublons(bd.Hangman.ToFind, choice)
		}
	} else {
		if choice == bd.Hangman.Word && len(choice) == len(bd.Hangman.Word) {
			bd.Scoreboard.Tab = append(bd.Scoreboard.Tab, Score{Score: bd.Hangman.Attempts * 100, Name: bd.Set.Name, Top: 0})
			bd.Scoreboard.Tab2 = append(bd.Scoreboard.Tab2, Score{Score: bd.Hangman.Attempts * 100, Name: bd.Set.Name, Top: 0})

			SortingSelection(bd.Scoreboard.Tab)
			FixTop(bd.Scoreboard.Tab)
			HighScore()
			http.Redirect(w, r, "/win", 303)
		} else if choice != bd.Hangman.Word && len(choice) > 1 {
			bd.Hangman.Attempts -= 2
		}
	}

	if choice != "" {
		bd.Hangman.ToFind = append(bd.Hangman.ToFind, choice)
	}

	if (len(choice) == 1) && (len(classic.Verif(classic.ListToString(bd.Hangman.WordUser), "_")) == 0) {
		bd.Scoreboard.Tab = append(bd.Scoreboard.Tab, Score{Score: bd.Hangman.Attempts * 100, Name: bd.Set.Name, Top: 0})
		bd.Scoreboard.Tab2 = append(bd.Scoreboard.Tab2, Score{Score: bd.Hangman.Attempts * 100, Name: bd.Set.Name, Top: 0})

		SortingSelection(bd.Scoreboard.Tab)
		FixTop(bd.Scoreboard.Tab)
		HighScore()
		http.Redirect(w, r, "/win", 303)
	}
	if bd.Hangman.Attempts <= 0 {
		bd.Hangman.Attempts = 10
		bd.Scoreboard.Tab = append(bd.Scoreboard.Tab, Score{Score: 0, Name: bd.Set.Name, Top: 0})
		bd.Scoreboard.Tab2 = append(bd.Scoreboard.Tab2, Score{Score: 0, Name: bd.Set.Name, Top: 0})
		SortingSelection(bd.Scoreboard.Tab)
		FixTop(bd.Scoreboard.Tab)
		HighScore()
		http.Redirect(w, r, "/loser", 303)
	}

	err := t.ExecuteTemplate(w, "hangman1.0.html", bd)
	if err != nil {
		return
	}
}

func Loser(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "loser1.0" + ".html")

	if r.FormValue("restart") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	if r.FormValue("scoreboard") == "submit" {
		bd.Set.CurrentPage = "/loser"
		http.Redirect(w, r, "/scoreboard", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "loser1.0.html", bd)
	if err != nil {
		return
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "404" + ".html")
	if r.FormValue("home") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "404.html", bd)
	if err != nil {
		return
	}
}

func Rules(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "rules" + ".html")
	if r.FormValue("back") == "submit" {
		http.Redirect(w, r, bd.Set.CurrentPage, http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "rules.html", bd)
	if err != nil {
		return
	}
}

func Win(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "win1.0" + ".html")
	if r.FormValue("restart") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	if r.FormValue("scoreboard") == "submit" {
		bd.Set.CurrentPage = "/win"
		http.Redirect(w, r, "/scoreboard", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "win1.0.html", bd)
	if err != nil {
		return
	}
}

func Setting(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "setting" + ".html")
	if r.FormValue("lg") == "en" {
		bd.Set.CurrentLanguage = bd.Set.Language.En
	}
	if r.FormValue("lg") == "fr" {
		bd.Set.CurrentLanguage = bd.Set.Language.Fr
	}
	if r.FormValue("lg") == "es" {
		bd.Set.CurrentLanguage = bd.Set.Language.Es
	}
	if r.FormValue("lg") == "ge" {
		bd.Set.CurrentLanguage = bd.Set.Language.Ge
	}
	if r.FormValue("apply") == "submit" || r.FormValue("back") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "setting.html", bd)
	if err != nil {
		return
	}
}

func doublons(array []string, choice string) int {
	for i := 0; i < len(array); i++ {
		if array[i] == choice {
			return -1
		}
	}
	return 0
}

func Ranking(w http.ResponseWriter, r *http.Request) {
	start, _ := template.ParseFiles("./Source/Web/" + "scores1.0" + ".html")
	if r.FormValue("back") == "submit" {
		http.Redirect(w, r, bd.Set.CurrentPage, http.StatusSeeOther)
	}
	if r.FormValue("more") == "submit" {
		http.Redirect(w, r, "/board", http.StatusSeeOther)
	}
	err := start.ExecuteTemplate(w, "scores1.0.html", bd)
	if err != nil {
		return
	}
}

func Scoreboard(w http.ResponseWriter, r *http.Request) {
	start, _ := template.ParseFiles("./Source/Web/" + "board" + ".html")
	if r.FormValue("back") == "submit" {
		http.Redirect(w, r, "/scoreboard", http.StatusSeeOther)
	}
	err := start.ExecuteTemplate(w, "board.html", bd)
	if err != nil {
		return
	}
}

func InitScoreboard() {
	var array []Score
	for i := 1; i <= 9; i++ {
		array = append(array, Score{Top: i, Name: "N/A", Score: 0})
	}
	bd.Scoreboard.Tab = array
}

func SortingSelection(array []Score) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if array[j].Score >= array[minIndex].Score {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}

func FixTop(array []Score) {
	for i := 0; i < len(array); i++ {
		array[i].Top = i + 1
	}
}

func HighScore() {
	bd.Scoreboard.Player1 = bd.Scoreboard.Tab[0]
	bd.Scoreboard.Player2 = bd.Scoreboard.Tab[1]
	bd.Scoreboard.Player3 = bd.Scoreboard.Tab[2]
	bd.Scoreboard.Player4 = bd.Scoreboard.Tab[3]
	bd.Scoreboard.Player5 = bd.Scoreboard.Tab[4]
	bd.Scoreboard.Player6 = bd.Scoreboard.Tab[5]
	bd.Scoreboard.Player7 = bd.Scoreboard.Tab[6]
	bd.Scoreboard.Player8 = bd.Scoreboard.Tab[7]
	bd.Scoreboard.Player9 = bd.Scoreboard.Tab[8]
}

func Clue(difficulty string, word string, word2 []string) {
	rand.Seed(time.Now().UnixNano())
	if difficulty == "words3.txt" && bd.Hangman.ClueNbr >= 1 {
		index := classic.Verif(classic.ListToString(word2), "_")
		nbr := rand.Intn(len(index))
		bd.Hangman.WordUser[index[nbr]] = string(word[index[nbr]])
		bd.Hangman.ClueNbr -= 1
	}
}

func InitRankingBoard() {
	var array []Score
	array = append(array, Score{Top: -1, Name: "RV", Score: 999})
	array = append(array, Score{Top: -1, Name: "smith", Score: 900})
	array = append(array, Score{Top: -1, Name: "Chat chat", Score: 800})
	array = append(array, Score{Top: -1, Name: "Luc1ll3", Score: 700})
	array = append(array, Score{Top: -1, Name: "Leo", Score: 600})
	array = append(array, Score{Top: -1, Name: "REMI", Score: 500})
	array = append(array, Score{Top: -1, Name: "ADAN", Score: 200})

	bd.Scoreboard.Tab2 = array

	array = append(array, Score{Top: -1, Name: "N/A", Score: 0})
	array = append(array, Score{Top: -1, Name: "N/A", Score: 0})

	bd.Scoreboard.Tab = array
	FixTop(bd.Scoreboard.Tab)
	HighScore()

}

func main() {
	variable()
	InitRankingBoard()
	http.HandleFunc("/home", Home)
	http.HandleFunc("/404", NotFound)
	http.HandleFunc("/rules", Rules)
	http.HandleFunc("/setting", Setting)
	http.HandleFunc("/loser", Loser)
	http.HandleFunc("/win", Win)
	http.HandleFunc("/scoreboard", Ranking)
	http.HandleFunc("/board", Scoreboard)
	http.HandleFunc("/", Hangman)

	fmt.Println("http://localhost" + port + "/home")

	fs := http.FileServer(http.Dir("Source"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
