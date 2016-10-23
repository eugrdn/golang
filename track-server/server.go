package main

import (
	"log"
	"net/http"
	handlers "github.com/golang/track-server/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/favicon.ico", handlers.HandlerIcon) 
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal("Fatal error", err)
	}

}
