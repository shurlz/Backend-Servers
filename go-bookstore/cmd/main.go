package main

import (
	"log"
	"net/http"
	"github.com/gorrila/mux"
	_ "github.com/jinzhu/gorm/dialetics/mysql"
	"github.com/shurlz/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
