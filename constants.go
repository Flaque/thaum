package main

const THAUM_FILES = "thaum_files"
const HELP_TEMPLATE = `
NAME:
  {{.Name}} - {{.Usage}}

NOTES:
  Thaum requires a "thaum_files" in your project. It will
  continue to look up for a template in the parent directories
  until it finds a sibling.

USAGE:
  thaum <template> <name>

EXAMPLE:
  thaum component MyWindow

VERSION:
  {{.Version}}

`
