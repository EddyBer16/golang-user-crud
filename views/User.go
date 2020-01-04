package views

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
}
