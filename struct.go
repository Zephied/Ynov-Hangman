package hangmanclassic

type Data struct {
	Word       string
	HiddenWord string
	Life       int
	LetterUsed []string
	Ascii      bool
	AsciiFile  string
}
