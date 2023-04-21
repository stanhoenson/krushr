package utils

import (
	"strings"
)

func StringArrayIncludesSubstring(arr []string, substr string) bool {
	for _, s := range arr {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}
