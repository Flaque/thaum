package main

import (
	"os"

	"github.com/urfave/cli"
)

// Called when thaum is actually run.
func onRun(c *cli.Context) error {
	template := c.Args().Get(0)
	name := c.Args().Get(1)

	compile(template, name)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "thaum"
	app.Usage = "Generate micro-boilerplates"
	app.Action = onRun

	app.Run(os.Args)
}
