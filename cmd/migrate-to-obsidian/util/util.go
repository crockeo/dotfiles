package util

import "strings"

func EscapePath(title string) string {
	return strings.Replace(title, "/", "\\/", -1)
}
