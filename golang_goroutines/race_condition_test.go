package golanggoroutines

import (
	"fmt"
	"testing"
)

func TestRacecondition(t *testing.T) {
	x := 0

	for i := 1; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	fmt.Println("Counter = ", x)
}
