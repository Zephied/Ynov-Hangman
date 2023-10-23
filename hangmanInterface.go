package hangmanclassic

import (
	"fmt"
	"os"
)

func HangmanInterface(hangman []byte, word string, hiddenWord string, life int, letterUsed []string, save bool) {
	var state1, state2 int
	var status bool
	var choice string

	fmt.Println("Good luck, you have", life, "attempts.")
	fmt.Println(hiddenWord)
	for life > 0 {
		fmt.Print("Choose:")
		fmt.Scan(&choice)
		if choice == "STOP" {
			Save(word, hiddenWord, life, letterUsed)
		}
		status = CheckUsedLetters(&letterUsed, choice)
		if status {
			hiddenWord, status = CheckLetter(word, hiddenWord, choice)
			if status {
				fmt.Println(hiddenWord)
				if hiddenWord == word {
					fmt.Println("You won!")
					if os.Args[1] == "--startWith" {
						os.Remove(os.Args[2])
					}
					os.Exit(0)
				}
			} else {
				if len(choice) > 1 {
					life = life - 2
				} else {
					life--
				}
				state1, state2, status = HangmanStats(life)
				if status {
					PrintHangman(state1, state2, hangman)
					if life == 0 {
						fmt.Println("You lost!")
						fmt.Println("The word was:", word)
						if os.Args[1] == "--startWith" {
							os.Remove(os.Args[2])
						}
						os.Exit(0)
					} else {
						fmt.Println("You have", life, "attempts left.")
					}
				}
			}
		}
	}
}

func HangmanInit(hangman []byte, words []string, wordsNb int) {
	life := 10
	var letterUsed []string
	var save bool = false

	word := RandomWord(words, wordsNb)
	hiddenWord, letter := HideWord(word)
	hiddenWord = Hint(hiddenWord, word, letter)

	HangmanInterface(hangman, word, hiddenWord, life, letterUsed, save)
}
