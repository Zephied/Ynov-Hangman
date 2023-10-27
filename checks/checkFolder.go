package checks

import "os"

func CheckFolder() {
	if _, err := os.Stat("save"); os.IsNotExist(err) {
		os.Mkdir("save", 0755)
	}
}
