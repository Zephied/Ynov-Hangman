package checks

import "fmt"

func CheckFile(err error) bool {
	status := true
	if err != nil {
		fmt.Println("no such file or directory")
		status = false
		return status
	} else {
		return status
	}
}
