package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sriram-bb63/go-bookstore/pkg/routes"
	_ "gorm.io/driver/sqlite"
)

func main() {
	fmt.Println("Started")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
