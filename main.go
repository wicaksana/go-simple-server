package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/health", health)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func root(writer http.ResponseWriter, req *http.Request) {
	log.Println("\"/\" accessed by " + req.RemoteAddr)
	fmt.Fprint(writer, "default")
}

func health(writer http.ResponseWriter, req *http.Request) {
	log.Println("\"/health\" accessed by " + req.RemoteAddr)
	fmt.Fprint(writer, "OK")
}
