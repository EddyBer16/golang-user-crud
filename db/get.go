package db

import (
	"../views"
)

func GetAll() ([]views.User, error) {
	rows, err := con.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	users := []views.User{}
	for rows.Next() {
		data := views.User{}
		rows.Scan(&data.ID, &data.Username, &data.Password)
		users = append(users, data)
	}
	return users, nil
}

func GetByID(id string) (views.User, error) {
	rows, err := con.Query("SELECT * FROM users WHERE ID=?", id)
	if err != nil {
		return views.User{}, err
	}
	user := views.User{}
	for rows.Next() {
		data := views.User{}
		rows.Scan(&data.ID, &data.Username, &data.Password)
		user = data
	}
	if user.ID == 0 && user.Username == "" && user.Password == "" {
		return views.User{}, nil
	} else {
		return user, nil
	}
}
