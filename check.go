package hangmanclassic

import (
	"fmt"
	"os"
)

func CheckArgs() bool { // verifie que toute les conditions soit remplie et correct avant d'effectuer quoi que ce soit
	checked := true
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("usage: go run src/main.go [option]\nif option is --startWith, you can load a game with save.txt at the second argument\nelse option has to be a file with words")
		checked = false
	}
	return checked
}

func CheckFile(err error) bool {
	status := true
	if err != nil {
		fmt.Println("no such file or directory")
		status = false
		return status
	} else {
		return status
	}
}

func CheckSave(err error) bool {
	status := true
	if err != nil {
		fmt.Println("error while saving")
		status = false
		return status
	} else {
		return status
	}
}

func CheckLetter(word string, hiddenWord string, choice string) (string, bool) {
	var i int
	var status bool = false

	if len(choice) > 1 {
		if choice == word {
			hiddenWord = word
			status = true
			return hiddenWord, status
		} else {
			return hiddenWord, status
		}
	}
	for i < len(word) {
		if string(word[i]) == choice {
			hiddenWord = hiddenWord[:i] + choice + hiddenWord[i+1:]
			status = true
		}
		i++
	}
	return hiddenWord, status
}

func CheckUsedLetters(letterUsed *[]string, choice string) bool {
	var i int
	for _, letter := range *letterUsed {
		if choice == letter {
			if len(choice) > 1 {
				fmt.Println("You already used this word.")
			} else {
				fmt.Println("You already used this letter.")
			}
			return false
		} else {
			i++
		}
	}
	*letterUsed = append(*letterUsed, choice)
	return true
}

func CheckFolder() {
	if _, err := os.Stat("save"); os.IsNotExist(err) {
		os.Mkdir("save", 0755)
	}
}
