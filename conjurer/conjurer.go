package conjurer

import (
	"github.com/ryu/epic_affinity/user"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func Conjure() {
	http.HandleFunc("/", handleIndex)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public/"))))
	http.HandleFunc("/users", user.Index)
	http.HandleFunc("/user/create", user.Create)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
