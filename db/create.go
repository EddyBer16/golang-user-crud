package db

import (
	"fmt"
)

func CreateUser(username, password string) error {
	data, err := con.Query("INSERT INTO users VALUES(null, ?, ?)", username, password)
	defer data.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
