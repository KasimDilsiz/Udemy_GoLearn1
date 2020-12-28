package main

import "fmt"

const (
	C0 = iota
	C1 = iota
	C2 = iota
)
const (
	C4 = iota
	C5 = iota
	C6 = iota
)

func main() {

	fmt.Println(C0, C1, C2) // "0 1 2"
	fmt.Println(C4, C5, C6) // "0 1 2" burada belırtmek ıstedıgımız bunu tekrar yazdıgımıda sıfırlanır
}

/*
const (
	C0 = iota
	C1
	C2
)
yukarıdakı ıfade yerıne bu kısa ıfadeyı de kullanabılırız.
Iota anahtar sozcugu ardasık tam sayı sabıtleri temsıl eder ; 0,1,2 ..
*/
