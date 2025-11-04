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
	fmt.Println("Начался w.Start()")
	wg := &sync.WaitGroup{}

	defer close(w.resultChan)

	for i := 0; i < w.numWorkers; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for {
				select {
				case <-w.ctx.Done():
					fmt.Println("Выход из воркера", goroutineID, " по контексту")
					return
				case num, ok := <-w.tasksChan:
					if !ok {
						fmt.Println("Воркер ", goroutineID, " закончил работу")
						return
					}
					w.resultChan <- w.workersFunc(num)
				}
			}
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

func generator(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer func() {
			close(ch)
			fmt.Println("Генератор закончил и закрыл канал")
		}()

		fmt.Println("Генератор начал работу")
		for i := 1; i <= 1000; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Выход из generator по контексту")
				return
			case ch <- i:
				fmt.Println("Генератор отправил ", i, " в канал")
			}
		}
	}()
	return ch
}

func reader(ctx context.Context, ch <-chan int) {
	defer fmt.Println("Чтение закончено, канал закрыт")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Выход из reader по контексту")
			return
		case num, exists := <-ch:
			if !exists {
				return
			}
			fmt.Printf("Прочитано значение %d\n", num)
		}
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	start := time.Now()
	results := make(chan int)
	w := WorkerPool{ctx, 10, timesTwo, generator(ctx), results}

	go w.Start()

	reader(ctx, results)

	fmt.Println("Мейн закончился cпустя ", time.Since(start))
}
