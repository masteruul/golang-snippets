package main

import (
	"fmt"
	"os"
)

//this program is use init function that provide by golang

func Init() {
	fmt.Println("Initialitation process...")
	fmt.Println("Done !!!")
	fmt.Println("Hello", os.Getenv("USER"))
}

func main() {
	Init()
	fmt.Println("Program running....")
}
