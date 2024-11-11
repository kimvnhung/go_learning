package main

import (
	"fmt"
	"log"

	submodule "github.com/kimvnhung/gosubmodule"
)

func main() {
	fmt.Println("Hello, World!")
	log.Printf("Using package from private repo")

	p1 := &submodule.Point{X: 1, Y: 2}

	p2 := &submodule.Point{X: 3, Y: 4}

	p3 := p1.Middle(p2)

	log.Printf("P1(x, y) = (%f, %f)", p3.X, p3.Y)
}
