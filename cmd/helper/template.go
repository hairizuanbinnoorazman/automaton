package helper

import (
	"fmt"
	"io"
	"path"
	"text/template"
)

// RenderOutput renders the output in a markdown file
// This file will not be used effectively yet and this might involve rethinking certain aspects of the package design.
// This involves internal package changes
func RenderOutput(w io.Writer, templateFile string, a interface{}) error {
	_, templateFileValue := path.Split(templateFile)
	t := template.Must(template.New(templateFileValue).ParseFiles(templateFile))
	err := t.Execute(w, a)

	if err != nil {
		fmt.Println("Unable to render template")
		fmt.Println(err.Error())
		return err
	}
	// Add a new line
	io.WriteString(w, "\n\n")
	return nil
}
