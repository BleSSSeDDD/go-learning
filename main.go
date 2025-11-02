package main

import (
	"fmt"
	"sync"
	"time"
)

func writer() <-chan int {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Writer 1 writes i: ", i)
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i > -5; i-- {
			fmt.Println("Writer 2 writes i: ", i)
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
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		count := 0
		for v := range ch {
			count++
			fmt.Printf("Doubler 1 received %d (total: %d)\n", v, count)
			ch1 <- v * 2
		}
		fmt.Printf("Doubler 1 finished, processed %d values\n", count)
	}()
	go func() {
		defer wg.Done()
		count := 0
		for v := range ch {
			count++
			fmt.Printf("Doubler 2 received %d (total: %d)\n", v, count)
			ch1 <- v * 2
		}
		fmt.Printf("Doubler 2 finished, processed %d values\n", count)
	}()
	go func() {
		wg.Wait()
		close(ch1)
	}()

	return ch1
}

func reader(ch <-chan int) {
	for i := range ch {
		fmt.Println("Reader recieves: ", i)
	}
}

func main() {
	reader(doubler(writer()))
	time.Sleep(2 * time.Second)
}
