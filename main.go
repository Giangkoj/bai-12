package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg = sync.WaitGroup{}
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)

	wg.Add(2)
	go func() {
		for {
			select {
			case i := <-ch1:
				fmt.Println("channel 1: %v\n", i)
			case j := <-ch2:
				fmt.Println("channel 2: %v\n", j)
			default:
				break
			}
		}
		wg.Done()
	}()
	go func() {
		ch1 <- 42
		close(ch1)
		ch2 <- 27
		close(ch2)
		wg.Done()

	}()

}
