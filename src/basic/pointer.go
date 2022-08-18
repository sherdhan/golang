package main

import "fmt"

type alamat struct {
	negara, provinsi, kota string
}

type Man struct {
	name string
}

func change_alamat(alamat *alamat) {
	alamat.negara = "jepang"
}

//beri pointer * agar variable berubah
func (man *Man) married() {
	man.name = "Mr. " + man.name
}

func main() {
	alamat1 := alamat{"indonesia", "jawa timur", "malang"}
	alamat2 := &alamat1
	// pointer baru kosong
	alamat3 := new(alamat)

	// alamat2.kota = "surabaya"
	*alamat2 = alamat{"indonesia", "jawa timur", "surabaya"}

	fmt.Println(alamat1)
	fmt.Println(alamat2)
	fmt.Println(alamat3)

	alamat4 := &alamat{"indonesia", "jawa timur", "surabaya"}
	change_alamat(alamat4)
	fmt.Println(alamat4)

	nama := Man{name: "sherdhan"}
	fmt.Println(nama)
	nama.married()
	fmt.Println(nama)
}
