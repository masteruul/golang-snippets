package main

import "fmt"

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println("value ", i, " from ", from)
	}
}

func main() {
	fmt.Println("this is for go routine example")
	f("direct")

	go f("goroutine")
	fmt.Println("done")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
