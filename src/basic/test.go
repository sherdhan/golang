package main

import "fmt"

//Filter ...
type Filter func(string) string

// Blacklist ...
type Blacklist func(string) bool

func main() {

	nilai := 90000000
	fmt.Println(nilai)

	nama := "ahmad"
	e := nama[0]

	estring := string(e)
	fmt.Println(estring)

	arr := []int{
		1, 2, 3,
	}

	for key, value := range arr {
		fmt.Println("index ", key, " = ", value)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	total := sumall(10, 10, 10)
	fmt.Println(total)

	totalarr := sumall(arr...)
	fmt.Println(totalarr)

	sayhellowithfilter("test", namafilter)

	registeruser("ahmad", func(nama string) bool {
		return nama == "admin"
	})

	examplemap()

	runapp(true)
}

func sumall(nomor ...int) int {
	// contoh variadic
	total := 0
	for _, v := range nomor {
		total += v
	}
	return total
}

func callname(nama string) string {
	return "Halo " + nama
}

func sayhellowithfilter(name string, filter Filter) {
	namefiltered := filter(name)
	fmt.Println("Hello ", namefiltered)
}

func namafilter(nama string) string {
	if nama == "test" {
		return "..."
	}
	return nama
}

func registeruser(nama string, blacklist Blacklist) {
	if blacklist(nama) {
		fmt.Println("You are blocked", nama)
	} else {
		fmt.Println("Welcome ", nama)
	}
}

func examplemap(){
	data := map[string]string{
		"name" : "sherdhan",
		"address" : "malang",
	}
	datas := make(map[string]string)
	datas["name"] = "syarif" 
	datas["tes"] = "syarif"

	delete(datas, "tes")

	fmt.Println(data)
	fmt.Println(datas)
}

func runapp(status bool) {
	defer endapp()
	if status {
		panic("error pada status")
	}

}

func endapp() {
	message := recover()
	fmt.Println(message)
	fmt.Println("Aplikasi selesai")
}
