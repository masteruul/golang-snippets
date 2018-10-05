package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println(w, "welcome dude \n")
}

func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(w, "hello "+p.ByName("name"))
}

func main() {
	fmt.Println("Hello there ... " + os.Getenv("USER"))

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":9003", router))
}
