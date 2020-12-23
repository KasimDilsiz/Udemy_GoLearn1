package main

import "fmt"

type circle struct {
	radius int
}

func (a *circle) area() int {
	return a.radius * a.radius * 3

}
func (a *circle) perimeter() int {
	return 2 * 3 * a.radius

}

func main() {
	b := circle{radius: 4}
	fmt.Println(b.perimeter(), "cevre")
	fmt.Println(b.area(), "alan")

}
