package files

import (
	"fmt"
	"github.com/spf13/afero"
	"log"
	"os"
	mustache "github.com/Flaque/thaum/thaum/mustache"
	output "github.com/Flaque/thaum/thaum/output"
	thaumErrors "github.com/Flaque/thaum/thaum/errors"
	util "github.com/Flaque/thaum/thaum/util"
)

// Global Afero Filesystem variable
var AppFs afero.Fs = afero.NewOsFs()

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

	if exists(outputPath) {
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

// Returns a variable map with empty items for each name in the path
func getEmptyVarsFromFile(path string) []string {
	bytes, err := afero.ReadFile(AppFs, path)
	if err != nil {
		log.Fatal(err)
	}

	return mustache.FindVariables(string(bytes))
}

// Compiles a template and moves it over
func Compile(template Template) {

	for _, d := range template.Dirs {
		createCompiledDir(createOutputPath(d, template.Name, template.Variables))
	}

	for _, f := range template.Files {
		compileTemplateFile(f)
	}
}
