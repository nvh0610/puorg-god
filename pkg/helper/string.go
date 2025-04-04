package helper

import "strings"

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToArray(s string) []string {
	return strings.Split(s, ",")
}
