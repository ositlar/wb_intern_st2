package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Flags
var (
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
)

func init() {
	flag.IntVar(&A, "A", 0, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&B, "B", 0, "\"before\" печатать +N строк до совпадения")
	flag.IntVar(&C, "C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&c, "c", false, "\"count\" (количество строк)")
	flag.BoolVar(&i, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&v, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&F, "F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	flag.BoolVar(&n, "n", false, "\"line num\", печатать номер строки")
}

func main() {
	flag.Parse()
	regex := os.Args[len(os.Args)-2]    //Фильтр
	filename := os.Args[len(os.Args)-1] //Файл для поиска строк

	file, err := os.ReadFile(filename)
	content := strings.Split(strings.ReplaceAll(string(file), "\r", ""), "\n")
	if err != nil {
		log.Fatal(err)
	}
	if A > 0 && B > 0 {
		log.Fatal("Некорректно введены флаги для выборки")
	}
	if n && c {
		log.Fatal("Некорректно введены флаги для выборки")
	}

	if c {
		var count int
		if v {
			if i {
				if F {
					for j := range content {
						if !strings.EqualFold(content[j], regex) {
							count++
						}
					}
					fmt.Println(count)
					os.Exit(1)
				} else {
					for j := range content {
						if !strings.Contains(strings.ToLower(content[j]), strings.ToLower(regex)) {
							count++
						}
					}
					fmt.Println(count)
					os.Exit(1)
				}
			} else {
				if F {
					for j := range content {
						if content[j] != regex {
							count++
						}
					}
					fmt.Println(count)
					os.Exit(1)
				} else {
					for j := range content {
						if !strings.Contains(content[j], regex) {
							count++
						}
					}
					fmt.Println(count)
					os.Exit(1)
				}
			}
		} else {
			if i {
				if F {
					for j := range content {
						if strings.EqualFold(content[j], regex) {
							count++
						}
					}
					fmt.Println(count)
					os.Exit(1)
				} else {
					for j := range content {
						if strings.Contains(strings.ToLower(content[j]), strings.ToLower(regex)) {
							count++
						}
					}
					fmt.Println(count)
					os.Exit(1)
				}
			} else {
				if F {
					for j := range content {
						if content[j] == regex {
							count++
						}
					}
					fmt.Println(count)
					os.Exit(1)
				} else {
					for j := range content {
						if strings.Contains(content[j], regex) {
							count++
						}
					}
					fmt.Println(count)
					os.Exit(1)
				}
			}
		}

	}

	if v {
		if !F {
			//Поиск по паттерну
			for _, s := range content {
				if !strings.Contains(s, regex) {
					fmt.Println(s)
				}
			}
			os.Exit(1)
		} else {
			//Поиск по точному совпадению
			for _, s := range content {
				if !strings.EqualFold(s, regex) {
					fmt.Println(s)
				}
			}
			os.Exit(1)
		}

	} else {
		if !F {
			//Поиск по паттерну
			if A > 0 {
				if !n {
					if targetIndex := FindIndex(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex; i <= targetIndex+A && i < len(content); i++ {
							fmt.Println(content[i])
						}
					}
				} else {
					if targetIndex := FindIndex(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex; i <= targetIndex+A && i < len(content); i++ {
							fmt.Println(i)
						}
					}
				}

			}
			if B > 0 {
				if !n {
					if targetIndex := FindIndex(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex; i >= targetIndex-B && i >= 0; i-- {
							fmt.Println(content[i])
						}
					}
				} else {
					if targetIndex := FindIndex(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex; i >= targetIndex-B && i >= 0; i-- {
							fmt.Println(i)
						}
					}
				}

			}
			if C > 0 {
				if !n {
					if targetIndex := FindIndex(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex - C; i >= targetIndex+C && i >= 0 && i < len(content); i++ {
							fmt.Println(content[i])
						}
					}
				} else {
					if targetIndex := FindIndex(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex - C; i >= targetIndex+C && i >= 0 && i < len(content); i++ {
							fmt.Println(i)
						}
					}
				}
			}
		} else {
			//Поиск по точному совпадению
			if A > 0 {
				if !n {
					if targetIndex := FindIndexFixed(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex; i <= targetIndex+A && i < len(content); i++ {
							fmt.Println(content[i])
						}
					}
				} else {
					if targetIndex := FindIndexFixed(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex; i <= targetIndex+A && i < len(content); i++ {
							fmt.Println(i)
						}
					}
				}
			}
			if B > 0 {
				if !n {
					if targetIndex := FindIndexFixed(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex; i >= targetIndex-B && i >= 0; i-- {
							fmt.Println(i)
						}
					}
				} else {
					if targetIndex := FindIndexFixed(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex; i >= targetIndex-B && i >= 0; i-- {
							fmt.Println(i)
						}
					}
				}
			}
			if C > 0 {
				if !n {
					if targetIndex := FindIndexFixed(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex - C; i >= targetIndex+C && i >= 0 && i < len(content); i++ {
							fmt.Println(content[i])
						}
					}
				} else {
					if targetIndex := FindIndexFixed(content, regex); targetIndex == -1 {
						fmt.Println("Строка не найдена")
						os.Exit(1)
					} else {
						for i := targetIndex - C; i >= targetIndex+C && i >= 0 && i < len(content); i++ {
							fmt.Println(i)
						}
					}
				}
			}
		}
	}

}

// FindIndex ...
func FindIndex(content []string, target string) int {
	for j := range content {
		if i {
			if strings.Contains(strings.ToLower(content[j]), strings.ToLower(target)) {
				return j
			}
		} else {
			if strings.Contains(content[j], target) {
				return j
			}
		}

	}
	return -1
}

// FindIndexFixed ...
func FindIndexFixed(content []string, target string) int {
	for j := range content {
		if i {
			if strings.EqualFold(content[j], target) {
				return j
			}
		} else {
			if content[j] == target {
				return j
			}
		}
	}
	return -1
}
