package files

import (
	"fmt"
	constants "github.com/Flaque/thaum/thaum/constants"
	thaumErrors "github.com/Flaque/thaum/thaum/errors"
	"github.com/spf13/afero"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

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

// Strips a real path of style `<thaum_files>/<template>/blahblah`
// to just `blahblah`
func stripTemplatePrefix(template string, path string) string {
	splitPoint := fmt.Sprintf("%s/%s", constants.ThaumFiles, template)
	partialPath := strings.Split(path, splitPoint)[1]
	return strings.TrimPrefix(partialPath, "/")
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

// Create an map of strings with their values empty
func emptyStringMap(strs []string) map[string]string {
	vars := make(map[string]string)
	for _, name := range strs {
		vars[name] = ""
	}
	return vars
}

// Makes sure that there is a thaum_files somewhere
func validThaumPath() (string, error) {
	return existsAbove(cwd(), constants.ThaumFiles)
}

// Returns the available thaum templates
func ThaumTemplates() ([]string, error) {
	thaumPath, err := validThaumPath()
	if err != nil {
		return []string{}, err
	}

	files, err := afero.ReadDir(AppFs, thaumPath)
	if err != nil {
		return []string{}, err
	}

	var names []string
	for _, f := range files {
		if f.IsDir() {
			names = append(names, f.Name())
		}
	}

	return names, nil
}

// Ignores .DS_store
func IsDsStore(path string) bool {
	return filepath.Base(path) == ".DS_Store"
}

// Trims the .thaum extension from the path
func RemoveThaumExtension(path string) string {
	return strings.TrimSuffix(path, ".thaum")
}
