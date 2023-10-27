package init

import (
	"bufio"
	"fmt"
	structs "hangman_classic/structs"
	"os"
)

func InitAsciiArt(data *structs.Data) []string {
	file, err := os.Open(data.AsciiFile)
	if err != nil {
		fmt.Println("Error: cannot open file")
		os.Exit(1)
	}
	defer file.Close()

	var asciiArt []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var section string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			asciiArt = append(asciiArt, section)
			section = ""
		} else {
			section += line + "\n"
		}
	}
	return asciiArt
}
