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
	kflag      int
	nflag      bool
	rflag      bool
	uflag      bool
)

func init() {
	flag.StringVar(&filePath, "path", "/", "path to file")
	flag.StringVar(&outputPath, "outputPath", ".\\", "path to output file")
	flag.IntVar(&kflag, "k", -1, "указание колонки для сортировки")
	flag.BoolVar(&nflag, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&rflag, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&uflag, "u", false, "не выводить повторяющиеся строки")
}

func main() {
	flag.Parse()
	fmt.Println(filePath, kflag, nflag, rflag, uflag)
	fContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Split(strings.ReplaceAll(string(fContent), "\r", ""), "\n")
	if !nflag { // Сортировка по строкам
		if kflag >= 0 {
			if !rflag {
				//Сортировка по возрастанию
				sort.Slice(strs, func(i, j int) bool {
					//Проверка на то, что введеная длина для сортировки меньше чем количество столбцов в строке
					if len(strings.Split(strs[i], " ")) < kflag || len(strings.Split(strs[j], " ")) < kflag {
						log.Fatal("Out of range: kFlag is bigger than string size")
					}
					return len(strings.Split(strs[i], " ")[kflag]) < len(strings.Split(strs[j], " ")[kflag])
				})
			} else {
				//Сортировка по убыванию
				sort.Slice(strs, func(i, j int) bool {
					//Проверка на то, что введеная длина для сортировки меньше чем количество столбцов в строке
					if len(strings.Split(strs[i], " ")) < kflag || len(strings.Split(strs[j], " ")) < kflag {
						log.Fatal("Out of range: kFlag is bigger than string size")
					}
					return len(strings.Split(strs[i], " ")[kflag]) > len(strings.Split(strs[j], " ")[kflag])
				})
			}
		} else {
			//Если нет флага -k или он введен неправильно (отрицательный)
			if !rflag {
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
		if !rflag {
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

	if uflag {
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

// ParseString ...
func ParseString(s string) int {
	d, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
