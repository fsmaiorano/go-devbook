package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// CreateUser creates a new user
func TestCreateUser(t *testing.T) {
	t.Run("Create user", func(t *testing.T) {
		t.Log("Starting UnitTest Create user")

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/users", nil)
		CreateUser(w, r)
		if w.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	})
}
