package controllers

import (
	"api/src/database"
	"api/src/helpers"
	"api/src/models"
	"api/src/repositories"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
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

	if err := user.Prepare(); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

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

	line, err := db.Query("SELECT * FROM users WHERE id = @id;", sql.Named("id", ID))
	if err != nil {
		w.Write([]byte("Error preparing statement: " + err.Error()))
		return
	}

	var user models.User
	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scanning line: " + err.Error()))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error encoding user: " + err.Error()))
		return
	}
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
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(request, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database: " + err.Error()))
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	err = repository.Update(user)
	if err != nil {
		w.Write([]byte("Error updating user: " + err.Error()))
		return
	}

	defer db.Close()
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

	repository := repositories.NewRepositoryOfUsers(db)
	err = repository.Delete(ID)
	if err != nil {
		w.Write([]byte("Error deleting user: " + err.Error()))
		return
	}

	defer db.Close()
}
