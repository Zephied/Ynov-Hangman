package hangmanclassic

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(word string, hiddenWord string, life int, letterUsed []string) {
	var name string
	b, err := json.Marshal(SaveStruct{word, hiddenWord, life, letterUsed})
	if CheckSave(err) {
		fmt.Print("choose a name for your save: ")
		fmt.Scanln(&name)
		name = "save/" + name + ".txt"
		os.Create(name)
		os.WriteFile(name, b, 0644)
		fmt.Println("game saved in", name)
		os.Exit(0)
	}
}
