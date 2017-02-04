package testingutil

import (
	"fmt"
	"github.com/spf13/afero"
	"testing"
)

// Keeps track of files that we've used so we can clean up.
var testRegistry []string
var AppFs afero.Fs = afero.NewOsFs()

/*
Creates a tmp directory for us to use.
*/
func TmpDir(dir string) string {
	name, err := afero.TempDir(AppFs, dir, "dir")
	if err != nil {
		panic(fmt.Sprint("unable to work with test dir", err))
	}
	testRegistry = append(testRegistry, name)

	return name
}

func TmpThaumFiles(dir string) string {
	name := dir + "/thaum_files"
	err := AppFs.Mkdir(name, 0755)
	if err != nil {
		panic(fmt.Sprint("unable to work with test dir", err))
	}
	testRegistry = append(testRegistry, name)

	return name
}

/*
Creates a tmp file for us to use when testing
*/
func TmpFile(dir string) afero.File {
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
func RemoveAllTestFiles(t *testing.T) {
	for _, path := range testRegistry {
		if err := AppFs.RemoveAll(path); err != nil {
			t.Error(AppFs.Name(), err)
		}
	}

	testRegistry = make([]string, 0)
}
