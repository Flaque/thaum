package main

import (
	"os"
	"fmt"

	"github.com/urfave/cli"
)

// Called when thaum is actually run.
func onRun(c *cli.Context) error {
	template := c.Args().Get(0)
	name := c.Args().Get(1)

	if len(c.Args()) == 0 {
		cli.ShowAppHelp(c)
		return nil
	}

	if name == "" {
		ErrorLog(fmt.Sprintf("Thaum requires a name for your %q template. Example: \n\n $ thaum <template> <name> ", template))
		return nil
	}

	compile(template, name)

	return nil
}

func main() {
	cli.AppHelpTemplate = HELP_TEMPLATE

	app := cli.NewApp()
	app.Name = "thaum"
	app.Usage = "Generate micro-boilerplates"
	app.Action = onRun
	app.Version = "0.2.0"

	app.Run(os.Args)
}
