package hangmanclassic

import (
	"math/rand"
)

func RandomWord(words []string, wordsNb int) string {
	index := rand.Intn(wordsNb)
	return words[index]
}
