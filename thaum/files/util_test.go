package files

import (
	thaumErrors "github.com/Flaque/thaum/thaum/errors"
	testUtil "github.com/Flaque/thaum/thaum/testingutil"
	. "github.com/franela/goblin"
	"path/filepath"
	"testing"
	"os"
)

// Tests the exists() function
func TestExists(t *testing.T) {
	defer testUtil.RemoveAllTestFiles(t)

	g := Goblin(t)
	g.Describe("exists()", func() {
		g.It("should return true if the file exists", func() {
			tmp := testUtil.TmpFile("")
			g.Assert(exists(tmp.Name())).Equal(true)
		})

		g.It("should return false if the file doesn't exist", func() {
			g.Assert(exists("IDontExist.jpg")).Equal(false)
		})
	})
}

// Tests the existsAbove() function
func TestExistsAbove(t *testing.T) {
	defer testUtil.RemoveAllTestFiles(t)

	// Create our test environment
	outerDir := testUtil.TmpDir("")       // Outermost directory
	queryDir := testUtil.TmpDir(outerDir) // What we're looking for
	subDir := testUtil.TmpDir(outerDir)   // Some random directory
	subsubDir := testUtil.TmpDir(subDir)  // Another random dir inside the subDir

	// Run tests
	g := Goblin(t)
	g.Describe("existsAbove()", func() {
		g.It("should return true if the folder exists above", func() {
			query := filepath.Base(queryDir)
			result, err := existsAbove(subsubDir, query)
			if err != nil {
				g.Fail(err)
			}

			g.Assert(result).Equal(queryDir)
		})

		g.It("should return err if the folder does not exist above", func() {
			query := "IDontExistTrustMe"
			_, err := existsAbove(subsubDir, query)
			g.Assert(err).Equal(thaumErrors.NoTemplateFolderAnywhere)
		})
	})
}

// More or less tests cwd()
func TestCwd(t *testing.T) {
	defer testUtil.RemoveAllTestFiles(t)

	g := Goblin(t)
	g.Describe("cwd()", func() {
		g.It("should return... something?", func() {
			g.Assert(len(cwd()) > 0).IsTrue()
		})
	})
}

func TestThaumTemplates(t *testing.T) {
	defer testUtil.RemoveAllTestFiles(t)

	g := Goblin(t)
	g.Describe("ThaumTemplates()", func() {
		g.It("correctly returns available templates", func() {
			mySrc, _, template := testUtil.TmpThaumEnvironment("")

			os.Chdir(mySrc)
			ts, err := ThaumTemplates()
			g.Assert(err).Equal(nil)
			g.Assert(ts[0]).Equal(filepath.Base(template))
		})
	})
}
