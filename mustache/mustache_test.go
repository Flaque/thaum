package mustache

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestRender(t *testing.T) {
	g := Goblin(t)
	g.Describe("Mustache render()", func() {
		g.It("Should correctly compile a template", func() {
			template := "Hello {{name}}!"
			g.Assert(render(template, "test")).Equal("Hello test!")
		})
	})
}
