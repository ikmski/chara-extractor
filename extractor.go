package main

import "sort"

func extract(text []string) []rune {

	m := make(map[rune]rune)

	for _, line := range text {

		chars := []rune(line)
		for _, char := range chars {

			_, ok := m[char]
			if !ok {
				m[char] = char
			}
		}
	}

	runes := []rune{}
	for key, _ := range m {
		runes = append(runes, key)
	}

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return runes
}
