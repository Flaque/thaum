package mustache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRender(t *testing.T) {
	template := "Hello {{firstName}} {{lastName}}!"
	variables := map[string]string{
		"firstName": "test",
		"lastName":  "mcTest",
	}
	assert.Equal(t, Render(template, variables), "Hello test mcTest!", "Rendered template does not equal expected values")
}

// Helper function for TestFindVariables
func contains(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func TestFindVariables(t *testing.T) {
	assert := assert.New(t)

	template := "Hello {{firstName}} {{lastName}}! I like you {{firstName}}"
	result := FindVariables(template)

	assert.Equal(len(result), 2, "Length of variables is not 2")
	assert.Contains(result, "firstName", "Variables do not contain 'firstName'")
	assert.Contains(result, "lastName", "Variables do not contain 'lastName'")
}
