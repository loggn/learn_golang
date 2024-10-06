package utils

import "strings"

func SpliString(str string) []string {
	arr := strings.Split(str, " ")
	return arr
}

func WordCount(str string) map[string]int {
	arr := strings.Split(str, " ")
	m := make(map[string]int)
	a:
	for _, value := range arr{
		for str := range m {
			if value == str {
				m[value]++
				continue a
			}
		}
		m[value] = 1
	}
	return m
}