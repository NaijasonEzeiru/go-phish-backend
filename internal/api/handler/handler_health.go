package handler

import (
	"net/http"

	"github.com/naijasonezeiru/go-phish-backend/internal/api/helper"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, 200, "Server up and running")
}

func HandleErr(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusInternalServerError, "something went wrong")
}
