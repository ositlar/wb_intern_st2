package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	f string
	d string
	s bool
)

func init() {
	flag.StringVar(&f, "f", "", "\"fields\" - выбрать поля (колонки)")
	flag.StringVar(&d, "d", "\t", "\"delimiter\" - использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "\"separated\" - только строки с разделителем")
}

func main() {
	flag.Parse()
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if s && !strings.Contains(line, d) {
			continue
		}
		cols := strings.Split(line, d)
		if len(f) > 0 {
			selectedCols := selectColumns(cols, f)
			fmt.Println(strings.Join(selectedCols, d))
		} else {
			fmt.Println(line)
		}
	}
}

func selectColumns(cols []string, fields string) []string {
	fieldIndices := parseFieldIndices(fields)
	if len(fieldIndices) > len(cols) {
		fieldIndices = fieldIndices[:len(cols)]
	}

	selectedCols := make([]string, len(fieldIndices))
	for i, idx := range fieldIndices {
		selectedCols[i] = cols[idx-1]
	}
	return selectedCols
}

func parseFieldIndices(fieldStr string) []int {
	fields := strings.Split(fieldStr, ",")
	fieldIndices := make([]int, len(fields))
	for i, f := range fields {
		fieldIndices[i], _ = strconv.Atoi(f)
	}
	return fieldIndices
}
