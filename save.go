package hangmanclassic

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(word string, hiddenWord string, life int, letterUsed []string) {
	b, err := json.Marshal(SaveStruct{word, hiddenWord, life, letterUsed})
	if CheckSave(err) {
		os.WriteFile("save.txt", b, 0644)
		fmt.Println("game saved in save.txt")
		os.Exit(0)
	}
}
