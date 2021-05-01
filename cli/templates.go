package cli

import (
	_ "embed"
	"text/template"
)

//go:embed templates/help.txt
var helpRaw string
var helpTemplate *template.Template

type helpTemplateParams struct {
	GithubRepoLink string
}

func init() {
	helpTemplate = template.Must(template.New("help").Parse(helpRaw))
}
