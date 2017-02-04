package constants

const ThaumFiles = "thaum_files"
const HelpTemplate = `
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

// Output statements
const NotifyThaumFiles  = "Using thaum_files at: "
const NotifyCreatedFile = "Created file: "
const NotifyCreatedDir  = "Created dir: "
