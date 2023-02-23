package main

import (
	"classic"
	"fmt"
	"net/http"
	"os"
	"web/go_file/pages"
)

const port = ":8080"

func main() {
	if len(os.Args) == 2 {
		classic.Hangman()
	} else {
		Server()
	}
}

func Server() {
	pages.Variable()
	pages.InitRankingBoard()

	http.HandleFunc("/home", pages.Home)
	http.HandleFunc("/404", pages.NotFound)
	http.HandleFunc("/rules", pages.Rules)
	http.HandleFunc("/setting", pages.Setting)
	http.HandleFunc("/loser", pages.Loser)
	http.HandleFunc("/win", pages.Win)
	http.HandleFunc("/scoreboard", pages.Ranking)
	http.HandleFunc("/board", pages.Scoreboard)
	http.HandleFunc("/", pages.Hangman)

	fmt.Println("http://localhost" + port + "/home")

	fs := http.FileServer(http.Dir("../Source"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
