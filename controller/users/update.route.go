package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"../../db"
	"../../views"
)

func HandUpdateUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	data := views.User{}
	json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data)
	if err := db.UpdateUser(id, data); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There's some error"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User update successfully"))
	}
}
