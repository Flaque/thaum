package main

import (
	"fmt"
	"github.com/spf13/afero"
	"log"
	"os"
	"path"
	"strings"
	mustache "github.com/Flaque/thaum/mustache"
	output "github.com/Flaque/thaum/output"
	thaumErrors "github.com/Flaque/thaum/errors"
	constants "github.com/Flaque/thaum/constants"
)

// Global Afero Filesystem variable
var AppFs afero.Fs = afero.NewOsFs()

// Returns true if the path is a real file that exists
func exists(path string) bool {
	exists, err := afero.Exists(AppFs, path)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}

// Returns the path if a query exists above the "from" directory
func existsAbove(from string, query string) (string, error) {

	myPath := path.Join(from, query)
	if exists(myPath) {
		return myPath, nil
	}

	// We've hit root!
	if path.Join(from, "../") == from {
		return "", thaumErrors.NoTemplateFolderAnywhere
	}

	// Search in parent directory
	return existsAbove(path.Join(from, "../"), query)
}

// Returns the current working directory
func cwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

// Returns the path of a template if that template exists
func findTemplate(template string, thaumPath string) (string, error) {

	// Check if thaum_files folder exists
	path := thaumPath
	if !exists(path) {
		return "", thaumErrors.NoTemplateFolder
	}

	// Check if this template exists
	path = fmt.Sprintf("%s/%s", path, template)
	if !exists(path) {
		return "", thaumErrors.NoTemplate
	}

	return path, nil // Success!
}

// Strips a real path of style `<thaum_files>/<template>/blahblah`
// to just `blahblah`
func stripTemplatePrefix(template string, path string) string {
	splitPoint := fmt.Sprintf("%s/%s", constants.ThaumFiles, template)
	partialPath := strings.Split(path, splitPoint)[1]
	return strings.TrimPrefix(partialPath, "/")
}

// Creates a compiled file in the output
func createCompiledFile(inputPath string, outputPath string, name string) {
	file, err := AppFs.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	content := mustache.RenderFile(inputPath, name)
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	output.Write("Created file: " + outputPath)
}

func compileTemplate(inputPath string, template string, name string) error {
	stat, err := os.Stat(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	outputPath := stripTemplatePrefix(template, inputPath)
	outputPath = mustache.Render(outputPath, name) // Compile any {{}}s in paths

	// Skip root
	if outputPath == "." || outputPath == "" {
		return nil
	}

	if exists(outputPath) {
		output.ErrorAsObject(thaumErrors.NoOverwrite)
		os.Exit(1)
	}

	if stat.IsDir() {
		err := AppFs.Mkdir(outputPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		output.Write("Created folder: " + outputPath)
	} else {
		createCompiledFile(inputPath, outputPath, name)
	}

	return nil
}

// Compiles a template and moves it over
func compile(template string, name string) {

	thaumPath, err := existsAbove(cwd(), constants.ThaumFiles)
	if err != nil {
		output.ErrorAsObject(err)
		return
	}

	output.Search(fmt.Sprintf("Using thaum_files at: %q", thaumPath))

	// Find the path for the template; make sure template exists
	path, err := findTemplate(template, thaumPath)

	if err != nil {
		output.ErrorAsObject(err)
		return
	}

	// Create Walk function
	walkFn := func(inputPath string, info os.FileInfo, err error) error {
		return compileTemplate(inputPath, template, name)
	}

	// Actually walk through here.
	afero.Walk(AppFs, path, walkFn)
}
