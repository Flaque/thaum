package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func Test{{name}}(t *testing.T) {
	g := Goblin(t)
	g.Describe("{{name}} render()", func() {
		g.It("Should do something", func() {
			// Put your test here!
		})
	})
}
