package pages

import "net/http"
import (
	"html/template"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("../Source/Web/" + "404" + ".html")
	if r.FormValue("home") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "404.html", Bd)
	if err != nil {
		return
	}
}
