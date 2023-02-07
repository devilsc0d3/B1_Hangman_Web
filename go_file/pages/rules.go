package pages

import (
	"html/template"
	"net/http"
)

func Rules(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("../Source/Web/" + "rules" + ".html")
	if r.FormValue("back") == "submit" {
		http.Redirect(w, r, Bd.Set.CurrentPage, http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "rules.html", Bd)
	if err != nil {
		return
	}
}
