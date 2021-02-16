package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/users/sahil-9/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}