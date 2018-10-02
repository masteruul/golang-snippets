package main

import (
	"fmt"
	"math"
	"os"
)

type geometry interface {
	luas() float64
	keliling() float64
}

type segi4 struct {
	panjang, lebar float64
}

type lingkaran struct {
	jari2 float64
}

func (s segi4) luas() float64 {
	return s.panjang * s.lebar
}

func (s segi4) keliling() float64 {
	return 2 * (s.panjang + s.lebar)
}

func (l lingkaran) luas() float64 {
	return math.Pi * l.jari2 * l.jari2
}

func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.jari2
}

func hitung(g geometry) {
	fmt.Println(g)
	fmt.Println(g.luas())
	fmt.Println(g.keliling())
}

//for initiate program.
func Init() {
	fmt.Println("Hello bro ", os.Getenv("USER"))
}
func main() {
	Init()
	fmt.Println("in this section, i want to implement interface in golang")
	//initiate 2 struct data
	s := segi4{lebar: 3, panjang: 4}
	l := lingkaran{jari2: 5}
	hitung(s)
	hitung(l)
}
