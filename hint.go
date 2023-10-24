package hangmanclassic

import (
	"math/rand"
)

func Hint(data *Data, letter int) (string, []string) {
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
			data.HiddenWord = CheckRecurence(index, data.HiddenWord, use, data.Word)
			data.LetterUsed = append(data.LetterUsed, string(data.Word[index]))
			change = true
		}
		i++
	}
	if !change {
		index := rand.Intn(letter)
		data.HiddenWord = data.HiddenWord[:index] + string(data.Word[index]) + data.HiddenWord[index+1:]
	}
	return data.HiddenWord, data.LetterUsed
}
