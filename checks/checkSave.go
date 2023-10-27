package checks

import "fmt"

func CheckSave(err error) bool {
	status := true
	if err != nil {
		fmt.Println("error while saving")
		status = false
		return status
	} else {
		return status
	}
}
