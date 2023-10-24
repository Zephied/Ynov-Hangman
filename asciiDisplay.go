package hangmanclassic

func AsciiArtDisplay(data *Data) {
	//convert letter into ascii art using content/standard.txt and hiddenWord
	//display ascii art
	file, status := ReadFile(data.AsciiFile)
	if status {
		asciiArt := ConvertToAsciiArt(file, data.HiddenWord)
		for i := 0; i < len(asciiArt); i++ {
			for j := 0; j < len(asciiArt[i]); j++ {
				print(string(asciiArt[i][j]))
			}
			println()
		}
	}
}

func ConvertToAsciiArt(file []byte, hiddenWord string) [][]byte {
	//convert letter into ascii art using content/standard.txt and hiddenWord
	//return ascii art
	asciiArt := make([][]byte, 0)
	for i := 0; i < len(hiddenWord); i++ {
		asciiArt = append(asciiArt, make([]byte, 0))
	}
	for i := 0; i < len(file); i++ {
		if file[i] == 10 {
			continue
		}
		asciiArt[i%len(hiddenWord)] = append(asciiArt[i%len(hiddenWord)], file[i])
	}
	return asciiArt
}
