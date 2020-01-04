package controller

import (
	"net/http"
	"strings"

	"../../db"
)

func HandDeleteUser(w http.ResponseWriter, r *http.Request) {
	// CHECK THE ID PROVIDED
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id != "" {
		// HAND ERROR
		if err := db.DeleteUser(id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("There's some error"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("User deleted successfully"))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID was not provided"))
	}
}
