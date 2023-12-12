package main

import (
	checks "hangman_classic/checks"
	display "hangman_classic/display"
	initiate "hangman_classic/init"
	structs "hangman_classic/structs"
	"os"
)

func main() {
	var hangman []byte

	var data structs.Data
	var Inite structs.Inite

	data.Ascii = false
	if checks.CheckArgs() {
		if os.Args[1] == "--startWith" {
			initiate.Load()
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
		Inite.Wordlist, Inite.Status = initiate.ReadFile(filename)
		if Inite.Status {
			Inite.Words, Inite.WordsNb = initiate.WordSlice(Inite.Wordlist)
			hangman, Inite.Status = initiate.ReadFile("content/hangman.txt")
			if Inite.Status {
				Inite.WordsNb = Inite.WordsNb - 1
				checks.CheckFolder()
				initiate.HangmanInit(hangman, &Inite, &data)
			}
		}
	}
}
