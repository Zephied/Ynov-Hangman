package display

func Help() {
	println("Usage: ./hangman [OPTION]... [FILE]...\nOPTION can be a file path of word or --startWith [FILE] to load a saved game.\nFILE is the path of the file containing the words.\nyou can add [pathfile] to display the hangman in ascii art.\n--help display this help.")
}
