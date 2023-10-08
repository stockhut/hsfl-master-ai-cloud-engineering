package main

import (
	"fmt"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
)

func main() {
	fmt.Println("Hello from Auth!")

	r := router.New()

	r.GET("/", handleHi)

	err := http.ListenAndServe("localhost:8080", r)
	panic(err)
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi!")
}
