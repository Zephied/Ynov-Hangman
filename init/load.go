package init

import (
	"encoding/json"
	structs "hangman_classic/structs"
	"os"
)

func Load() {
	filename := os.Args[2]
	stats, status := ReadFile(filename)
	saved := true

	if status {
		hangman, status := ReadFile("content/hangman.txt")
		if status {
			var data structs.Data
			json.Unmarshal(stats, &data)
			HangmanInterface(hangman, &data, saved)
		}
	}
}
