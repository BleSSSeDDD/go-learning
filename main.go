package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func isDigit(s byte) bool {
	return s >= '0' && s <= '9'
}

func isLetter(s byte) bool {
	if s >= 'a' && s <= 'z' {
		return true
	}
	return false
}

func toLover(s byte) byte {
	if s >= 'A' && s <= 'Z' {
		return s + 'a' - 'A'
	}
	return s
}

func isPalindrome(wg *sync.WaitGroup, s string) bool {
	defer wg.Done()
	i, j := 0, len(s)-1
	for i < j {
		if !(isLetter(toLover(s[i])) || isDigit(toLover(s[i]))) {
			i++
			continue
		} else if !(isLetter(toLover(s[j])) || isDigit(toLover(s[j]))) {
			j--
			continue
		} else if toLover(s[i]) == toLover(s[j]) {
			i++
			j--
			continue
		}
		return false
	}
	return true
}

func isPalindromeFaster(wg *sync.WaitGroup, s string) bool {
	defer wg.Done()
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !(isDigit(s[i]) || isLetter(toLover(s[i]))) {
			i++
		}
		for i < j && !(isDigit(s[j]) || isLetter(toLover(s[j]))) {
			j--
		}
		if toLover(s[i]) != toLover(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	base := "A man, a plan, a canal: Panama"
	var sb strings.Builder
	for i := 0; i < 100000; i++ {
		sb.WriteString(base)
	}
	s := sb.String()

	s += "X"

	fmt.Printf("Длина строки: %d символов\n", len(s))

	wg := &sync.WaitGroup{}
	wg.Add(2)

	start1 := time.Now()
	go func() {
		defer fmt.Printf("Версия 1: %v\n", time.Since(start1))
		result1 := isPalindrome(wg, s)
		fmt.Printf("Результат 1: %t\n", result1)
	}()

	start2 := time.Now()
	go func() {
		defer fmt.Printf("Версия 2: %v\n", time.Since(start2))
		result2 := isPalindromeFaster(wg, s)
		fmt.Printf("Результат 2: %t\n", result2)
	}()

	wg.Wait()
	fmt.Println("Готово!")
}
