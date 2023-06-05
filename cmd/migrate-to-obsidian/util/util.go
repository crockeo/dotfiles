package util

import (
	"regexp"
	"strings"
)

func EscapePath(title string) string {
	return strings.Replace(title, "/", "\\/", -1)
}

func Match(re *regexp.Regexp, s string) (map[string]string, bool) {
	match := re.FindStringSubmatch(s)
	if match == nil {
		return nil, false
	}
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		result[name] = match[i]
	}
	return result, true
}
