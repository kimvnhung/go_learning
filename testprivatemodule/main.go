package main

import (
	"fmt"
	"log"

	submodule "github.com/kimvnhung/gosubmodule"
)

func main() {
	fmt.Println("Hello, World!")
	log.Printf("Using package from private repo")

	p1 := submodule.Point{}
	p1.X = 1
	p1.Y = 2

	log.Printf("P1(x, y) = (%f, %f)", p1.X, p1.Y)
}
