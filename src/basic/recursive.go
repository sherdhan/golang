package main

import "fmt"

func factorial_loop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorial_recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial_recursive(value-1)
	}
}

func main() {
	loop := factorial_loop(5)
	fmt.Println("factorial", loop)

	recursive := factorial_recursive(5)
	fmt.Println("recursive ", recursive)
}
