package main

import (
	"fmt"
	"log"
	"net/http"

	"./controller/users"
	"./db"
)

func main() {
	routes := controller.Router()
	db := db.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":3000", routes))
	fmt.Print("server in port 3000")
}
