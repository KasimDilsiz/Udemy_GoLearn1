package main

import "fmt"

func main() {
	s := "Hello , playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
}

// CIKAN EKRAN GORUNTUSUNDEKI DEGERLERIMIZ BIZIM ASKII TABLOSUNDAKI DEGERLERDIR
// BURADA ASLINDA BIR DIZGENIN BİR DILIM BAYTTAN OLUSTUGUNU GORDUK
//YANI BIR DIZGEYI ALIP BAYT DILIMINE CEVIREBILIRIZ VE SONUCU CIKTI OLARAK ALABILIRIZ .
