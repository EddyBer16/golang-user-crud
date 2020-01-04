package controller

import (
	"encoding/json"
	"net/http"

	"../../db"
	"../../views"
)

func HandPostUser(w http.ResponseWriter, r *http.Request) {
	user := views.User{}
	json.NewDecoder(r.Body).Decode(&user)
	// HAND ERROR
	if err := db.CreateUser(user.Username, user.Password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There's some error"))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created successfully"))
	}
}
