package main

import (
  "fmt"
  "os"
  "log"
  "github.com/spf13/afero"
)

// Global Afero Filesystem variable
var AppFs afero.Fs = afero.NewOsFs()

// Returns true if the path is a real file that exists
func exists(path string) (bool) {
  exists, err := afero.Exists(AppFs, path)
  if err != nil { log.Fatal(err) }
  return exists
}

// Returns the path of a template if that template exists
func findTemplate(template string) (string, error) {

  // Check if thaum_files folder exists
  path := fmt.Sprintf("./%s", THAUM_FILES)
  if !exists(path) {
    return "", ErrNoTemplateFolder
  }

  // Check if this template exists
  path = fmt.Sprintf("%s/%s", path, template)
  if !exists(path) {
    return "", ErrNoTemplate
  }

  return path, nil // Success!
}

// The [WalkFn](https://golang.org/pkg/path/filepath/#WalkFunc).
func walkFiles(path string, info os.FileInfo, err error) error {
  if err != nil { return err }

  fmt.Println(path)

  return nil
}
