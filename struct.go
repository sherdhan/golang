package main

import (
	"errors"
	"fmt"
)

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

	hasil, err := pembagi(10, 2)
	if err == nil {
		fmt.Println("hasil = ", hasil)
	} else {
		fmt.Println("Error: ", err)
	}

}

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

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 || nilai == 0 {
		return 0, errors.New("Pembagi atau nilai tidak boleh 0")
	}
	hasil := nilai / pembagi
	return hasil, nil
}
