package main

import (
	"github.com/hewo233/cxsj_homework1/common"
	"github.com/hewo233/cxsj_homework1/routes"
	"log"
	"net/http"
	"os"
)

func main() {

	/*
		fmt.Print("Enter Password: ")
		passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password")
			return
		}
		dbPassword := string(passwordBytes)
	*/

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("DB_PASSWORD not set")
		return
	}

	common.UserDB, _ = common.InitDB(dbPassword)
	defer common.CloseDB(common.UserDB)

	routes.SetupRoutes()

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
