package main

import "fmt"

func buddyStrings(word string, goal string) bool {
	//check if the length of s and goal are the same
	if len(word) != len(goal) {
		return false
	}
	//make a slice with the different chars
	differences := make([]int, 0)
	dictionary := make(map[rune]int)
	for x, y := range word {
		dictionary[y]++
		if word[x] != goal[x] {
			differences = append(differences, x)
			if len(differences) > 2 {
				return false
			}
		}
	}
	
	if len(differences) == 0 {
		for _, num := range dictionary {
			if num > 1 {
				return true
			}
		}
		return false
	}

	if len(differences) == 1 {
		return false
	}
	//if the 2 different letters can be swapped and equal the goal, return true
	letters := []byte(word)
	letters[differences[0]], letters[differences[1]] = letters[differences[1]], letters[differences[0]]
	return string(letters) == goal
}

func main() {
	word := "ab"
	goal := "ba"
	fmt.Println(buddyStrings(word, goal))
}
