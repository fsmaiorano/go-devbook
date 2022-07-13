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
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(request, &user); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare("signup"); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusCreated, user)
}

// GetUser gets a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	user, err := repository.FindById(ID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusOK, user)
}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database: " + err.Error()))
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	users, err := repository.FindByNameOrNick(nameOrNick)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	if len(users) == 0 {
		helpers.Json(w, http.StatusNotFound, users)
		return
	}

	helpers.Json(w, http.StatusOK, users)
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadGateway, err)
		return
	}

	userIDToken, err := security.ExtractTokenUserId(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if ID != userIDToken {
		helpers.Error(w, http.StatusForbidden, err)
		return
	}

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(request, &user); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	err = repository.Update(ID, user)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusNoContent, nil)
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error parsing id: " + err.Error()))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database: " + err.Error()))
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	err = repository.Delete(ID)
	if err != nil {
		w.Write([]byte("Error deleting user: " + err.Error()))
		return
	}

	helpers.Json(w, http.StatusNoContent, nil)
}
