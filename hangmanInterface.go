package hangmanclassic

import "fmt"

func HangmanInterface(hangman []byte, words []string, wordsNb int) {
	life := 10
	var state1, state2 int
	var status bool
	var choice string

	fmt.Println("Good luck, you have", life, "attempts.")
	word := RandomWord(words, wordsNb)
	hiddenWord, letter := HideWord(word)
	hiddenWord = Hint(hiddenWord, word, letter)
	fmt.Println(hiddenWord)
	for life > 0 {
		fmt.Print("Choose:")
		fmt.Scan(&choice)
		hiddenWord, status = CheckLetter(word, hiddenWord, choice)
		if status {
			fmt.Println(hiddenWord)
			if hiddenWord == word {
				fmt.Println("You won!")
				break
			}
		} else {
			life--
			state1, state2, status = HangmanStats(life)
			if status {
				PrintHangman(state1, state2, hangman)
				if life == 0 {
					fmt.Println("You lost!")
					fmt.Println("The word was:", word)
				} else {
					fmt.Println("You have", life, "attempts left.")
				}
			}
		}
	}
}
