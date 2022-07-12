package controllers

import (
	"api/src/database"
	"api/src/helpers"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login is a function that handles the login request
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	storedUser, err := repository.AuthenticationFindByEmail(user.Email)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(storedUser.Password, user.Password); err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte("Login successful"))
}
