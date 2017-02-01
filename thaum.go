package main

import (
	"os"
	"fmt"

	"github.com/urfave/cli"
	output "github.com/Flaque/thaum/output"
	constants "github.com/Flaque/thaum/constants"
	files "github.com/Flaque/thaum/files"
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
		output.Error(fmt.Sprintf("Thaum requires a name for your %q template. Example: \n\n $ thaum <template> <name> ", template))
		return nil
	}

	files.Compile(template, name)

	return nil
}

func main() {
	cli.AppHelpTemplate = constants.HelpTemplate

	app := cli.NewApp()
	app.Name = "thaum"
	app.Usage = "Generate micro-boilerplates"
	app.Action = onRun
	app.Version = "0.2.0"

	app.Run(os.Args)
}
