package main

import "errors"

/**
 * Package Level Errors
 */
var (
  ErrNoTemplateFolder = errors.New("Thaum can't find your thaum_files!")
  ErrNoTemplate = errors.New("Thaum can't that template in your folder.")
)
