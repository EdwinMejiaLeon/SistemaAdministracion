package controller

import (
	"net/http"

	"github.com/EdwinMejiaLeon/sistema/controller/users"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/login.html")
}

func GetInLoginHandler(w http.ResponseWriter, r *http.Request) {

	Email, Password := r.PostFormValue("Email"), r.PostFormValue("Password")

	if users.ValidateLogin(Email, Password) {
		users.GetDataUser(Email, Password)
		http.ServeFile(w, r, "./templates/home.html")
	}

	http.ServeFile(w, r, "./templates/login.html")
}
