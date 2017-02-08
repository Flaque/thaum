package testingutil

import (
  "fmt"
  "testing"
  "github.com/spf13/afero"
  filet "github.com/flaque/filet"
)

var AppFs afero.Fs = afero.NewOsFs()

/*
Creates a temporary `thaum_files` directory.
*/
func TmpThaumFiles(t *testing.T, dir string) string {
  name := dir + "/thaum_files"
  err := AppFs.Mkdir(name, 0755)
  if err != nil {
    panic(fmt.Sprint("unable to work with test dir", err))
  }
  filet.Files = append(filet.Files, name)

  return name
}

/*
Creates a temporary thaum-test environment for us
*/
func TmpThaumEnvironment(t *testing.T, dir string) (string, string, string) {
  myLocation := filet.TmpDir(t, dir)
  mySrc := filet.TmpDir(t, myLocation)
  thaum_files := TmpThaumFiles(t, myLocation)
  myTemplate := filet.TmpDir(t, thaum_files)
  filet.TmpFile(t, myTemplate, "")
  return mySrc, thaum_files, myTemplate
}
