package main

import (
	"os"
	"bufio"
	"strings"

	"github.com/urfave/cli"
	constants "github.com/Flaque/thaum/thaum/constants"
	files "github.com/Flaque/thaum/thaum/files"
	output "github.com/Flaque/thaum/thaum/output"
)

func askForVariables(template files.Template) files.Template {
	reader := bufio.NewReader(os.Stdin)

	output.Space()
	for v := range template.Variables {
		output.VariableLabel(v)
		value, _ := reader.ReadString('\n')
		template.Variables[v] = strings.TrimSpace(value)
	}
	output.Space()

	return template.Update()
}

// Called when thaum is actually run.
func onRun(c *cli.Context) error {
	template := c.Args().Get(0)

	if len(c.Args()) == 0 {
		cli.ShowAppHelp(c)
		return nil
	}

	templateStruct, err := files.GetTemplate(template)
	if err != nil {
		output.ErrorAsObject(err)
	}

	templateStruct = askForVariables(templateStruct)
	files.Compile(templateStruct)

	return nil
}

func buildApp() *cli.App {
	cli.AppHelpTemplate = constants.HelpTemplate

	app := cli.NewApp()
	app.Name = "thaum"
	app.Usage = "Generate micro-boilerplates"
	app.Action = onRun
	app.Version = "0.2.0"
	return app
}

func main() {
	app := buildApp()
	app.Run(os.Args)
}
