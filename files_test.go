package main

import (
	"testing"
	"github.com/spf13/afero"
	"path/filepath"
	"fmt"
	. "github.com/franela/goblin"
	thaumErrors "github.com/Flaque/thaum/errors"
)

// Keeps track of files that we've used so we can clean up.
var testRegistry []string

/*
Creates a tmp directory for us to use.
 */
func tmpDir(dir string) string {
	name, err := afero.TempDir(AppFs, dir, "dir")
	if err != nil {
		panic(fmt.Sprint("unable to work with test dir", err))
	}
	testRegistry = append(testRegistry, name)

	return name
}

/*
Creates a tmp file for us to use when testing
 */
func tmpFile(dir string) afero.File {
	file, err := afero.TempFile(AppFs, dir, "file")
	if err != nil {
		panic(fmt.Sprint("unable to work with tmp dir", err))
	}

	testRegistry = append(testRegistry, file.Name())

	return file
}

/*
Removes all files in our test registry
 */
func removeAllTestFiles(t *testing.T) {
	for _, path := range testRegistry {
		if err := AppFs.RemoveAll(path); err != nil {
			t.Error(AppFs.Name(), err)
		}
	}

	testRegistry = make([]string, 0)
}

// Tests the exists() function
func TestExists(t *testing.T) {
	defer removeAllTestFiles(t)

	g := Goblin(t)
	g.Describe("exists()", func() {
		g.It("should return true if the file exists", func() {
			tmp := tmpFile("")
			g.Assert(exists(tmp.Name())).Equal(true)
		})

		g.It("should return false if the file doesn't exist", func() {
			g.Assert(exists("IDontExist.jpg")).Equal(false)
		})
	})
}

// Tests the existsAbove() function
func TestExistsAbove(t *testing.T) {
	defer removeAllTestFiles(t)

	// Create our test environment
	outerDir  := tmpDir("")       // Outermost directory
	queryDir  := tmpDir(outerDir) // What we're looking for
	subDir    := tmpDir(outerDir) // Some random directory
	subsubDir := tmpDir(subDir)   // Another random dir inside the subDir

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
	defer removeAllTestFiles(t)

	g := Goblin(t)
	g.Describe("cwd()", func() {
		g.It("should return... something?", func() {
			g.Assert(len(cwd()) > 0).IsTrue()
		})
	})
}
