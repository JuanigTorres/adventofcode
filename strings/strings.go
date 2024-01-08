package strings

import (
	"strings"
)

func Lines(str string) []string {
	return strings.Split(str, "\n")
}

func DeleteSubstring(str, substr string) string {
	return strings.Replace(str, substr, "", 1)
}

func Words(str string) []string {
	values := strings.Split(str, " ")

	r := make([]string, 0, len(values))
	for _, v := range values {
		if v != "" {
			r = append(r, strings.TrimSpace(v))
		}
	}
	return r
}
