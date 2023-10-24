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
			var data Data
			json.Unmarshal(stats, &data)
			HangmanInterface(hangman, &data, saved)
		}
	}
}
