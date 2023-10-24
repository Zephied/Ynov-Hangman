package hangmanclassic

import (
	"fmt"
	"os"
)

func HangmanInterface(hangman []byte, data *Data, save bool) {
	var state1, state2 int
	var status bool
	var choice string

	fmt.Println("Good luck, you have", data.Life, "attempts.")
	if data.Ascii {
		AsciiArtDisplay(data)
	} else {
		fmt.Println(data.HiddenWord)
	}
	for data.Life > 0 {
		fmt.Print("Choose:")
		fmt.Scan(&choice)
		if choice == "STOP" {
			Save(data)
		}
		status = CheckUsedLetters(&data.LetterUsed, choice)
		if status {
			data.HiddenWord, status = CheckLetter(data, choice)
			if status {
				if !data.Ascii {
					fmt.Println(data.HiddenWord)
				} else {
					AsciiArtDisplay(data)

				}
				if data.HiddenWord == data.Word {
					fmt.Println("You won!")
					if os.Args[1] == "--startWith" {
						os.Remove(os.Args[2])
					}
					os.Exit(0)
				}
			} else {
				if len(choice) > 1 {
					data.Life = data.Life - 2
				} else {
					data.Life--
				}
				state1, state2, status = HangmanStats(data.Life)
				if status {
					PrintHangman(state1, state2, hangman)
					if data.Life == 0 {
						fmt.Println("You lost!")
						fmt.Println("The word was:", data.Word)
						if os.Args[1] == "--startWith" {
							os.Remove(os.Args[2])
						}
						os.Exit(0)
					} else {
						fmt.Println("You have", data.Life, "attempts left.")
					}
				}
			}
		}
	}
}

func HangmanInit(hangman []byte, words []string, wordsNb int, data *Data) {
	data.Life = 10
	data.LetterUsed = []string{}

	var letter int

	save := false
	data.Word = RandomWord(words, wordsNb)
	data.HiddenWord, letter = HideWord(data.Word)
	data.HiddenWord, data.LetterUsed = Hint(data, letter)

	HangmanInterface(hangman, data, save)
}
