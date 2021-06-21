package users

import (
	"github.com/EdwinMejiaLeon/sistema/db"
)

type LoggedInUser struct {
	InfoUserLogged User
	LoggedUser     bool
}

var UserLogged LoggedInUser

func GetDataUser(email string, password string) error {
	query := "SELECT id,name,email,password,role,state FROM users WHERE email = '" + email + "' AND password = '" + password + "'"
	dataBase := db.GetConnection()
	defer dataBase.Close()

	rows, err := dataBase.Query(query)
	if err != nil {
		return err
	}

	for rows.Next() {

		err = rows.Scan(
			&UserLogged.InfoUserLogged.ID,
			&UserLogged.InfoUserLogged.Name,
			&UserLogged.InfoUserLogged.Email,
			&UserLogged.InfoUserLogged.Password,
			&UserLogged.InfoUserLogged.Role,
			&UserLogged.InfoUserLogged.State,
		)

		if err != nil {
			return err
		}

		UserLogged.LoggedUser = true
	}
	return nil
}
