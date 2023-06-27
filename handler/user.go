package handler

import (
	"aytodo/database"
	"aytodo/database/dbHelper"
	"aytodo/utils"
	"net/http"
	"time"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	userID, userErr := dbHelper.GetUserIDByPassword(body.Email, body.Password)
	if userErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, userErr, "failed to find user")
		return
	}
	// create user session
	sessionToken := utils.HashString(body.Email + time.Now().String())
	sessionErr := dbHelper.CreateUserSession(database.Todo, userID, sessionToken)
	if sessionErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, sessionErr, "failed to create user session")
		return
	}
	utils.RespondJSON(w, http.StatusCreated, struct {
		Token string `json:"token"`
	}{
		Token: sessionToken,
	})
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if parseError := utils.ParseBody(r.Body, &body); parseError != nil {
		utils.RespondError(w, http.StatusInternalServerError, parseError, "failed to check user existence")
		return
	}
	exists, existsErr := dbHelper.IsUserExists(body.Email)
	if existsErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, existsErr, "failed to check user existence")
		return
	}
	if exists {
		utils.RespondError(w, http.StatusBadRequest, nil, "user already exists")
		return
	}
	hashedPassword, hasErr := utils.HashPassword(body.Password)
	if hasErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, hasErr, "failed to secure password")
		return
	}

	sessionToken := utils.HashString(body.Email + time.Now().String())
	userID, saveErr := dbHelper.CreateUser(database.Todo, body.Name, body.Email, hashedPassword)
	if saveErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, saveErr, "failed to save user")
	}
	sessionErr := dbHelper.CreateUserSession(database.Todo, userID, sessionToken)
	if sessionErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, sessionErr, "failed to create user session")
	}

	utils.RespondJSON(w, http.StatusCreated, struct {
		Token string `json:"token"`
	}{
		Token: sessionToken,
	})
}

// func LogoutUser(w http.ResponseWriter, r *http.Request) {
// 	token := r.Header.Get("x-api-key")
// 	userCtx := middleware.UserContext(r)
// 	err := dbHelper.DeleteSessionToken(userCtx.ID, token)
// 	if err != nil {
// 		utils.RespondError(w, http.StatusInternalServerError, err, "failed to logout user")
// 		return
// 	}
// 	w.WriteHeader(http.StatusAccepted)
// }
