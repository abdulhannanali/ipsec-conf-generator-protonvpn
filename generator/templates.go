package generator

import (
	"embed"
	"text/template"
)

//go:embed templates/*
var fs embed.FS

var connectionTemplate *template.Template
var connectionFileTemplate *template.Template

func init() {
	connectionFileTemplate = initTemplate("connection_file", "templates/connection_file.txt")
	connectionTemplate = initTemplate("connection", "templates/connection.txt")
}

func initTemplate(name string, path string) *template.Template {
	data, err := fs.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return template.Must(template.New(name).Parse(string(data)))
}
