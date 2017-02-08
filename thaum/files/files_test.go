package files

import (
	. "github.com/franela/goblin"
	"testing"
  filet "github.com/flaque/filet"
)

func TestCreateCompiledFile(t *testing.T) {
	defer filet.CleanUp(t)

	g := Goblin(t)
	g.Describe("createCompiledFile()", func() {
		g.It("Should create a file", func() {
			inPath := filet.TmpFile(t, filet.TmpDir(t, ""), "").Name()
			outPath := filet.TmpDir(t, "") + "/test"

			createCompiledFile(TemplateFile{"", inPath, make(map[string]string)},
				outPath)

			g.Assert(exists(outPath)).Equal(true)
		})
	})
}

func TestCreateCompiledDir(t *testing.T) {
	defer filet.CleanUp(t)

	g := Goblin(t)

	g.Describe("createCompiledFile()", func() {
		g.It("Should create a file", func() {
			outPath := filet.TmpDir(t, "") + "/test"

			createCompiledDir(outPath)

			g.Assert(exists(outPath)).Equal(true)
		})
	})
}
