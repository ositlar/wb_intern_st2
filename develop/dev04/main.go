package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func SortString(s string) string {
	runes := strings.Split(s, "")
	sort.Strings(runes)
	return strings.Join(runes, "")
}

func RemoveDuplicates(words []string) []string {
	seen := make(map[string]struct{}, len(words))
	j := 0
	for _, v := range words {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		words[j] = v
		j++
	}
	return words[:j]
}

func main() {
	output := make(map[string][]string)
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(input)
	var flag bool
	for _, s := range input {
		flag = false
		w := strings.ToLower(s)
		sorted := SortString(w)
		for key := range output {
			keySorted := SortString(key)
			if keySorted == sorted {
				output[key] = append(output[key], s)
				flag = true
				break
			}
		}
		if flag {
			continue
		} else {
			output[s] = append(output[s], s)
		}
	}
	for key := range output {
		sort.Strings(output[key])
		output[key] = RemoveDuplicates(output[key])
		if len(output[key]) == 1 {
			delete(output, key)
		}
	}
	fmt.Println(output)
}

// func findAnagramGroups(words []string) map[string][]string {
// 	groups := make(map[string][]string)

// 	for _, word := range words {
// 		added := false

// 		word = strings.ToLower(word)
// 		sorted := sortLetters(word)
// 		for key := range groups {
// 			sortedKey := sortLetters(key)
// 			if sortedKey == sorted {
// 				groups[key] = append(groups[key], word)
// 				added = true
// 				break
// 			}
// 		}
// 		if added {
// 			continue
// 		} else {
// 			groups[word] = append(groups[word], word)
// 		}
// 	}

// 	for key := range groups {
// 		sort.Strings(groups[key])
// 		groups[key] = removeDuplicates(groups[key])
// 		if len(groups[key]) == 1 {
// 			delete(groups, key)
// 		}
// 	}

// 	return groups
// }

// func sortLetters(word string) string {
// 	letters := strings.Split(word, "")
// 	sort.Strings(letters)
// 	return strings.Join(letters, "")
// }

// func removeDuplicates(words []string) []string {
// 	seen := make(map[string]struct{}, len(words))
// 	j := 0
// 	for _, v := range words {
// 		if _, ok := seen[v]; ok {
// 			continue
// 		}
// 		seen[v] = struct{}{}
// 		words[j] = v
// 		j++
// 	}
// 	return words[:j]
// }

// func main() {
// 	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
// 	result := findAnagramGroups(words)
// 	fmt.Println(result)
// }
