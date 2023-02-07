package pages

import "classic"

type game struct {
	Title    string
	Word     string
	WordUser []string
	Attempts int
	ToFind   []string
	File     string
	ClueNbr  int
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

var Bd = base{}

func Variable() {
	//initialisation of board ranking
	InitScoreboard()
	FixTop(Bd.Scoreboard.Tab)
	HighScore()

	//initialisation of languages
	Bd.Set.Language.Fr = []string{"Pac-Hangman Adventure", "facile", "moyen", "difficile", "entre un nom", "lancer",
		"Bonne chance ", "Vous avez", "essaies", "entrez une lettre ou un mot", "envoye", "lettre deja essayer : ", "rejouer",
		"tu as Gagne", "tableaux des scores", "rejoué",
		"Retour", "Parametre", "Regles", "classement", "Indices ?", "Plus", "le mot etait",
	}
	Bd.Set.Language.En = []string{"Pac-Hangman Adventure", "easy", "medium", "hard", "enter a name", "start",
		"good luck", "you have", "tries", "enter a letter or word", "sent", "letter already tried:", "replay",
		"you won", "scoreboard", "replay",
		"Back", "Setting", "Rules", "High Score", "Clues ?", "More", "The word was",
	}
	Bd.Set.Language.Es = []string{"Pac-Hangman Adventure", "facil", "medio", "dificil", "Introduce un apodo", "iniciar",
		"buena suerte", "tienes", "intentos", "introduce una letra o palabra", "enviado", "letra ya intentada : ", "reproducir",
		"has ganado", "marcador", "reproducir",
		"Atras", "parametro", "Reglas", "clasificacion", "Pistas ?", "Mas", "La palabra era",
	}
	Bd.Set.Language.Ge = []string{"Pac-Hangman Adventure", "leicht", "mittel", "schwer", "einen Namen eingeben", "starten",
		"Viel Gluck", "Sie haben", "Versuche", "Geben Sie einen Buchstaben oder ein Wort ein", "Gesendet", "Buchstabe bereits versucht:", "Wiederholen",
		"Sie haben gewonnen", "Anzeigetafel", "Wiederholung",
		"Zuruck", "Einstellung", "Regeln", "Hohe Punktzahl", "Hinweise?", "Mehr", "Das Wort war",
	}
	Bd.Set.CurrentLanguage = Bd.Set.Language.En

	//initialisation of the game
	var Word = classic.RandomWord("words.txt")
	var data = game{
		Title: "...", Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""), File: "word.txt", ClueNbr: 0,
	}
	Bd.Hangman = data
}

func InitScoreboard() {
	var array []Score
	for i := 1; i <= 9; i++ {
		array = append(array, Score{Top: i, Name: "N/A", Score: 0})
	}
	Bd.Scoreboard.Tab = array
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
	Bd.Scoreboard.Player1 = Bd.Scoreboard.Tab[0]
	Bd.Scoreboard.Player2 = Bd.Scoreboard.Tab[1]
	Bd.Scoreboard.Player3 = Bd.Scoreboard.Tab[2]
	Bd.Scoreboard.Player4 = Bd.Scoreboard.Tab[3]
	Bd.Scoreboard.Player5 = Bd.Scoreboard.Tab[4]
	Bd.Scoreboard.Player6 = Bd.Scoreboard.Tab[5]
	Bd.Scoreboard.Player7 = Bd.Scoreboard.Tab[6]
	Bd.Scoreboard.Player8 = Bd.Scoreboard.Tab[7]
	Bd.Scoreboard.Player9 = Bd.Scoreboard.Tab[8]
}

func InitRankingBoard() {
	var array []Score
	array = append(array, Score{Top: -1, Name: "RV", Score: 200})
	array = append(array, Score{Top: -1, Name: "smith", Score: 900})
	array = append(array, Score{Top: -1, Name: "Chat chat", Score: 300})
	array = append(array, Score{Top: -1, Name: "Luc1ll3", Score: 1000})
	array = append(array, Score{Top: -1, Name: "Leo", Score: 600})
	array = append(array, Score{Top: -1, Name: "REMI", Score: 500})
	array = append(array, Score{Top: -1, Name: "ADAN", Score: 800})

	Bd.Scoreboard.Tab2 = array

	array = append(array, Score{Top: -1, Name: "N/A", Score: 0})
	array = append(array, Score{Top: -1, Name: "N/A", Score: 0})

	SortingSelection(array)
	Bd.Scoreboard.Tab = array
	FixTop(Bd.Scoreboard.Tab)
	HighScore()
}
