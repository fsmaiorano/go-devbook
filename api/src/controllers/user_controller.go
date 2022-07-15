package controllers

import (
	"api/src/database"
	"api/src/helpers"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"encoding/json"
	"errors"
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

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	err = repository.Delete(ID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	helpers.Json(w, http.StatusNoContent, nil)
}

// Follow a user by id
func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadGateway, err)
		return
	}

	followerID, err := security.ExtractTokenUserId(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID == followerID {
		helpers.Error(w, http.StatusForbidden, errors.New("you can't follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	err = repository.Follow(userID, followerID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusNoContent, nil)
}

// Unfollow a user by id
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadGateway, err)
		return
	}

	followerID, err := security.ExtractTokenUserId(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID == followerID {
		helpers.Error(w, http.StatusForbidden, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	err = repository.Unfollow(userID, followerID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusNoContent, nil)
}

// GetFollowers returns a list of followers
func Followers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadGateway, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	followers, err := repository.GetFollowers(userID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusOK, followers)
}

// Following returns the users that the user is following
func Following(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadGateway, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	following, err := repository.GetFollowing(userID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusOK, following)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadGateway, err)
		return
	}

	userIDToken, err := security.ExtractTokenUserId(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		helpers.Error(w, http.StatusForbidden, errors.New("you can't update the password of another user"))
		return
	}

	var pass models.Password
	err = json.NewDecoder(r.Body).Decode(&pass)
	if err != nil {
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

	storedPassword, err := repository.GetUserPassword(userID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(storedPassword, pass.CurrentPassword); err != nil {
		helpers.Error(w, http.StatusUnauthorized, errors.New("invalid password"))
		return
	}

	err = repository.UpdatePassword(userID, pass.NewPassword)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusNoContent, nil)
}
