package files

import (
	mustache "github.com/Flaque/thaum/thaum/mustache"
	output "github.com/Flaque/thaum/thaum/output"
	util "github.com/Flaque/thaum/thaum/util"
	"github.com/spf13/afero"
	"os"
)

// A complete representation of a template that's
// stored in thaum_files.
type Template struct {
	Dirs      []string
	Files     []TemplateFile
	Variables map[string]string
	Name      string
}

// An individual file inside of the template
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

// Walks the path folder and returns all the files to compile
func getTemplateFromFiles(path string, templateName string) Template {
	var filesToCompile []TemplateFile
	var dirs []string
	nameSet := make(map[string]string)

	// Create Walk function
	walkFn := func(inputPath string, info os.FileInfo, err error) error {

		// Handle dirs
		if info.IsDir() {
			names := mustache.FindVariables(inputPath)
			nameSet = util.AddStringsToSet(names, nameSet)
			dirs = append(dirs, inputPath)
			return nil
		}

		// Ignore DS Stores
		if IsDsStore(info.Name()) {
			return nil
		}

		// Work with files
		names := getEmptyVarsFromFile(inputPath)
		nameSet = util.AddStringsToSet(names, nameSet)

		filesToCompile = append(filesToCompile,
			TemplateFile{templateName, inputPath, emptyStringMap(names)})
		return nil
	}

	// Actually walk through here.
	afero.Walk(AppFs, path, walkFn)

	return Template{dirs, filesToCompile, nameSet, templateName}
}

// From a template name, returns a complete template
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
	thaumPath, err := validThaumPath()
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
