package main

import (
	. "hangman_classic"
	"os"
)

func main() {
	var wordlist []byte
	var status bool
	var words []string
	var wordsNb int
	var hangman []byte

	if CheckArgs() {
		filename := os.Args[1]
		wordlist, status = ReadFile(filename)
		if status {
			words, wordsNb = WordSlice(wordlist)
			hangman, status = ReadFile("content/hangman.txt")
			if status {
				wordsNb = wordsNb - 1
				HangmanInterface(hangman, words, wordsNb)
			}
		}
	}
}
