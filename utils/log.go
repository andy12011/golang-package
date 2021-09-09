package utils

import "runtime"

func GetExcuteStackMessage() string {
	var buf []byte
	runtime.Stack(buf, false)

	return string(buf)
}