package users

import (
	"github.com/EdwinMejiaLeon/sistema/db"
)

func ValidateLogin(email string, password string) bool {
	query := "SELECT email,password,state FROM users WHERE email = '" + email + "' AND password = '" + password + "'"
	dataBase := db.GetConnection()
	defer dataBase.Close()

	rows, err := dataBase.Query(query)
	if err != nil {
		return false
	}

	for rows.Next() {
		u := User{}
		err = rows.Scan(
			&u.Email,
			&u.Password,
			&u.State,
		)

		if err != nil {
			return false
		}

		if (u.Email == email) && (u.Password == password) && u.State {
			return true
		}
	}

	return false
}
