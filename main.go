package main

import (
	"fmt"
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
	s := "A man, a plan, a canal: Panama"

	wg := &sync.WaitGroup{}
	wg.Add(2)
	start := time.Now()
	go func() {
		fmt.Println(isPalindrome(wg, s))
		fmt.Println(time.Since(start))
	}()
	go func() {
		fmt.Println(isPalindromeFaster(wg, s))
		fmt.Println(time.Since(start))
	}()
	wg.Wait()
	fmt.Println("Готово!")
}
