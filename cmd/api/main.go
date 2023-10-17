package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"video/cmd/api/router"
	"video/configs/dbconnection"
	"video/pkg/service"
)

func main() {
	dataBase := dbconnection.DbConnection()
	defer dataBase.Close()
	App := service.Application{DB: dataBase}
	routerMake := router.MakeRouter(App)
	fmt.Println("Starting a server on port: 8080")
	http.ListenAndServe(":8080", routerMake)
}
