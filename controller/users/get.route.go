package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"../../db"
)

func HandGetUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	// CHECK IF IN THE URL WAS PROVIDED AN ID
	if id != "" {
		data, err := db.GetByID(id)
		// HAND ERROR AND WRITE TO SERVER
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		// CHECK IF THE USER EXISTS
		if data.ID == 0 && data.Username == "" && data.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No data"))
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	} else {
		data, err := db.GetAll()
		// HAND ERROR AND WRITE TO SERVER
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}
