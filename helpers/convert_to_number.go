package helpers

import (
	"strconv"
)

func Convert_string_to_int(number string) int64 {
	numbr, _ := strconv.ParseInt(number, 0, 64)
	return numbr
}
func Convert_string_to_float(number string) float64 {
	numbr, _ := strconv.ParseFloat(number, 64)
	return numbr
}
