package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	count := countWords("asdhkajshdkjh")
	fmt.Println(count)
}

func countWords(text string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.Fields(text)
	for _, word := range words {
		cleanWord := strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return r
			}
			return -1
		}, strings.ToLower(word))

		if cleanWord != "" {
			wordCount[cleanWord]++
		}
	}

	return wordCount
}
