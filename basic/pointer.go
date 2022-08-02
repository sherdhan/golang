package main

import "fmt"

type alamat struct {
	negara, provinsi, kota string
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
}
