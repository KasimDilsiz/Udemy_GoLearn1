package main

// haritadan bir şeyi silmek için
// delete(<map name>,"key") ifadesini kullanırız .
import "fmt"

func main() {

	m := map[string]int{
		"James": 32,
		"Miss":  27,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)

	fmt.Println(m["Miss"])
	fmt.Println(m["Ian Fleming"])

	// eğer bir şeyi kesin olarak sildiğimizi görmek istiyorsak , virgük, ok kullanabiliriz .
	if v, ok := m["Miss"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss")

	}
	fmt.Println(m)
}
