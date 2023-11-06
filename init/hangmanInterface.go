package init

import (
	"fmt"
	check "hangman_classic/checks"
	display "hangman_classic/display"
	saving "hangman_classic/saving"
	structs "hangman_classic/structs"
	"os"
)

func HangmanInterface(hangman []byte, data *structs.Data, save bool) {
	var state1, state2 int
	var status bool
	var choice string
	var asciiArt []string
	var chars []byte = []byte{' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ':', ';', '<', '=', '>', '?', '@', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '[', '\\', ']', '^', '_', '`', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '{', '|', '}', '~'}

	fmt.Println("Good luck, you have", data.Life, "attempts.")
	if data.Ascii {
		asciiArt = InitAsciiArt(data)
		display.PrintInAscii(data, asciiArt, chars)
	} else {
		fmt.Println(data.HiddenWord)
	}
	for data.Life > 0 {
		fmt.Print("Choose:")
		fmt.Scan(&choice)
		if choice == "STOP" {
			saving.Save(data)
		}
		status = check.CheckUsedLetters(&data.LetterUsed, choice)
		if status {
			data.HiddenWord, status = check.CheckLetter(data, choice)
			if status {
				if !data.Ascii {
					fmt.Println(data.HiddenWord)
				} else {
					display.PrintInAscii(data, asciiArt, chars)
				}
				if data.HiddenWord == data.Word {
					fmt.Println("You won!")
					if os.Args[1] == "--startWith" {
						os.Remove(os.Args[2])
					}
					os.Exit(0)
				}
			} else {
				if len(choice) > 1 {
					data.Life = data.Life - 2
				} else {
					data.Life--
				}
				state1, state2, status = structs.HangmanStats(data.Life)
				if status {
					display.PrintHangman(state1, state2, hangman)
					if data.Ascii {
						display.PrintInAscii(data, asciiArt, chars)
					} else {
						fmt.Println(data.HiddenWord)
					}
					if data.Life == 0 {
						fmt.Println("You lost!")
						fmt.Println("The word was:", data.Word)
						if os.Args[1] == "--startWith" {
							os.Remove(os.Args[2])
						}
						os.Exit(0)
					} else {
						fmt.Println("You have", data.Life, "attempts left.")
					}
				}
			}
		}
	}
}
