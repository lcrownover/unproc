package util

import (
	"strings"
)

func removeEmptyElements(s []string) []string {
	var new_s []string
	for _, e := range s {
		es := strings.TrimSpace(e)
		if es != "" {
			new_s = append(new_s, es)
		}
	}
	return new_s
}

func CountLines(input string) int32 {
	lines := strings.Split(input, "\n")
	lines = removeEmptyElements(lines)
	lineCount := int32(len(lines))
	return lineCount
}
