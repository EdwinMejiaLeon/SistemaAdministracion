package controller

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/welcome.html")
}
