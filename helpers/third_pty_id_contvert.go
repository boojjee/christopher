package helpers

import (
	// "fmt"
	"strings"
)

func Substr_thirdid(tid string, source string) string {
	switch source {
	case "RunKeeper":
		// s := strings.Split("/fitnessActivities/x1/xxxx", "/")
		// fmt.Println(s[2])
		s := strings.Split(tid, "/")
		return s[2]
	case "Nike Plus":
		return tid
	case "Strava":
		return tid
	default:
		return "none"
	}
}
