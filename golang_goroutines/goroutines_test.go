package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHello() {
	fmt.Println("Test Hello")
}

func TestCreategoroutines(t *testing.T) {
	go RunHello()
	fmt.Println("cek goroutines")

	time.Sleep(1 * time.Second)
}

func Displaynumber(number int) {
	fmt.Println("nomor ke", number)
}

func TestManygoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go Displaynumber(i)
	}

	time.Sleep(1 * time.Second)
}
