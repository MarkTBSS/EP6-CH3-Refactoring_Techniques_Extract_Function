package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person *Person) getPrice() string {
	if person.Age < 12 {
		return "free"
	}
	return "$15"
}

//=====smell===========================================================
func printReceipt(person Person) {
	amount := person.getPrice()
	//receipt details
	fmt.Println("name:", person.Name)
	fmt.Println("total amount:", amount)
}

//=====refactor===========================================================
func printReceiptRefactor(person Person) {
	amount := person.getPrice()
	printDetails(person, amount)
}
func printDetails(person Person, amount string) {
	fmt.Println("name:", person.Name)
	fmt.Println("total amount:", amount)
}
func main() {
	mark := Person{Name: "Mark", Age: 12}
	printReceipt(mark)
	ai_chan := Person{Name: "Ai", Age: 11}
	printReceipt(ai_chan)
	printReceiptRefactor(mark)
	printReceiptRefactor(ai_chan)
}
