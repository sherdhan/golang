package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Runsync(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("test")
	time.Sleep(1 * time.Second)
}

func TestWaitgroup(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go Runsync(&group)
	}

	group.Wait()
	fmt.Println("done")
}
