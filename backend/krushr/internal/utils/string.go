package utils

import (
	"fmt"
	"strings"
)

func StringArrayIncludesSubstring(arr []string, substr string) bool {
	for _, s := range arr {
		fmt.Printf("%v, %v, %v", s, substr, strings.Contains(s, substr))
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}
