package hangmanclassic

import (
	"fmt"
	"os"
)

func CheckArgs() bool { // verifie que toute les conditions soit remplie et correct avant d'effectuer quoi que ce soit
	checked := true
	if len(os.Args) <= 2 && len(os.Args) > 5 {
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

func CheckLetter(data *Data, choice string) (string, bool) {
	var i int
	var status bool = false

	if len(choice) > 1 {
		if choice == data.Word {
			data.HiddenWord = data.Word
			status = true
			return data.HiddenWord, status
		} else {
			return data.HiddenWord, status
		}
	}
	for i < len(data.Word) {
		if string(data.Word[i]) == choice {
			data.HiddenWord = data.HiddenWord[:i] + choice + data.HiddenWord[i+1:]
			status = true
		}
		i++
	}
	return data.HiddenWord, status
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

func CheckRecurence(index int, hiddenWord string, use string, word string) string {
	for j := 0; j < len(hiddenWord); j++ {
		if string(word[j]) == use {
			hiddenWord = hiddenWord[:j] + use + hiddenWord[j+1:]
		}
	}
	return hiddenWord
}
