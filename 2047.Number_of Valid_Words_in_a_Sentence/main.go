package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countValidWords(sentence string) int {
	//split the words by space
	split := strings.Split(sentence, " ")
	count := 0

	//in case of 2 spaces
	for _, word := range split {
		if word == "" {
			continue
		}

		var hyphen, islower, invalid bool
		wordLength := len(word)

		for i, letter := range word {
			if unicode.IsDigit(letter) {
				invalid = true
				break
			} else if letter == '-' {
				if !islower || i == wordLength-1 || hyphen || i < wordLength-1 && !unicode.IsLower(rune(word[i+1])) {
					invalid = true
					break
				}

				hyphen = true
			} else if unicode.IsPunct(letter) {
				if i != wordLength-1 {
					invalid = true
					break
				}

			} else {
				islower = true
			}
		}

		if !invalid {
			count++
		}
	}

	return count
}

func main() {
	senternce := "alice and  bob are playing stone-game10"
	fmt.Println(countValidWords(senternce))
}
