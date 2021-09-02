package utils


import (
	"bytes"
	"strings"
)

func StringConnectByBuffer(args ...string) string {
	var buffer bytes.Buffer
	for _, str := range args {
		buffer.WriteString(str)
	}

	return buffer.String()
}

func StringConnectByBuilder(args ...string) string {
	var builder strings.Builder

	for _, str := range args {
		builder.WriteString(str)
	}

	return builder.String()
}
