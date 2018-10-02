package main

import (
	"fmt"
	"os"
)

func Init() {
	fmt.Println("Hello ", os.Getenv("USER"))
}

func main() {
	Init()

}
