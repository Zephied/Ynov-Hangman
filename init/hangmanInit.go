package init

import (
	structs "hangman_classic/structs"
)

func HangmanInit(hangman []byte, inite *structs.Inite, data *structs.Data) {
	data.Life = 10
	data.LetterUsed = []string{}

	var letter int

	save := false
	data.Word = RandomWord(inite.Words, inite.WordsNb)
	data.HiddenWord, letter = HideWord(data.Word)
	data.HiddenWord, data.LetterUsed = Hint(data, letter)

	HangmanInterface(hangman, data, save)
}
