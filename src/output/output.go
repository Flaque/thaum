package output

import (
  "fmt"
  "github.com/kortschak/ct"
  padUtf8 "github.com/willf/pad/utf8"
)

var (
  warn = (ct.Fg(ct.Red)).Paint
  highlight = (ct.Fg(ct.Blue) | ct.Bold).Paint
)

func Error(text string) {
  fmt.Println(warn(padUtf8.Right("🚨  Error:", 10, " ")), text)
}

func ErrorAsObject(err error) {
  Error(fmt.Sprintf("%v", err))
}

func Search(text string) {
  fmt.Println("🔍  " + text)
}

func Write(text string) {
  fmt.Println("✍️  " + text)
}
