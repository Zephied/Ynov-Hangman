package saving

import (
	"encoding/json"
	"fmt"
	. "hangman_classic/checks"
	. "hangman_classic/structs"
	"os"
)

func Save(data *Data) {
	var name string
	b, err := json.Marshal(data)
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
