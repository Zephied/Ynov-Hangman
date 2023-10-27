package init

import (
	structs "hangman_classic/structs"
)

func HangmanInit(hangman []byte, words []string, wordsNb int, data *structs.Data) {
	data.Life = 10
	data.LetterUsed = []string{}

	var letter int

	save := false
	data.Word = RandomWord(words, wordsNb)
	data.HiddenWord, letter = HideWord(data.Word)
	data.HiddenWord, data.LetterUsed = Hint(data, letter)

	HangmanInterface(hangman, data, save)
}
