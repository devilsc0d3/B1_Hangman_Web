package pages

import (
	"html/template"
	"net/http"
)

func Ranking(w http.ResponseWriter, r *http.Request) {
	start, _ := template.ParseFiles("../Source/Web/" + "ranking" + ".html")
	if r.FormValue("back") == "submit" {
		http.Redirect(w, r, Bd.Set.CurrentPage, http.StatusSeeOther)
	}
	if r.FormValue("more") == "submit" {
		http.Redirect(w, r, "/board", http.StatusSeeOther)
	}
	err := start.ExecuteTemplate(w, "ranking.html", Bd)
	if err != nil {
		return
	}
}

func Scoreboard(w http.ResponseWriter, r *http.Request) {
	start, _ := template.ParseFiles("../Source/Web/" + "scoreboard" + ".html")
	if r.FormValue("back") == "submit" {
		http.Redirect(w, r, "/scoreboard", http.StatusSeeOther)
	}
	err := start.ExecuteTemplate(w, "scoreboard.html", Bd)
	if err != nil {
		return
	}
}
