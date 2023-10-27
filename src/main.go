package main

import (
	. "hangman_classic/checks"
	. "hangman_classic/display"
	. "hangman_classic/init"
	. "hangman_classic/structs"
	"os"
)

func main() {
	var wordlist []byte
	var status bool
	var words []string
	var wordsNb int
	var hangman []byte

	var data Data

	data.Ascii = false
	if CheckArgs() {
		if os.Args[1] == "--startWith" {
			Load()
		}
		if os.Args[1] == "--help" {
			Help()
			os.Exit(0)
		}
		if len(os.Args) > 2 && len(os.Args) < 4 {
			data.Ascii = true
			data.AsciiFile = os.Args[2]
		}
		filename := os.Args[1]
		wordlist, status = ReadFile(filename)
		if status {
			words, wordsNb = WordSlice(wordlist)
			hangman, status = ReadFile("content/hangman.txt")
			if status {
				wordsNb = wordsNb - 1
				CheckFolder()
				HangmanInit(hangman, words, wordsNb, &data)
			}
		}
	}
}
