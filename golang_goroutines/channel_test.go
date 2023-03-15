package golanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreatechannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "sherdhan"
		fmt.Println("go func done")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(1 * time.Second)
}

func Getresponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "data sherdhan"
}

func TestGetrespinse(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go Getresponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func Onlyin(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "sherdhan"
	fmt.Println("in data done")
}

func Onlyout(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChannelinout(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go Onlyin(channel)
	go Onlyout(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedchannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "sherdhan"
	channel <- "syarif"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func TestRangechannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("done")
}

func TestSelectchannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go Getresponse(channel1)
	go Getresponse(channel2)

	count := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 adalah ", data)
			count++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 adalah ", data)
			count++
		default:
			fmt.Println("menunggu data")
		}
		if count == 2 {
			break
		}
	}

}
