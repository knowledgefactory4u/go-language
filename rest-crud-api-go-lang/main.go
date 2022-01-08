package main

import (
	"log"
	"net/http"
	"rest-crud-api-go-lang/config"
	routers "rest-crud-api-go-lang/router"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()
	log.Println("Starting the HTTP server on port 9080")
	router := mux.NewRouter().StrictSlash(true)
	routers.InitaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":9080", router))
}
