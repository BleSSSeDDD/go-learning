package main

import (
	"fmt"
	"sync"
)

func writer() <-chan int {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i > -10; i-- {
			ch <- i
		}
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func doubler(ch <-chan int) <-chan int {
	ch1 := make(chan int)
	go func() {
		for v := range ch {
			ch1 <- v * 2
		}
		close(ch1)
	}()
	return ch1
}

func reader(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	reader(doubler(writer()))
}
