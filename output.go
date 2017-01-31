package main

import (
  "fmt"
  "github.com/kortschak/ct"
  padUtf8 "github.com/willf/pad/utf8"
)

var (
  warn = (ct.Fg(ct.Red)).Paint
  highlight = (ct.Fg(ct.Blue) | ct.Bold).Paint
)

func ErrorLog(text string) {
  fmt.Println(warn(padUtf8.Right("🚨  Error:", 10, " ")), text)
}

func ErrorAsObjectLog(err error) {
  ErrorLog(fmt.Sprintf("%v", err))
}

func SearchLog(text string) {
  fmt.Println("🔍  " + text)
}

func WriteLog(text string) {
  fmt.Println("✍️  " + text)
}
