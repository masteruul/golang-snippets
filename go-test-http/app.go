package main

import (
	"log"
	"net/http"

	. "github.com/masteruul/golang-snippets/go-test-http/handler"
)

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/hitung/", HitungHandler)

	message := "server is running now..."
	log.Println(message)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
