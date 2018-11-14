package main

import (
	"log"
	"net/http"

	"github.com/leiysky/cloudgo-io/handlers"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}
	http.Handle("/", handlers.StaticFileHandler("assets"))
	http.HandleFunc("/post", handlers.FormHandler)
	http.HandleFunc("/unknown", handlers.UnknownHandler)
	log.Fatal(server.ListenAndServe())
}
