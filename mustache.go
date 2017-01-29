package main

import "github.com/hoisie/mustache"

// Renders a mustache template string with the "name" parameter.
func render(template string, name string) (string) {
  return mustache.Render(template, map[string]string{"name":name})
}
