package mustache

import (
	"github.com/hoisie/mustache"
	"regexp"
	"strings"
)

// Renders a mustache template string with the "name" parameter.
func Render(template string, variables map[string]string) string {
	return mustache.Render(template, variables)
}

// Renders a mustache template file with the "name" parameter.
func RenderFile(path string, variables map[string]string) string {
	return mustache.RenderFile(path, variables)
}

// Helper for keys
func getKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Returns all the variable names inside of the content
func FindVariables(content string) []string {
	variables := make(map[string]bool)

	re := regexp.MustCompile("({{)[A-Za-z]*(}})") // matches {{blah}} types
	matches := re.FindAllString(content, -1)
	for _, match := range matches {
		match := strings.TrimPrefix(match, "{{")
		match = strings.TrimSuffix(match, "}}")
		variables[match] = true
	}
	return getKeys(variables)
}
