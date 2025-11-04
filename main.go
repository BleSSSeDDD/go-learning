package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type IntOrString interface {
	~int | ~string
}

type WorkerPool[T IntOrString] struct {
	ctx         context.Context
	numWorkers  int
	workersFunc func(T) T
	tasksChan   <-chan T
	resultChan  chan<- T
}

func (w WorkerPool[T]) Start(wMain *sync.WaitGroup) {
	fmt.Println("Начался w.Start()")
	wg := &sync.WaitGroup{}

	defer close(w.resultChan)
	defer wMain.Done()

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
					result := w.workersFunc(num)
					select {
					case w.resultChan <- result:
					case <-w.ctx.Done():
						fmt.Println("Воркер", goroutineID, "не смог отправить результат, выход по контексту")
						return
					}
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

func duplicateString(s string) string {
	return s + s
}

func generatorInt(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer func() {
			close(ch)
			fmt.Println("Генератор для Int закончил и закрыл канал")
		}()

		fmt.Println("Генератор для Int начал работу")
		for i := 1; i <= 100; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Выход из generator для Int по контексту")
				return
			case ch <- i:
				fmt.Println("Генератор для Int отправил ", i, " в канал")
			}
		}
	}()
	return ch
}

func generatorString(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		defer func() {
			close(ch)
			fmt.Println("Генератор для String закончил и закрыл канал")
		}()

		fmt.Println("Генератор для String начал работу")
		for i := 1; i < 100; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Выход из generator для String  по контексту")
				return
			case ch <- strconv.Itoa(i):
				fmt.Println("Генератор для String отправил ", i, " в канал")
			}
		}
	}()
	return ch
}

func readerInt(ctx context.Context, wMain *sync.WaitGroup, ch <-chan int) {
	defer fmt.Println("Чтение закончено, канал для Int закрыт")
	defer wMain.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Выход из reader для Int по контексту")
			return
		case num, exists := <-ch:
			if !exists {
				return
			}
			fmt.Printf("Прочитано значение %d в канале для Int\n", num)
		}
	}
}

func readerString(ctx context.Context, wMain *sync.WaitGroup, ch <-chan string) {
	defer fmt.Println("Чтение закончено, канал для String закрыт")
	defer wMain.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Выход из reader для String по контексту")
			return
		case num, exists := <-ch:
			if !exists {
				return
			}
			fmt.Printf("Прочитано значение %s для String\n", num)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Millisecond)
	defer cancel()

	start := time.Now()
	resultsInt := make(chan int)
	resultsString := make(chan string)

	wInt := WorkerPool[int]{ctx, 10, timesTwo, generatorInt(ctx), resultsInt}
	wString := WorkerPool[string]{ctx, 10, duplicateString, generatorString(ctx), resultsString}

	wg.Add(4)

	go readerString(ctx, wg, resultsString)
	go readerInt(ctx, wg, resultsInt)
	go wInt.Start(wg)
	go wString.Start(wg)

	wg.Wait()
	fmt.Println("Мейн закончился cпустя ", time.Since(start))
}
