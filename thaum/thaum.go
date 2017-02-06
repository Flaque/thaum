package main

import (
	"bufio"
	"os"
	"strings"

	constants "github.com/Flaque/thaum/thaum/constants"
	files "github.com/Flaque/thaum/thaum/files"
	output "github.com/Flaque/thaum/thaum/output"
	"github.com/urfave/cli"
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
		templateNames, err := files.ThaumTemplates()
		if err != nil {
			output.ErrorAsObject(err)
			os.Exit(1)
		}
		output.ListTemplates(templateNames)
		return nil
	}

	templateStruct, err := files.GetTemplate(template)
	if err != nil {
		output.ErrorAsObject(err)
		os.Exit(1)
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
	app.Version = "0.4.0"
	return app
}

func main() {
	app := buildApp()
	app.Run(os.Args)
}
