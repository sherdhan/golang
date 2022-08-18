package main

import "fmt"

func sum_all(nomor ...int) int {
	total := 0
	for _, v := range nomor {
		total += v
	}
	return total
}

func main() {
	// bentuk biasa
	total := sum_all(10, 10, 10, 10)
	fmt.Println(total)

	// bentuk array
	slice := []int{10, 10, 10, 10}
	total_slice := sum_all(slice...)
	fmt.Println(total_slice)
}
