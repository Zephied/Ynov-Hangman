package init

func WordSlice(wordlist []byte) ([]string, int) {
	var list []string
	var nb int
	var wordnb int
	i := 0
	for i < len(wordlist) {
		if wordlist[i] == '\n' {
			list = append(list, string(wordlist[nb:i]))
			nb = i + 1
			wordnb++
		}
		i++
	}
	return list, wordnb
}
