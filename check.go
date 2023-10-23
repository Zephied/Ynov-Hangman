package hangmanclassic

import (
	"fmt"
	"os"
)

func CheckArgs() bool { // verifie que toute les conditions soit remplie et correct avant d'effectuer quoi que ce soit
	checked := true
	if len(os.Args) != 2 {
		fmt.Println("wrong number of arguments\nhangman require [wordlist] for first argument")
		checked = false
	}
	return checked
}

func CheckFile(err error) bool {
	status := true
	if err != nil {
		fmt.Println("no such file or directory")
		status = false
		return status
	} else {
		return status
	}
}
