package checks

import (
	"fmt"
	"os"
)

func CheckArgs() bool { // verifie que toute les conditions soit remplie et correct avant d'effectuer quoi que ce soit
	checked := true
	if len(os.Args) <= 2 && len(os.Args) > 5 {
		fmt.Println("usage: go run src/main.go [option]\nif option is --startWith, you can load a game with save.txt at the second argument\nelse option has to be a file with words")
		checked = false
	}
	return checked
}
