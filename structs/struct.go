package structs

type Data struct {
	Word       string
	HiddenWord string
	Life       int
	LetterUsed []string
	Ascii      bool
	AsciiFile  string
}

type Inite struct {
	Wordlist []byte
	Words    []string
	WordsNb  int
	Status   bool
}
