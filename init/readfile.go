package init

import (
	check "hangman_classic/checks"
	"os"
)

func ReadFile(filename string) ([]byte, bool) { // lis le fichier qui lui est envoyer et renvois sous forme de tableau de mot
	var file []byte
	var err error
	file, err = os.ReadFile(filename)
	return file, check.CheckFile(err)
}
