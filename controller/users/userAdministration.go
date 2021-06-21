package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/EdwinMejiaLeon/sistema/controller/email"
	"github.com/EdwinMejiaLeon/sistema/db"
	"github.com/gorilla/mux"
)

type ChangeTable struct {
	Id             int
	UsuarioId      int
	SupervisorId   int
	NameSupervisor string
	PreviousChange string
	NewChange      string
	Date           string
}

func MenuAdminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/home.html")
}

func UserAdminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/admin/userAdmin/userIndex.html")
}

func PageEditar(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/admin/userAdmin/userEdit.html")
}

func UserListPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/supervisor/userList.html")
}

func RoleUser(w http.ResponseWriter, r *http.Request) {

	response := UserLogged.InfoUserLogged.Role

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InfoUser(w http.ResponseWriter, r *http.Request) {

	response := UserLogged

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// get users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	var allUser []User
	record := User{}

	query := "SELECT id,name,email,password,role,state FROM users"
	dataBase := db.GetConnection()
	defer dataBase.Close()

	rows, _ := dataBase.Query(query)

	for rows.Next() {

		rows.Scan(
			&record.ID,
			&record.Name,
			&record.Email,
			&record.Password,
			&record.Role,
			&record.State,
		)

		allUser = append(allUser, record)
	}
	//  add info login user
	allUser = append(allUser, UserLogged.InfoUserLogged)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(allUser)
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	query := "DELETE FROM users WHERE id=" + vars["id"] + ""

	dataBase := db.GetConnection()
	defer dataBase.Close()

	dataBase.Exec(query)

	http.ServeFile(w, r, "./templates/admin/userAdmin/userIndex.html")
}

// get user for id
func GetUser(w http.ResponseWriter, r *http.Request) {
	var allUser []User
	record := User{}

	vars := mux.Vars(r)

	query := "SELECT * FROM users WHERE id=" + vars["id"] + ""
	dataBase := db.GetConnection()
	defer dataBase.Close()

	rows, _ := dataBase.Query(query)

	for rows.Next() {

		rows.Scan(
			&record.ID,
			&record.Name,
			&record.Email,
			&record.Password,
			&record.Role,
			&record.State,
		)

		allUser = append(allUser, record)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(allUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	dataBase := db.GetConnection()
	defer dataBase.Close()

	if UserLogged.InfoUserLogged.Role == "supervisor" {
		storeChangeData(vars["id"], r.PostFormValue("Name"), r.PostFormValue("Email"), r.PostFormValue("Role"))
	}

	queryUpdate := `
				UPDATE users SET 
				name = $1, email = $2, role = $3 
				WHERE id = $4
			`

	stmt, err := dataBase.Prepare(queryUpdate)
	if err != nil {
		fmt.Println("error : >>>>>>>>>>", err)
	}
	defer stmt.Close()

	stmt.Exec(r.PostFormValue("Name"), r.PostFormValue("Email"), r.PostFormValue("Role"), vars["id"])

	http.ServeFile(w, r, "./templates/admin/userAdmin/userIndex.html")
}

func ActiveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	query := "UPDATE users SET state = true WHERE id=" + vars["id"] + ""

	dataBase := db.GetConnection()
	defer dataBase.Close()

	dataBase.Exec(query)

	nameUser, emailUser := "Jhon Edwin Mejia Leon", "jemejial@unal.edu.co"
	email.SendEmail(nameUser, emailUser)

	http.ServeFile(w, r, "./templates/admin/userAdmin/userIndex.html")
}

func CountUser(w http.ResponseWriter, r *http.Request) {

	query := "SELECT COUNT(*) FROM users"
	var data, response string = "", ""
	dataBase := db.GetConnection()

	defer dataBase.Close()

	rows, _ := dataBase.Query(query)

	for rows.Next() {

		rows.Scan(
			&data,
		)
	}
	response = "La cantidad de usuarios en el sistema son : " + data

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func storeChangeData(id string, nameNew string, emailNew string, roleNew string) {

	dataBase := db.GetConnection()
	defer dataBase.Close()
	nameOld, emailOld, roleOld, responseOld, responseNew := "", "", "", "", ""

	querySelect := "SELECT name,email,role FROM users WHERE id =" + id
	dataOld, err := dataBase.Query(querySelect)

	if err != nil {
		fmt.Println("error : >>>>>>>>>>", err)
	}

	for dataOld.Next() {

		err = dataOld.Scan(
			&nameOld,
			&emailOld,
			&roleOld,
		)

		if err != nil {
			fmt.Println("error : >>>>>>>>>>", err)
		}
	}

	if nameOld != nameNew {
		responseOld += " Nombre: " + nameOld
		responseNew += " Nombre: " + nameNew
	}

	if emailOld != emailNew {
		responseOld += " Email: " + emailOld
		responseNew += " Email: " + emailNew
	}

	if roleOld != roleNew {
		responseOld += " Rol: " + roleOld
		responseNew += " Rol: " + roleNew
	}

	insertDataChanges(id, responseOld, responseNew)
}

func insertDataChanges(userId string, previousChange string, newChange string) {
	query := `	INSERT INTO
				view_supervisors (usuario_id,supervisor_id,namesupervisor,previouschange,newchange, date)
				VALUES($1,$2,$3,$4,$5,$6)`
	dataBase := db.GetConnection()
	defer dataBase.Close()

	stmt, err := dataBase.Prepare(query)

	if err != nil {
		fmt.Println("error : >>>>>>>>>>", err)
	}
	defer stmt.Close()

	stmt.Exec(userId, UserLogged.InfoUserLogged.ID, UserLogged.InfoUserLogged.Name, previousChange, newChange, time.Now())
}

func SupervisorTableData(w http.ResponseWriter, r *http.Request) {

	var allRecords []ChangeTable
	record := ChangeTable{}

	query := "SELECT id,usuario_id,supervisor_id,namesupervisor,previousChange,newChange,date FROM view_supervisors"
	dataBase := db.GetConnection()
	defer dataBase.Close()

	rows, err := dataBase.Query(query)

	if err != nil {
		fmt.Println("error >>>>>>>>>>", err)
	}

	for rows.Next() {

		err = rows.Scan(
			&record.Id,
			&record.UsuarioId,
			&record.SupervisorId,
			&record.NameSupervisor,
			&record.PreviousChange,
			&record.NewChange,
			&record.Date,
		)

		if err != nil {
			fmt.Println("error >>>>>>>>>>", err)
		}

		allRecords = append(allRecords, record)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(allRecords)

}
