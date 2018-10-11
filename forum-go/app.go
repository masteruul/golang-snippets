package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	. "github.com/masteruul/golang-snippets/forum-go/redis"
)

type User struct {
	name   string
	secret string
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "welcome dude \n")
}

func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello "+p.ByName("name"))
}

func HandlerInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "you hit Info Handler with")
}

func BasicAuth(h httprouter.Handle, u User) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == u.name && password == u.secret {
			HandlerInfo(w, r, ps)
		} else {
			w.Header().Set("WWW-Authenticate", "Basic realm = Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func Protected(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header()
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header()
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func Redis(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header()
	w.WriteHeader(200)
	w.Write([]byte("OK"))

	UseRedis()
}

func main() {
	fmt.Println("Hello there ... " + os.Getenv("USER"))

	user := User{name: "uul", secret: "rahasia"}
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/protected/", BasicAuth(Protected, user))
	router.POST("/login", Login)
	router.POST("/register", Register)
	router.GET("/redis", Redis)

	log.Fatal(http.ListenAndServe(":9003", router))
}
