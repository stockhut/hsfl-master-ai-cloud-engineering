package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Hello from the static file server")

	router := http.NewServeMux()
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", router))
}
