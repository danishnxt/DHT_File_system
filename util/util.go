package util

import "fmt"

func CheckError(err_location string, err error) {
	if err != nil {
		if err_location != "" {
			fmt.Println(" ** Error ** ==> ", err_location)
		}
		fmt.Println(err)
	}
}

func UtilMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
