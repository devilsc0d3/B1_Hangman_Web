package pages

import "net/http"
import "html/template"

func Loser(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("../Source/Web/" + "loser" + ".html")

	if r.FormValue("restart") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	if r.FormValue("scoreboard") == "submit" {
		Bd.Set.CurrentPage = "/loser"
		http.Redirect(w, r, "/scoreboard", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "loser.html", Bd)
	if err != nil {
		return
	}
}
