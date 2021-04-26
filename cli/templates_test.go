package cli

import (
	"strings"
	"testing"
)

func TestHelpTemplate(t *testing.T) {
	var builder strings.Builder
	err := helpTemplate.Execute(&builder, helpTemplateParams{GithubRepoLink: "https://github.com/x/xyz"})

	if err != nil {
		t.Errorf("Didn't expect error:\n%v", err)
	}

	if len(builder.String()) < 1 {
		t.Errorf("template is not producing any output")
	}
}
