package checks

import (
	. "hangman_classic/structs"
)

func CheckRecurence(index int, data *Data, use string) string {
	for j := 0; j < len(data.HiddenWord); j++ {
		if string(data.Word[j]) == use {
			data.HiddenWord = data.HiddenWord[:j] + use + data.HiddenWord[j+1:]
		}
	}
	return data.HiddenWord
}
