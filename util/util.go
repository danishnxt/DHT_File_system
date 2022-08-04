package util

import "fmt"

func CheckError(err error) {
	// check if the file is ok
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}

func UtilMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
