package main

import "errors"

/**
 * Package Level Errors
 */
var (
  ErrNoTemplateFolder = errors.New("Thaum can't find your thaum_files!")
  ErrNoTemplate = errors.New("Thaum can't that template in your folder.")
  ErrNoOverwrite = errors.New(`Thaum found a file in your template with
    the same name as an existing file.`)
  ErrNoTemplateFolderAnywhere = errors.New(`Thaum can't find
    your thaum_files anywhere in the parent directories!`)
)
