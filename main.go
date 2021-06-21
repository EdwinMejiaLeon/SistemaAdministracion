package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EdwinMejiaLeon/sistema/controller"
	"github.com/EdwinMejiaLeon/sistema/controller/dataCovid"
	"github.com/EdwinMejiaLeon/sistema/controller/users"
	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter().StrictSlash(true)

	route.HandleFunc("/", controller.IndexHandler)
	route.HandleFunc("/login", controller.LoginHandler).Methods("GET")
	route.HandleFunc("/login", controller.NowUserHandler).Methods("POST")
	route.HandleFunc("/home", controller.GetInLoginHandler).Methods("POST")
	route.HandleFunc("/register", controller.RegisterHandler).Methods("GET")
	route.HandleFunc("/index", controller.HomeHandler)

	// admin
	route.HandleFunc("/menu", users.MenuAdminHandler).Methods("GET")
	route.HandleFunc("/UserRole", users.RoleUser).Methods("GET")
	route.HandleFunc("/InfoUser", users.InfoUser).Methods("GET")
	route.HandleFunc("/SupervisorData", users.SupervisorTableData).Methods("GET")
	// admin user
	route.HandleFunc("/UserAdmin", users.UserAdminHandler).Methods("GET") //index of admin users
	route.HandleFunc("/UserAll", users.GetUsers).Methods("GET")           //json of all users
	route.HandleFunc("/UserDelete/{id}", users.DeleteUser).Methods("POST")
	route.HandleFunc("/UserEditPage/{id}", users.PageEditar).Methods("GET")
	route.HandleFunc("/UserGetEdit/{id}", users.GetUser).Methods("GET")   //json of a user
	route.HandleFunc("/UserAll/{id}", users.UpdateUser).Methods("POST")   //update a user
	route.HandleFunc("/UserActive/{id}", users.ActiveUser).Methods("GET") //update a user
	// admin count
	route.HandleFunc("/UserCount", users.CountUser).Methods("GET")

	// supervisor
	route.HandleFunc("/UserList", users.UserListPage).Methods("GET")

	// user
	route.HandleFunc("/User/{id}", users.PageEditar).Methods("GET")
	// covid
	route.HandleFunc("/User/covidInformationPage/{id}", dataCovid.InfoCovid).Methods("GET")
	route.HandleFunc("/User/covidSurveyPage/{id}", dataCovid.SurveyPage).Methods("GET")
	route.HandleFunc("/User/covidInformationPage/{id}", dataCovid.SurveyStore).Methods("POST")
	route.HandleFunc("/covidInformationPage/{id}", dataCovid.Survey).Methods("GET") //API GET

	fmt.Printf("Servidor corriendo en http://localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", route))
}
