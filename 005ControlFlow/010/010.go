package main

import "fmt"

func main() {

	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("Bond James")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")

	}
}
