package checks

import "fmt"

func CheckUsedLetters(letterUsed *[]string, choice string) bool {
	var i int
	for _, letter := range *letterUsed {
		if choice == letter {
			if len(choice) > 1 {
				fmt.Println("You already used this word.")
			} else {
				fmt.Println("You already used this letter.")
			}
			return false
		} else {
			i++
		}
	}
	*letterUsed = append(*letterUsed, choice)
	return true
}
