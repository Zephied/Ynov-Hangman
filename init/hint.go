package init

import (
	check "hangman_classic/checks"
	structs "hangman_classic/structs"
	"math/rand"
)

func Hint(data *structs.Data, letter int) (string, []string) {
	var i int
	maxHint := letter/2 - 1
	change := false
	var use string

	if maxHint < 1 {
		maxHint = 1
	}
	rng := rand.Intn(maxHint)
	for i < rng {
		index := rand.Intn(letter)
		if data.HiddenWord[index] == '_' {
			data.HiddenWord = data.HiddenWord[:index] + string(data.Word[index]) + data.HiddenWord[index+1:]
			use = string(data.Word[index])
			data.HiddenWord = check.CheckRecurence(index, data, use)
			data.LetterUsed = append(data.LetterUsed, string(data.Word[index]))
			change = true
		}
		i++
	}
	if !change {
		index := rand.Intn(letter)
		data.HiddenWord = data.HiddenWord[:index] + string(data.Word[index]) + data.HiddenWord[index+1:]
		data.HiddenWord = check.CheckRecurence(index, data, use)
		data.LetterUsed = append(data.LetterUsed, string(data.Word[index]))

	}
	return data.HiddenWord, data.LetterUsed
}
