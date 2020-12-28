package main

/*1-string veya dilimin yorumlanmamış baytları ----> %s
2-Go sözdizimi ile güvenli bir şekilde kaçan çift tırnaklı bir dize----->%q
3) taban 16, küçük harf, bayt başına iki karakter----->%x
4) taban 16, büyük harf, bayt başına iki karakter----->%X
*/
import "fmt"

func main() {
	s := "Hello , playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\t", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("at the index position %d we have hex %#x\n", i, v)
	}
}
