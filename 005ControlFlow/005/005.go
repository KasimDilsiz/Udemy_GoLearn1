package main

import "fmt"

func main() {
	//x := 83 / 40 bunun cıktısı 2 dır bolumu alır
	//y := 83 % 40 bunu cıktısı 3 tur kalanı alır
	// 	fmt.Println(x, y)
	x := 1
	for {
		x++

		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue

		}
		fmt.Println(x)

	}
	fmt.Println("done")
}
