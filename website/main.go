package main

import (
	"log"
	"net/http"

	"github.com/golang/website/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	userC := controllers.UserC

	router := httprouter.New()
	router.GET("/", userC.Redirector)
	router.GET("/users", userC.UserListRender)

	log.Fatal(http.ListenAndServe(":8085", router))
}