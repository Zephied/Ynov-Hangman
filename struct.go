package hangmanclassic

type SaveStruct struct {
	Word       string
	HiddenWord string
	Life       int
	LetterUsed []string
}
