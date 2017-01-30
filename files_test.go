package main

import (
	"testing"
	"github.com/spf13/afero"
	"fmt"
	. "github.com/franela/goblin"
)

// Keeps track of files that we've used so we can clean up.
var testRegistry []string

/*
Creates a tmp directory for us to use.
 */
func tmpDir() string {
	name, err := afero.TempDir(AppFs, "", "test")
	if err != nil {
		panic(fmt.Sprint("unable to work with test dir", err))
	}
	testRegistry = append(testRegistry, name)

	return name
}

/*
Creates a tmp file for us to use when testing
 */
func tmpFile() afero.File {
	file, err := afero.TempFile(AppFs, "", "test")
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

func TestExists(t *testing.T) {
	defer removeAllTestFiles(t)

	g := Goblin(t)
	g.Describe("exists", func() {
		g.It("should return true if the file exists", func() {
			tmp := tmpFile()
			fmt.Println(tmp.Name())
			g.Assert(exists(tmp.Name())).Equal(true)
		})

		g.It("should return false if the file doesn't exist", func() {
			g.Assert(exists("IDontExist.jpg")).Equal(false)
		})
	})
}
