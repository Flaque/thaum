package files

import (
	"fmt"
	thaumErrors "github.com/Flaque/thaum/thaum/errors"
	mustache "github.com/Flaque/thaum/thaum/mustache"
	output "github.com/Flaque/thaum/thaum/output"
	"github.com/spf13/afero"
	"log"
	"os"
)

// Global Afero Filesystem variable
var AppFs afero.Fs = afero.NewOsFs()
var overwrite bool = false

// Compile a template file
func compileTemplateFile(f TemplateFile) {
	stat, err := os.Stat(f.inputPath)
	if err != nil {
		log.Fatal(err)
	}

	outputPath := stripTemplatePrefix(f.name, f.inputPath)
	outputPath = mustache.Render(outputPath, f.variables)

	// Skip root
	if outputPath == "." || outputPath == "" {
		return
	}

	if exists(outputPath) && overwrite {
		output.ErrorAsObject(thaumErrors.NoOverwrite)
		os.Exit(1)
	}

	if stat.IsDir() {
		createCompiledDir(outputPath)
	} else {
		createCompiledFile(f, outputPath)
	}
}

// Creates a compiled file in the output
func createCompiledFile(f TemplateFile, outputPath string) {
	file, err := AppFs.Create(outputPath)
	if err != nil {
		fmt.Println("createCompiledFile Failed to create file: " + outputPath)
		os.Exit(1)
	}

	content := mustache.RenderFile(f.inputPath, f.variables)
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	output.CreatedFile(outputPath)
}

// Creates a compiled directory in the output
func createCompiledDir(outputPath string) {
	if outputPath == "" {
		return
	}

	err := AppFs.Mkdir(outputPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	output.CreatedDir(outputPath)
}

// Creates a compiled output path.
func createOutputPath(templateName string, path string, variables map[string]string) string {
	outputPath := stripTemplatePrefix(path, templateName)
	return mustache.Render(outputPath, variables)
}

// Returns a variable map with empty items for each name in the path
func getEmptyVarsFromFile(path string) []string {
	bytes, err := afero.ReadFile(AppFs, path)
	if err != nil {
		log.Fatal(err)
	}

	return mustache.FindVariables(string(bytes))
}

// Compiles a template and moves it over
func Compile(template Template, overwrite bool) {
	overwrite = overwrite

	for _, d := range template.Dirs {
		createCompiledDir(createOutputPath(d, template.Name, template.Variables))
	}

	for _, f := range template.Files {
		compileTemplateFile(f)
	}
}
