package pages

import "net/http"
import "html/template"

func Setting(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("../Source/Web/" + "setting" + ".html")
	if r.FormValue("language") == "en" {
		Bd.Set.CurrentLanguage = Bd.Set.Language.En
	}
	if r.FormValue("language") == "fr" {
		Bd.Set.CurrentLanguage = Bd.Set.Language.Fr
	}
	if r.FormValue("language") == "es" {
		Bd.Set.CurrentLanguage = Bd.Set.Language.Es
	}
	if r.FormValue("language") == "ge" {
		Bd.Set.CurrentLanguage = Bd.Set.Language.Ge
	}
	if r.FormValue("apply") == "submit" || r.FormValue("back") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	err := page.ExecuteTemplate(w, "setting.html", Bd)
	if err != nil {
		return
	}
}
