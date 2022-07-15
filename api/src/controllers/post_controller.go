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

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreatePost"))
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreatePost"))
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreatePost"))
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreatePost"))
}
