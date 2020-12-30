package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\t %#U\t %#x\n", i, i, i)
	}

}
