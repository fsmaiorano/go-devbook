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

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := security.ExtractTokenUserId(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(request, &post); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	if err = post.Prepare(); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfPosts(db)
	post.ID, err = repository.Create(post, userID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusCreated, post)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewRepositoryOfPosts(db)
	post, erro := repository.FindByID(postID)
	if erro != nil {
		helpers.Error(w, http.StatusInternalServerError, erro)
		return
	}

	helpers.Json(w, http.StatusOK, post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := security.ExtractTokenUserId(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfPosts(db)
	posts, err := repository.FindAll(userID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusOK, posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, err := security.ExtractTokenUserId(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfPosts(db)

	storedPost, err := repository.FindByID(postID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if storedPost.AuthorID != userID {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(request, &post); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	if err = post.Prepare(); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	err = repository.Update(storedPost.ID, post)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusOK, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, err := security.ExtractTokenUserId(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryOfPosts(db)

	storedPost, err := repository.FindByID(postID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if storedPost.AuthorID != userID {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	err = repository.Delete(storedPost.ID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusNoContent, nil)
}

func FindPostsByUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewRepositoryOfPosts(db)
	posts, err := repository.FindByUserID(userID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.Json(w, http.StatusOK, posts)
}
