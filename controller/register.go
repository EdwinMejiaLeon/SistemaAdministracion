package controller

import (
	"log"
	"net/http"

	"github.com/EdwinMejiaLeon/sistema/controller/users"
)

//page register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/register.html")
}

func NowUserHandler(w http.ResponseWriter, r *http.Request) {

	accessPermission := true

	if r.PostFormValue("Role") == "admin" {
		accessPermission = false
	}

	// match, _ := regexp.MatchString(`/(?i)(\W|^)[\w.\-]{0,25}@(gmail|hotmail|unal)\.(p|com|co)(\W|$)/`, r.PostFormValue("Email"))
	// if !match {
	// 	log.Fatal("error ->>>>>>>>>>>> La estructura del correo no es permitida")
	// }

	nowUser := users.User{
		Name:     r.PostFormValue("Name"),
		Email:    r.PostFormValue("Email"),
		Password: r.PostFormValue("Password"),
		Role:     r.PostFormValue("Role"),
		State:    accessPermission,
	}

	err := users.UserCreate(nowUser)

	if err != nil {
		log.Fatal(err)
	}

	http.ServeFile(w, r, "./templates/login.html")
}
