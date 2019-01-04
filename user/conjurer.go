package user

import (
	"github.com/ryu/epic_affinity/config"
	"net/http"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	xu, err := AllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
	config.TPL.ExecuteTemplate(w, "user-index.html", xu)
}

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "POST" {
		_, err := CreateUser(r)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}

		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}
	config.TPL.ExecuteTemplate(w, "create-user.html", nil)
}


func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}
