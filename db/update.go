package db

import (
	"fmt"

	"../views"
)

func UpdateUser(id string, user views.User) error {
	fmt.Println(user.Username, user.Password)
	_, err := con.Query("UPDATE users SET username=?, password=? WHERE ID=?", user.Username, user.Password, id)
	if err != nil {
		return err
	}
	return nil
}
