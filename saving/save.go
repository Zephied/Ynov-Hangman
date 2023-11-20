package saving

import (
	"encoding/json"
	"fmt"
	checks "hangman_classic/checks"
	structs "hangman_classic/structs"
	"os"
)

func Save(data *structs.Data, save bool) {
	var name string
	b, err := json.Marshal(data)
	if checks.CheckSave(err) {
		if !save {
			fmt.Print("choose a name for your save: ")
			fmt.Scanln(&name)
			name = "save/" + name + ".txt"
			os.Create(name)
		} else {
			name = os.Args[2]
		}
		os.WriteFile(name, b, 0644)
		fmt.Println("game saved in", name)
		os.Exit(0)
	}
}
