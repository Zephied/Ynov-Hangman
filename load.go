package hangmanclassic

import (
	"encoding/json"
	"os"
)

func Load() {
	filename := os.Args[2]
	stats, status := ReadFile(filename)
	saved := true

	if status {
		hangman, status := ReadFile("content/hangman.txt")
		if status {
			var save SaveStruct
			json.Unmarshal(stats, &save)
			HangmanInterface(hangman, save.Word, save.HiddenWord, save.Life, save.LetterUsed, saved)
		}
	}
}
