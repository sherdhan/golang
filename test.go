package main

import "fmt"

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

	sayhello()

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

}

func sayhello() {
	fmt.Println("amang")
}

func sumall(nomor ...int) int {
	total := 0
	for _, v := range nomor {
		total += v
	}
	return total
}
