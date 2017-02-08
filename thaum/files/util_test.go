package files

import (
	thaumErrors "github.com/Flaque/thaum/thaum/errors"
	testUtil "github.com/Flaque/thaum/thaum/testingutil"
	. "github.com/franela/goblin"
	"os"
	"path/filepath"
	"testing"
	filet "github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
)

// Tests the exists() function
func TestExists(t *testing.T) {
	defer filet.CleanUp(t)

	g := Goblin(t)
	g.Describe("exists()", func() {
		g.It("should return true if the file exists", func() {
			tmp := filet.TmpFile(t, "", "")
			g.Assert(exists(tmp.Name())).Equal(true)
		})

		g.It("should return false if the file doesn't exist", func() {
			g.Assert(exists("IDontExist.jpg")).Equal(false)
		})
	})
}

// Tests the existsAbove() function
func TestExistsAbove(t *testing.T) {
	defer filet.CleanUp(t)

	// Create our test environment
	outerDir := filet.TmpDir(t, "")       // Outermost directory
	queryDir := filet.TmpDir(t, outerDir) // What we're looking for
	subDir := filet.TmpDir(t, outerDir)   // Some random directory
	subsubDir := filet.TmpDir(t, subDir)  // Another random dir inside the subDir

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
	defer filet.CleanUp(t)

	g := Goblin(t)
	g.Describe("cwd()", func() {
		g.It("should return... something?", func() {
			g.Assert(len(cwd()) > 0).IsTrue()
		})
	})
}

func TestThaumTemplates(t *testing.T) {
	defer filet.CleanUp(t)

	g := Goblin(t)
	g.Describe("ThaumTemplates()", func() {
		g.It("correctly returns available templates", func() {
			mySrc, _, template := testUtil.TmpThaumEnvironment(t, "")

			os.Chdir(mySrc)
			ts, err := ThaumTemplates()
			g.Assert(err).Equal(nil)
			g.Assert(ts[0]).Equal(filepath.Base(template))
		})
	})
}

func TestIsDsStore(t *testing.T) {

	g := Goblin(t)
	g.Describe("IsDsStore()", func() {
		g.It("correctly identifies a DS Store filepath", func() {
			const testPath = "path/to/.DS_Store"
			g.Assert(IsDsStore(testPath)).Equal(true)
		})
	})
}

func TestRemoveThaumExtension(t *testing.T) {
	assert.Equal(t, RemoveThaumExtension("this.thaum"),
		"this", true)
}
