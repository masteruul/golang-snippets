package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	. "github.com/masteruul/golang-snippets/forum-go/common"
	. "github.com/masteruul/golang-snippets/forum-go/conn/posgresql"
	. "github.com/masteruul/golang-snippets/forum-go/handler"
)

func main() {
	fmt.Println("Hello there ... " + os.Getenv("USER"))
	dbinfo := DB{
		Name:"forum",
		User:"postgres"
		Password:"",
		Addr:"127.0.0.1"
	}
	db,err:=InitDB(dbinfo)
	if err != nil {
		log.Println("[ERROR]initiate DB")
	}
	user := User{Name: "uul", Secret: "rahasia"}
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/protected/", BasicAuth(Protected, user))
	router.POST("/login", Login)
	router.POST("/register", Register)
	router.GET("/redis", Redis)
	router.GET("/user", UserIndex)
	router.POST("/user", CreateUser)
	router.GET("/user/:id:", ShowUser)

	log.Fatal(http.ListenAndServe(":9003", router))
}
