package main

import "fmt"

func main() {
	nilai := 90000000
	fmt.Println(nilai)

	nama := "ahmad"
	e := nama[0]
	estring := string(e)
	fmt.Println(estring)

	arr := [...]int{
		1, 2, 3,
	}

	sayhello()

	for key, value := range arr {
		fmt.Println("index ", key, " = ", value)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}

func sayhello() {
	fmt.Println("amang")
}
