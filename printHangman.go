package hangmanclassic

import "fmt"

func PrintHangman(start int, end int, hangman []byte) {
	fmt.Println(string(hangman[start:end]))
}
