package main

// demınkı kodu yıne ıota kullanarak soyle kısaltabılırız
import "fmt"

const (
	_ = iota
	//kb=1024
	kb = 1 << 10
	mb = 1024 * kb
	gb = 1024 * mb
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}

/* aynı kodu soyle de yazabılırız
const  (
	_=iota
	kb=1<<(iota*10)
	mb=1<<(iota*10)
	gb=1<<(iota*10)
)
*/
