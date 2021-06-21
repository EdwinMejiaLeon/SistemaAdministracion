package dataCovid

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EdwinMejiaLeon/sistema/db"
	"github.com/gorilla/mux"
)

type DataCovid struct {
	Question1 string
	Question2 string
	Question3 string
	Question4 string
	Question5 string
}

func InfoCovid(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/basic/covid/index.html")
}

func Survey(w http.ResponseWriter, r *http.Request) {

	var allSurveyCovid []DataCovid
	survey := DataCovid{}

	query := "SELECT question1,question2,question3,question4,question5 FROM survey_covids"
	dataBase := db.GetConnection()
	defer dataBase.Close()

	rows, _ := dataBase.Query(query)

	for rows.Next() {

		err := rows.Scan(
			&survey.Question1,
			&survey.Question2,
			&survey.Question3,
			&survey.Question4,
			&survey.Question5,
		)

		if err != nil {
			fmt.Println("error>>>>>>>>>>>>>>>>>>", err)
		}

		allSurveyCovid = append(allSurveyCovid, survey)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(allSurveyCovid)
}

func SurveyPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/basic/covid/servey.html")
}

func SurveyStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := `	
				INSERT INTO
				survey_covids (user_id,question1,question2,question3,question4,question5)
				VALUES($1,$2,$3,$4,$5,$6)
				`
	dataBase := db.GetConnection()
	defer dataBase.Close()

	stmt, err := dataBase.Prepare(query)

	if err != nil {
		fmt.Println("error>>>>>>>>>>>>>>", err)
	}
	defer stmt.Close()

	stmt.Exec(vars["id"], r.PostFormValue("question1"), r.PostFormValue("question2"), r.PostFormValue("question3"), r.PostFormValue("question4"), r.PostFormValue("question5"))

	http.ServeFile(w, r, "./templates/home.html")
}
