package hangmanclassic

import (
	"math/rand"
)

func Hint(hiddenWord string, word string, letter int) string {
	var i int
	maxHint := letter/2 - 1
	change := false

	if maxHint < 1 {
		maxHint = 1
	}
	rng := rand.Intn(maxHint)
	for i < rng {
		index := rand.Intn(letter)
		if hiddenWord[index] == '_' {
			hiddenWord = hiddenWord[:index] + string(word[index]) + hiddenWord[index+1:]
			change = true
		}
		i++
	}
	if !change {
		index := rand.Intn(letter)
		hiddenWord = hiddenWord[:index] + string(word[index]) + hiddenWord[index+1:]
	}
	return hiddenWord
}
