package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var (
	filePath   string
	outputPath string
	kFlag      int
	nFlag      bool
	rFlag      bool
	uFlag      bool
)

func init() {
	flag.StringVar(&filePath, "path", "/", "path to file")
	flag.StringVar(&outputPath, "outputPath", ".\\", "path to output file")
	flag.IntVar(&kFlag, "k", -1, "указание колонки для сортировки")
	flag.BoolVar(&nFlag, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&rFlag, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&uFlag, "u", false, "не выводить повторяющиеся строки")
}

func main() {
	flag.Parse()
	fmt.Println(filePath, kFlag, nFlag, rFlag, uFlag)
	fContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Split(strings.ReplaceAll(string(fContent), "\r", ""), "\n")
	if !nFlag { // Сортировка по строкам
		if kFlag >= 0 {
			if !rFlag {
				//Сортировка по возрастанию
				sort.Slice(strs, func(i, j int) bool {
					//Проверка на то, что введеная длина для сортировки меньше чем количество столбцов в строке
					if len(strings.Split(strs[i], " ")) < kFlag || len(strings.Split(strs[j], " ")) < kFlag {
						log.Fatal("Out of range: kFlag is bigger than string size")
					}
					return len(strings.Split(strs[i], " ")[kFlag]) < len(strings.Split(strs[j], " ")[kFlag])
				})
			} else {
				//Сортировка по убыванию
				sort.Slice(strs, func(i, j int) bool {
					//Проверка на то, что введеная длина для сортировки меньше чем количество столбцов в строке
					if len(strings.Split(strs[i], " ")) < kFlag || len(strings.Split(strs[j], " ")) < kFlag {
						log.Fatal("Out of range: kFlag is bigger than string size")
					}
					return len(strings.Split(strs[i], " ")[kFlag]) > len(strings.Split(strs[j], " ")[kFlag])
				})
			}
		} else {
			//Если нет флага -k или он введен неправильно (отрицательный)
			if !rFlag {
				//Сортировка по возрастанию
				sort.Slice(strs, func(i, j int) bool {
					return len(strs[i]) < len(strs[j])
				})
			} else {
				//Сортировка по убыванию
				sort.Slice(strs, func(i, j int) bool {
					return len(strs[i]) > len(strs[j])
				})
			}
		}

	} else { //Сортировка по числам
		if !rFlag {
			//Сортировка по возрастанию
			sort.Slice(strs, func(i, j int) bool {
				return ParseString(strs[i]) < ParseString(strs[j])
			})
		} else {
			//Сортировка по убыванию
			sort.Slice(strs, func(i, j int) bool {
				return ParseString(strs[i]) > ParseString(strs[j])
			})
		}
	}

	if uFlag {
		strs = slices.Compact(strs)
	}
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, s := range strs {
		fmt.Println(s)
		file.WriteString(s + "\n")
	}
}

func ParseString(s string) int {
	d, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func SliceToSet(strs []string) []string {
	set := make(map[string]struct{})
	result := make([]string, 0)
	for _, value := range strs {
		set[value] = struct{}{}
	}
	for key := range set {
		result = append(result, key)
	}
	return result
}
