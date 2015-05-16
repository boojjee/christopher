package helpers

import (
	"strings"
)

func ErrSplit(text string) string {
	s := strings.Split(text, ":")
	return s[0]
}

func ErrCode(text string) string {
	s := strings.Split(text, ":")
	s1 := strings.Split(s[0], " ")
	return s1[1]
}
