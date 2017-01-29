package main

import (
  "fmt"
  "os"
  "log"

  "github.com/urfave/cli"
)

// Called when thaum is actually run.
func onRun(c *cli.Context) error {
  template := c.Args().Get(0)
  // name := c.Args().Get(1)

  path, err := findTemplate(template)
  if err != nil { log.Fatal(err) }
  fmt.Println(path)

  return nil
}

func main() {
  app := cli.NewApp()
  app.Name = "thaum"
  app.Usage = "Generate micro-boilerplates"
  app.Action = onRun

  app.Run(os.Args)
}
