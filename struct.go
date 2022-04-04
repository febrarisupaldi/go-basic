package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Married       bool
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huuu from", a.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	eko.sayHi("Joko")
	eko.sayHuu()

	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 35, true}
	fmt.Println(budi)

}
