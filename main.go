package main

import (
	"fmt"
	"github.com/hewo233/cxsj_homework1/common"
	"github.com/hewo233/cxsj_homework1/routes"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Print("Enter Password: ")
	passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading password")
		return
	}
	dbPassword := string(passwordBytes)

	common.UserDB, _ = common.InitDB(dbPassword)
	defer common.CloseDB(common.UserDB)

	routes.SetupRoutes()

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
