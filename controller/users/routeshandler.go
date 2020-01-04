package controller

import (
	"net/http"
)

func routesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet { // GET METHOD
			HandGetUser(w, r)

		} else if r.Method == http.MethodPost { // POST METHOD
			HandPostUser(w, r)

		} else if r.Method == http.MethodDelete { // DELETE METHOD
			HandDeleteUser(w, r)

		} else if r.Method == http.MethodPut { // PUT METHOD
			HandUpdateUser(w, r)
		}
	}
}
