package main

import "fmt"

// Customer ...
type Customer struct {
	nama, alamat string
	umur         int
}

func (customer Customer) callfunc(nama string) {
	fmt.Println("callfunc from ", nama, " to ", customer.nama)
}

// Hasname ...
type Hasname interface {
	getname() string
}

func sayhi(hasname Hasname) {
	fmt.Println("Hi ", hasname.getname())
}

//Person ...
type Person struct {
	name string
}

func (person Person) getname() string {
	return person.name
}

func main() {
	data := Customer{
		nama:   "sherdhan",
		alamat: "malang",
		umur:   23,
	}

	data.callfunc("ahmad")

	fmt.Println(data)

	new := Person{
		name: "sherdhan",
	}
	sayhi(new)

}
