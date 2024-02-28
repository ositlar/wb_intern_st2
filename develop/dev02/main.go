package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var s string
	flag := false
	fmt.Scan(&s)
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if flag {
			count, _ := strconv.Atoi(string(s[i]))
			for j := 0; j < count; j++ {
				//result = append(result, rune(s[i-1]))
				sb.WriteRune(rune(s[i-1]))
			}
			flag = false
		} else if s[i] == '\\' {
			flag = true
		} else if unicode.IsDigit(rune(s[i])) {
			//return "" // Некорректная строка, если встречено число без предшествующего символа
			fmt.Println("Некорректная строка, если встречено число без предшествующего символа")
			os.Exit(1)
		} else if i < len(s)-1 && unicode.IsDigit(rune(s[i+1])) {
			count, _ := strconv.Atoi(string(s[i+1]))
			for j := 0; j < count; j++ {
				sb.WriteRune(rune(s[i]))
			}
			i++
		} else {
			sb.WriteRune(rune(s[i]))
		}
	}
	fmt.Println(sb.String())
}
