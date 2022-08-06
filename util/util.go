package util

import (
	"DHT_NXT/consts"
	"fmt"
)

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

func BuildBuffer(msg_type int, buf []byte) []byte {
	new_buf := make([]byte, consts.FILE_BUF+1) // add space for 1
	// TODO: Find a more elegant solution to this?
	copy(new_buf, string(msg_type))
	copy(new_buf[1:], buf) // built buffer
	return new_buf
}
