package pages

import (
	"html/template"
	"net/http"
)

func Win(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("../Source/Web/" + "win" + ".html")
	if r.FormValue("restart") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	if r.FormValue("scoreboard") == "submit" {
		Bd.Set.CurrentPage = "/win"
		http.Redirect(w, r, "/scoreboard", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "win.html", Bd)
	if err != nil {
		return
	}
}
