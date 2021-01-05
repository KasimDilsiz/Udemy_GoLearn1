/*  package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Sinan",
		last:  "Tayak",
		age :22,
	}
	fmt.Println(p1)
}
yukarıda yazmış olduğumuz kod blogunun structun isminin olduğuna dikkat edelim
aşağıda onu anonim bir structt haline getirdik .  */

package main

import "fmt"

func main() {
	p1 :=
		struct {
			first string
			last  string
			age   int
		}{
			first: "Sinan",
			last:  "Tayak",
			age:   22,
		}

	fmt.Println(p1)
}
