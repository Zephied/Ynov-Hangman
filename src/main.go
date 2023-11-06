package main

import (
	checks "hangman_classic/checks"
	display "hangman_classic/display"
	inite "hangman_classic/init"
	structs "hangman_classic/structs"
	"os"
)

func main() {
	var wordlist []byte
	var status bool
	var words []string
	var wordsNb int
	var hangman []byte

	var data structs.Data

	data.Ascii = false
	if checks.CheckArgs() {
		if os.Args[1] == "--startWith" {
			inite.Load()
		}
		if os.Args[1] == "--help" {
			display.Help()
			os.Exit(0)
		}
		if len(os.Args) > 2 && len(os.Args) < 4 {
			data.Ascii = true
			data.AsciiFile = os.Args[2]
		}
		filename := os.Args[1]
		wordlist, status = inite.ReadFile(filename)
		if status {
			words, wordsNb = inite.WordSlice(wordlist)
			hangman, status = inite.ReadFile("content/hangman.txt")
			if status {
				wordsNb = wordsNb - 1
				checks.CheckFolder()
				inite.HangmanInit(hangman, words, wordsNb, &data)
			}
		}
	}
}
