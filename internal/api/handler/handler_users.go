package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naijasonezeiru/go-phish-backend/internal/api/helper"
	"github.com/naijasonezeiru/go-phish-backend/internal/database"
)

type successResponse struct {
	Message    string `json:"message"`
	Token      string `json:"token"`
	StatusCode int    `json:"statusCode"`
}

// HandlerRegister godoc
//
//	@Summary		resgister user
//	@Description	Register a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			username	formData	string			true	"Your Username"	minlength(4)	maxlength(30)	example(johndoe)
//	@Param			password	formData	string			true	"Your Password"	minlength(4)	maxlength(30)	example(password123)
//	@Success		200			{object}	successResponse	"ok"
//	@Failure		400			{object}	helper.ErrResponse
//	@Failure		404			{object}	helper.ErrResponse
//	@Failure		500			{object}	helper.ErrResponse
//	@Router			/users [post]
func HandlerRegister(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	hashedPassword, err := helper.HashPassword(params.Password)
	if err != nil {
		helper.RespondWithError(w, 500, fmt.Sprintf("Error hashing password: %s", err))
		return
	}
	user, err := helper.ConnectDB().DB.CreateUser(r.Context(), database.CreateUserParams{
		Username:     params.Username,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Couldn't create User %s", err))
		return
	}
	token, err := helper.CreateJWTToken(user.Username)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Couldn't create JWT token %s", err))
	}
	helper.RespondWithJSON(w, 201, successResponse{Message: "Successfully register", Token: token, StatusCode: 201})
}

// LoginHandler godoc
//
//	@Summary		user login
//	@Description	login
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			username	formData	string		true	"Your Username"	minlength(4)	maxlength(30)	example(johndoe)
//	@Param			password	formData	string		true	"Your Password"	minlength(4)	maxlength(30)	example(password123)
//	@Success		200			{object}	helper.User	"ok"
//	@Failure		400			{object}	helper.ErrResponse
//	@Failure		404			{object}	helper.ErrResponse
//	@Failure		500			{object}	helper.ErrResponse
//	@Router			/auth [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := helper.ConnectDB().DB.GetUserByUsername(r.Context(), params.Username)
	if err != nil {
		helper.RespondWithError(w, 400, "Username name or password is invalid")
		return
	}
	passwordMatch := helper.CheckPassword(params.Password, user.PasswordHash)
	if !passwordMatch {
		helper.RespondWithError(w, 403, "Username name or password is invalid")
		return
	}
	token, err := helper.CreateJWTToken(user.Username)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Couldn't create JWT token %s", err))
	}

	helper.RespondWithJSON(w, 201, helper.DatabaseUserToUser(user, token))
	// helper.RespondWithJSON(w, 201, user)
}

// HandlerGetMe godoc
//
//	@securityDefinitions.basic	BasicAuth
//	@name						Authorization
//	@in							header
//	@description				OAuth protects our entity endpoints
//	@Summary					persist login
//	@Description				Verifies user and returns the token bearer
//	@Tags						users
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	helper.User	"ok"
//	@Failure					400	{object}	helper.ErrResponse
//	@Failure					404	{object}	helper.ErrResponse
//	@Failure					500	{object}	helper.ErrResponse
//	@Router						/users/me [post]
func HandlerGetMe(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)

	user, err := helper.ConnectDB().DB.GetUserByUsername(r.Context(), userId)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error getting user: %s", err))
		return
	}
	token, err := helper.CreateJWTToken(user.Username)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Couldn't create JWT token %s", err))
	}
	helper.RespondWithJSON(w, 200, helper.DatabaseUserToUser(user, token))
	// helper.RespondWithJSON(w, 200, user)
}
