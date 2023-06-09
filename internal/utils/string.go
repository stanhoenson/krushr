package utils

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/jinzhu/inflection"
)

func StringArrayIncludesSubstring(arr []string, substr string) bool {
	for _, s := range arr {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// Converts a reflect string to a more user-friendly type name.
func ReflectTypeToName(s string) string {
	// Get the last part of the type
	parts := strings.Split(s, ".")
	lastPart := parts[len(parts)-1]
	// Remove "[]" suffix if present
	if strings.HasSuffix(lastPart, "[]") {
		lastPart = strings.TrimSuffix(lastPart, "[]")
		// Pluralize the last part
		lastPart = Pluralize(lastPart)
	}
	// Capitalize first letter
	r, n := utf8.DecodeRuneInString(lastPart)
	return string(unicode.ToUpper(r)) + lastPart[n:]
}

// Pluralizes a word.
func Pluralize(s string) string {
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "x") {
		return s + "es"
	} else if strings.HasSuffix(s, "y") {
		return strings.TrimSuffix(s, "y") + "ies"
	} else {
		return s + "s"
	}
}

func GetTypeString(entity interface{}) string {
	entityType := reflect.TypeOf(entity).Elem()

	// check if it's a slice/array type
	if entityType.Kind() == reflect.Slice || entityType.Kind() == reflect.Array {
		entityType = entityType.Elem()
		return inflection.Plural(entityType.Name())
	}

	return entityType.Name()
}

func FindDuplicates(arr []string) []string {
	duplicates := []string{}
	frequency := make(map[string]int)

	// Count the frequency of each string
	for _, str := range arr {
		frequency[str]++
	}

	// Check for duplicates
	for str, count := range frequency {
		if count > 1 {
			duplicates = append(duplicates, str)
		}
	}

	return duplicates
}
