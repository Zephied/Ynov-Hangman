package hangmanclassic

func CheckLetter(word string, hiddenWord string, choice string) (string, bool) {
	var i int
	var status bool = false

	if choice == word {
		hiddenWord = word
		status = true
		return hiddenWord, status
	}
	for i < len(word) {
		if string(word[i]) == choice {
			hiddenWord = hiddenWord[:i] + choice + hiddenWord[i+1:]
			status = true
		}
		i++
	}
	return hiddenWord, status
}
