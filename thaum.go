package main

import (
  "fmt"
  "os"
  "errors"
  "log"

  "github.com/urfave/cli"
  "github.com/hoisie/mustache"
  "github.com/spf13/afero"
)

const THAUM_FILES = "thaum_files"

var (
  ErrNoTemplateFolder = errors.New("Thaum can't find your thaum_files!")
  ErrNoTemplate = errors.New("Thaum can't that template in your folder.")
)

// Global Afero Filesystem variable
var AppFs afero.Fs = afero.NewOsFs()

func exists(path string) (bool) {
  exists, err := afero.Exists(AppFs, path)
  if err != nil { log.Fatal(err) }
  return exists
}

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

func render(template string, name string) (string) {
  return mustache.Render(template, map[string]string{"name":name})
}

func onRun(c *cli.Context) error {
  template := c.Args().Get(0)
  // name := c.Args().Get(1)

  path, err := findTemplate(template)
  if err != nil { log.Fatal(err) }

  fmt.Printf("%s", path)
  return nil
}

func main() {
  app := cli.NewApp()
  app.Name = "thaum"
  app.Usage = "Generate micro-boilerplates"
  app.Action = onRun

  app.Run(os.Args)
}
