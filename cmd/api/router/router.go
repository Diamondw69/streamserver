package router

import (
	"github.com/gorilla/mux"
	"video/pkg/service"
)

func MakeRouter(application service.Application) *mux.Router {
	router := mux.NewRouter()
	//other urls
	router.HandleFunc("/insert", application.InsertHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/stream/{id}", application.StreamPageHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/", application.MainPageHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin", application.AdminPageHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/{id}", application.SelectHandler).Methods("GET", "OPTIONS")
	return router
}
