package errors

import "errors"

/**
 * Package Level Errors
 */
var (
	NoTemplateFolder = errors.New("Thaum can't find your thaum_files!")
	NoTemplate       = errors.New("Thaum can't find that template in your folder.")
	NoOverwrite      = errors.New(`Thaum found a file in your template with the same name as an existing file.`)
	NoTemplateFolderAnywhere = errors.New(`Thaum can't find your thaum_files anywhere in the parent directories!`)
)
