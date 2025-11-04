package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	ctx         context.Context
	numWorkers  int
	workersFunc func(int) int
	tasksChan   <-chan int
	resultChan  chan<- int
}

func (w WorkerPool) Start() {
	wg := &sync.WaitGroup{}

	defer close(w.resultChan)

	for i := 0; i < w.numWorkers; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			for num := range w.tasksChan {
				fmt.Println("Воркер ", goroutineID, " записал в канал с результатами ", w.workersFunc(num), " при исходном числе ", num)
				fmt.Print("")
				w.resultChan <- w.workersFunc(num)
			}
			wg.Done()
		}(i + 1)
	}

	wg.Wait()
	fmt.Println("Конец w.Start()")
}

func timesTwo(num int) int {
	return num * 2
}

func squared(num int) int {
	return num * num
}

func generator() <-chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("Генератор начал работу")
		for i := 1; i <= 10000; i++ {
			ch <- i
		}
		close(ch)
		fmt.Println("Генератор закончил и закрыл канал")
	}()
	return ch
}

func reader(ch <-chan int) {
	for i := range ch {
		fmt.Printf("Прочитано значение %d\n", i)
		fmt.Print("")
	}
	fmt.Println("Чтение закончено, канал закрыт")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	start := time.Now()
	results := make(chan int)
	w := WorkerPool{ctx, 100, timesTwo, generator(), results}

	go w.Start()

	reader(results)

	fmt.Println("Мейн закончился cпустя ", time.Since(start))
}
