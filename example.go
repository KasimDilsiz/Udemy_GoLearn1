package main

import "fmt"

type conctactInterface interface {
	save()
	update()
	delete()
}
type contact struct {
	personName   string
	personNumber string
}
// Bir şeyler
func imple(k conctactInterface) {
	k.save()
}
func (r contact) delete() {
	fmt.Println(r.personName, "delete")

}
func (r contact) save() {
	fmt.Println(r.personName, "save")
}
func (r contact) update() {
	fmt.Println(r.personName, "uptade")
}

func main() {
	personObject := contact{personName: "ali", personNumber: "123r456"}
	fmt.Println(personObject)
	imple(personObject)

}
