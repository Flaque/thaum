package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "github.com/hoisie/mustache"
)

func onRun(c *cli.Context) error {
  template := c.Args().Get(0)
  name := c.Args().Get(1)
  fmt.Printf("%s, %s", template, name)
  return nil
}

func main() {
  app := cli.NewApp()
  app.Name = "thaum"
  app.Usage = "Generate micro-boilerplates"
  app.Action = onRun

  app.Run(os.Args)
}
