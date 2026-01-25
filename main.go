package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err1 := os.Open("text.txt")
	if err1 != nil {
		fmt.Println(f, err1)
	}

	scanner := bufio.NewScanner(f)

	scanner.Scan()

	fmt.Println(scanner.Text())

	defer f.Close()

	fr, err2 := os.ReadFile("text.txt")
	if err2 != nil {
		fmt.Println(fr, err2)
	}

	a := 0
	for _, c := range string(fr) {
		if c == '\n' {
			a++
		}
	}

	if fr[len(fr)-1] != '\n' {
		a++
	}

	fmt.Printf("Количество строк: %d", a)
}
