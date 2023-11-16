package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"video/cmd/api/router"
	"video/configs/dbconnection"
	"video/pkg/service"
)

func main() {
	dataBase := dbconnection.DbConnection()
	defer func(dataBase *sql.DB) {
		err := dataBase.Close()
		if err != nil {

		}
	}(dataBase)
	App := service.Application{DB: dataBase}
	routerMake := router.MakeRouter(App)
	fmt.Println("Starting a server on port: 8080")
	err := http.ListenAndServe(":8080", routerMake)
	if err != nil {
		return
	}
}
