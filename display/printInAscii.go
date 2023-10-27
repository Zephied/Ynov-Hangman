package display

import (
	"bufio"
	"fmt"
	. "hangman_classic/structs"
	"strings"
)

func PrintInAscii(data *Data, asciiArt []string, chars []byte) {

	var j int
	asciifiedWord := make([]string, 8)
	for i := 0; i < len(data.HiddenWord); i++ {

		for j = 0; j < len(chars); j++ {
			if chars[j] == data.HiddenWord[i] {
				break
			}
		}

		scanner := bufio.NewScanner(strings.NewReader(asciiArt[j+1]))
		scanner.Split(bufio.ScanLines)

		for i := 0; scanner.Scan(); i++ {
			line := scanner.Text()
			asciifiedWord[i] += line

		}
	}
	for i := 0; i < len(asciifiedWord); i++ {
		fmt.Println(asciifiedWord[i])
	}
	fmt.Println(string(data.HiddenWord))
}
