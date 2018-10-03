package main

import (
	"log"
	"net/http"

	. "github.com/masteruul/golang-snippets/go-test-http/handler"
)

func sameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	out := "you hit " + name + " API"
	w.Write([]byte(out))
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/same", sameHandler)
	http.HandleFunc("/hitung", HitungHandler)

	message := "server is running now..."
	log.Println(message)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
