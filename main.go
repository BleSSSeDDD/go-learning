package main

import (
	"fmt"
	"strings"
)

// Делает посимвольный сплит строки
// и идет с конца, пока первый символ не будет не пробелом
// потом начинает считать сколько не пробелов подряд
// и если появился пробел после того как уже было что-то в результате
// то возвращает результат
func lengthOfLastWord(s string) int {
	sarray := strings.Split(s, "")
	res := 0
	if len(sarray) == 1 {
		return 1
	}
	for i := len(sarray) - 1; i >= 0; i-- {
		if sarray[i] != " " {
			res++
			if i == 0 {
				return res
			}
			continue
		}
		if res != 0 {
			return res
		}
	}
	return 0
}

func main() {
	fmt.Println(lengthOfLastWord("a  фыф qwecqd sadasd     "))
}
