package main

import (
	"bytes"
	testUtil "github.com/Flaque/thaum/thaum/testingutil"
	. "github.com/franela/goblin"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
  filet "github.com/Flaque/filet"
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

	g := Goblin(t)
	g.Describe("Running thaum as ", func() {

		g.It("thaum myTemplate will mention thaum_files", func() {

			// Change directory back to where a user might use it.
			os.Chdir(mySrc)

			// Run Thaum
			app := buildApp()
			args := []string{"", filepath.Base(myTemplate)}
			text := captureStdout(func() {
				app.Run(args)
			})

			// Test that we mentioned which thaum files are being used
			g.Assert(strings.Contains(text, "Using thaum_files")).Equal(true)
		})

	})
}
