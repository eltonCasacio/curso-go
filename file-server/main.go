package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./file-server/public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
