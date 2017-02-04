package files

import (
	testUtil "github.com/Flaque/thaum/thaum/testingutil"
	. "github.com/franela/goblin"
	"testing"
)

func TestCreateCompiledFile(t *testing.T) {
	defer testUtil.RemoveAllTestFiles(t)

	g := Goblin(t)
	g.Describe("createCompiledFile()", func() {
		g.It("Should create a file", func() {
			inPath := testUtil.TmpFile(testUtil.TmpDir("")).Name()
			outPath := testUtil.TmpDir("") + "/test"

			createCompiledFile(TemplateFile{"", inPath, make(map[string]string)},
				outPath)

			g.Assert(exists(outPath)).Equal(true)
		})
	})
}

func TestCreateCompiledDir(t *testing.T) {
	defer testUtil.RemoveAllTestFiles(t)

	g := Goblin(t)

	g.Describe("createCompiledFile()", func() {
		g.It("Should create a file", func() {
			outPath := testUtil.TmpDir("") + "/test"

			createCompiledDir(outPath)

			g.Assert(exists(outPath)).Equal(true)
		})
	})
}
