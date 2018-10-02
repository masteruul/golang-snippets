package main

import (
	"fmt"
	"os"
)

//this program is use init function that provide by golang

func init() {
	fmt.Println("Initialitation process...")
	fmt.Println("Done !!!")
	fmt.Println("Hello", os.Getenv("USER"))
}

func main() {
	fmt.Println("Program running....")
}
