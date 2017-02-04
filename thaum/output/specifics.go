/**
 * This file is for specific messages to the user
 */

package output

import (
  "fmt"
  constants "github.com/Flaque/thaum/thaum/constants"
  padUtf8 "github.com/willf/pad/utf8"

)

func UsingThaumFilesAt(path string) {
  Search(fmt.Sprintf(constants.NotifyThaumFiles + "%q", path))
}

func CreatedFile(path string) {
  Write(constants.NotifyCreatedFile + path)
}

func CreatedDir(path string) {
  Write(constants.NotifyCreatedDir + path)
}

func VariableLabel(variable string) {
  fmt.Print(highlight(padUtf8.Left(variable, 12, " ") + ": "))
}
