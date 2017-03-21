package main

import (
	"bytes"
	filet "github.com/Flaque/filet"
	testUtil "github.com/Flaque/thaum/thaum/testingutil"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// Credit to https://gist.github.com/mindscratch/0faa78bd3c0005d080bf
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

var testRegistry []string

func TestThaum(t *testing.T) {
	defer filet.CleanUp(t)

	// Create thaum_files with myTemplate with filepath:
	// myLocation
	// 	- thaum_files
	// 		- myTemplate
	// 	- mySrc
	myLocation := filet.TmpDir(t, "")
	os.Chdir(myLocation) // Move the working directory to our temp location
	wd, _ := os.Getwd()

	mySrc := filet.TmpDir(t, wd)
	thaum_files := testUtil.TmpThaumFiles(t, wd)
	myTemplate := filet.TmpDir(t, thaum_files)
	// myFile      := testUtil.TmpFile(myTemplate)

	// Change directory back to where a user might use it.
	os.Chdir(mySrc)

	// Run Thaum
	app := buildApp()
	args := []string{"", filepath.Base(myTemplate)}
	text := captureStdout(func() {
		app.Run(args)
	})

	assert.Contains(t, text, "Using thaum_files", "thaum myTemplate does not contain thaum_files")
}
