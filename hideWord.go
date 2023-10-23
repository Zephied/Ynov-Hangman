package hangmanclassic

func HideWord(word string) (string, int) {
	var hiddenWord string
	var i int
	for i = 0; i < len(word); i++ {
		hiddenWord += "_"
	}
	return hiddenWord, i
}
