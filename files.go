package main

import (
  "fmt"
  "os"
  "log"
  "path/filepath"
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

// Strips a real path of style `<THAUM_FILES>/<template>/blahblah`
// to just `blahblah`
func stripTemplatePrefix(template string, path string) (string) {
  prefix := fmt.Sprintf("%s/%s", THAUM_FILES, template)
  p, err := filepath.Rel(prefix, path)
  if err != nil { log.Fatal(err) }
  return p
}

// Creates a compiled file in the output
func createCompiledFile(inputPath string, outputPath string, name string) {
  AppFs.Create(outputPath)
  content := renderFile(inputPath, name)
  afero.WriteFile(AppFs, outputPath, []byte(content), 0755)
}

func compileTemplate(inputPath string, template string, name string) error {
  stat, _ := os.Stat(inputPath) // TODO Check error
  outputPath := stripTemplatePrefix(template, inputPath)

  // Skip root
  if outputPath == "." { return nil }

  if exists(outputPath) {
    log.Fatal(ErrNoOverwrite)
  }

  if stat.IsDir() {
    AppFs.Mkdir(outputPath, 0755)
  } else {
    createCompiledFile(outputPath, inputPath, name)
  }

  fmt.Printf("Created: %s\n", outputPath)

  return nil
}

// Compiles a template and moves it over
func compile(template string, name string) {

  // Find the path for the template; make sure template exists
  path, err := findTemplate(template)
  if err != nil { log.Fatal(err) }

  // Create Walk function
  walkFn := func(inputPath string, info os.FileInfo, err error) error {
    return compileTemplate(inputPath, template, name)
  }

  // Actually walk through here.
  afero.Walk(AppFs, path, walkFn)
}
