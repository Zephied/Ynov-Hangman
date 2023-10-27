package checks

import (
	. "hangman_classic/structs"
)

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
