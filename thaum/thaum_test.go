package main

import (
	. "github.com/franela/goblin"
	"testing"
	"path/filepath"
	"os"
	"io"
	"bytes"
	"strings"
	testUtil "github.com/Flaque/thaum/thaum/testingutil"

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
	defer testUtil.RemoveAllTestFiles(t)

	// Create thaum_files with myTemplate with filepath:
	// myLocation
	// 	- thaum_files
	// 		- myTemplate
	// 	- mySrc
	myLocation  := testUtil.TmpDir("")
	os.Chdir(myLocation) // Move the working directory to our temp location
	wd, _ := os.Getwd()

	mySrc       := testUtil.TmpDir(wd)
	thaum_files := testUtil.TmpThaumFiles(wd)
	myTemplate  := testUtil.TmpDir(thaum_files)
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
