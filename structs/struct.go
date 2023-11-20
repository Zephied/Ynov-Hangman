package structs

type Data struct {
	Word       string
	HiddenWord string
	AsciiFile  string
	Life       int
	LetterUsed []string
	Ascii      bool
}

type Inite struct {
	Wordlist []byte
	Words    []string
	WordsNb  int
	Status   bool
}
