package main

import (
	"unicode"
	"fmt"
)

func main() {
	fmt.Println(Solution("Codility We test coders", 14))
	fmt.Println(Solution("The quick brown fox jumps over the lazy dog", 39))
	fmt.Println(Solution("Why not.", 6))
	fmt.Println(Solution("I can, but I wouldn't say it for her", 15))
	fmt.Println(Solution("I can, but I wouldn't say it for her", 22))
}

func Solution(message string, K int) string {

	if len(message) <= K {
		return message
	}

	init := K
	for {
		letter := message[init]
		if !unicode.IsLetter(rune(letter)) && string(letter) != "'" {
			break
		}
		init--
	}

	return message[:init]
}