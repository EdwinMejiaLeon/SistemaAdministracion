package users

import (
	"errors"
	"time"

	"github.com/EdwinMejiaLeon/sistema/db"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Role      string
	State     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// create user
func UserCreate(u User) error {
	query := `	INSERT INTO
				users (name,email,password,role,state)
				VALUES($1,$2,$3,$4,$5)`
	dataBase := db.GetConnection()
	defer dataBase.Close()

	stmt, err := dataBase.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(u.Name, u.Email, u.Password, u.Role, u.State)

	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("se esperaba 1 fila afectada")
	}

	return nil
}
