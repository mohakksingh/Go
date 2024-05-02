package main

import (
	"log"
	"fmt"
	"go-postgres/router"
	"net/http"
)

func main(){
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe("8080",r))
}
