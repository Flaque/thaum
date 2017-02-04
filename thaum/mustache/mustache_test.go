package mustache

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestRender(t *testing.T) {
	g := Goblin(t)
	g.Describe("Mustache render()", func() {
		g.It("Should correctly compile a template", func() {
			template := "Hello {{firstName}} {{lastName}}!"
			variables := map[string]string {
				"firstName" : "test",
				"lastName"  : "mcTest",
			}
			g.Assert(Render(template, variables)).Equal("Hello test mcTest!")
		})
	})
}

// Helper function for TestFindVariables
func contains(list []string, a string) bool{
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func TestFindVariables(t *testing.T) {
	g := Goblin(t)
	g.Describe("Mustache TestFindVariables()", func() {
		g.It("Should find all variables", func() {
			template := "Hello {{firstName}} {{lastName}}! I like you {{firstName}}"
			result := FindVariables(template)
			g.Assert(len(result)).Equal(2)
			g.Assert(contains(result, "firstName")).Equal(true)
			g.Assert(contains(result, "lastName")).Equal(true)
		})
	})
}
