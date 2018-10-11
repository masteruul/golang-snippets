package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	. "github.com/masteruul/golang-snippets/forum-go/common"
	. "github.com/masteruul/golang-snippets/forum-go/conn/redis"
)

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

		if hasAuth && user == u.Name && password == u.Secret {
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

func UserIndex(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//res, err := db.Query("SELECT * FROM user")
	err := "state on development"
	if err != "" {
		log.Println("[ERROR] cant fecth data from database")
	}
	w.Header()
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header()
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func ShowUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header()
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
