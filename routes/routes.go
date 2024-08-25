package routes

import (
	"github.com/hewo233/cxsj_homework1/controller"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/query", controller.QueryHandler)
	http.HandleFunc("/list", controller.ListHandler)
	http.HandleFunc("/ping", controller.PingHandler)
	http.HandleFunc("/add", controller.AddHandler)
	http.HandleFunc("/delete", controller.DeleteHandler)
	http.HandleFunc("/modify", controller.ModifyHandler)
}
