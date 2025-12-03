package main

import (
	"fmt"
	"reflect"
)

// Это та самая функция IsZeroValue из сгенерированного кода
func IsZeroValue(val interface{}) bool {
	return val == nil || reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
}

func main() {
	// Симулируем проблему которая была в коде
	var fals bool = false
	var tru bool = true
	var str string = ""
	var num int = 0

	fmt.Println("=== Zero Value проверки ===")
	fmt.Printf("false: IsZeroValue(%v) = %v\n", fals, IsZeroValue(fals))
	fmt.Printf("true:  IsZeroValue(%v) = %v\n", tru, IsZeroValue(tru))
	fmt.Printf("string \"\": IsZeroValue(%q) = %v\n", str, IsZeroValue(str))
	fmt.Printf("int 0: IsZeroValue(%v) = %v\n", num, IsZeroValue(num))

	fmt.Println("\n=== Что считалось 'zero value' ===")
	fmt.Printf("reflect.Zero(reflect.TypeOf(false)) = %v\n", reflect.Zero(reflect.TypeOf(fals)).Interface())
	fmt.Printf("reflect.Zero(reflect.TypeOf(true)) = %v\n", reflect.Zero(reflect.TypeOf(tru)).Interface())

	fmt.Println("\n=== Проблема в валидации ===")
	fmt.Println("Для required поля 'is_active':")
	fmt.Println("- false → IsZeroValue(false) = true → ОШИБКА 'required field is zero value'")
	fmt.Println("- true  → IsZeroValue(true) = false → УСПЕХ")
}
