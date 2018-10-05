package handler

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		name := r.URL.Path
		name = strings.TrimPrefix(name, "/")
		output := "Hello " + name + " you are online in " + os.Getenv("USER") + " device"
		w.Write([]byte(output))
	} else {
		message := "Method not allowed"
		w.Write([]byte(message))
	}
}

func HitungHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		input := r.URL.Path
		input = strings.TrimPrefix(input, "/hitung/")
		x, err := strconv.Atoi(input)
		if err != nil {
			w.Write([]byte("you typing a string not a number"))
		}
		rn := rand.Intn(100)
		sum := x + rn
		output := "magic number generate " + strconv.Itoa(rn) + " -> result of your number now is " + strconv.Itoa(sum)

		w.Write([]byte(output))
	} else {
		msg := "Method not allowed"
		w.Write([]byte(msg))

	}
}
