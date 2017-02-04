package files

import (
	constants "github.com/Flaque/thaum/thaum/constants"
	output "github.com/Flaque/thaum/thaum/output"
)

type Template struct {
	Dirs      []string
	Files     []TemplateFile
	Variables map[string]string
	Name      string
}

type TemplateFile struct {
	name      string
	inputPath string
	variables map[string]string
}

// Updates the TemplateFiles inside the template
func (t Template) Update() Template {
	for i, f := range t.Files {
		for name := range t.Variables {
			f.variables[name] = t.Variables[name] // Set value
			t.Files[i] = f
		}
	}
	return t
}

// From a template name, gets a template
func GetTemplate(template string) (Template, error) {
	// Make sure we can actually compile this template
	path, err := validateTemplatePath(template)
	if err != nil {
		return Template{}, err
	}

	return getTemplateFromFiles(path, template), nil
}

// Returns the path if this template exists
func validateTemplatePath(template string) (string, error) {

	// Make sure the thaum_files exists
	thaumPath, err := existsAbove(cwd(), constants.ThaumFiles)
	if err != nil {
		return "", err
	}

	// Find the path for the template; make sure template exists
	path, err := findTemplate(template, thaumPath)
	if err != nil {
		return "", err
	}

	// Path is good, tell the user and return path
	output.UsingThaumFilesAt(thaumPath)
	return path, nil
}
