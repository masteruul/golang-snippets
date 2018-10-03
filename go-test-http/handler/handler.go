package handler

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	name = strings.TrimPrefix(name, "/")
	output := "Hello " + name + " you are online in " + os.Getenv("USER") + " device"
	w.Write([]byte(output))
}

func HitungHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Path
	input = strings.TrimPrefix(input, "/hitung/")
	x, err := strconv.Atoi(input)
	if err != nil {
		w.Write([]byte("you typing a string not a number"))
		return
	}
	sum := x + rand.Intn(100)
	output := "result of your number now is " + strconv.Itoa(sum)

	w.Write([]byte(output))
}
