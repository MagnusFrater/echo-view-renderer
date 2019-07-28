package echoviewrenderer

import (
	"fmt"
	"html/template"
	"path/filepath"
)

// view defines a template and its base layout.
type view struct {
	templates *template.Template
}

func newView(tmpl string, funcMap map[string]interface{}) view {
	files := []string{getTemplatePath(tmpl)}
	files = append(files, generateSharedTemplatesPaths()...)

	t, err := template.New(tmpl).Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return view{
		templates: t,
	}
}

func getTemplatePath(tmpl string) string {
	return fmt.Sprintf("web/%s/%s.html", tmpl, tmpl)
}

func generateSharedTemplatesPaths() []string {
	matches, err := filepath.Glob("web/shared/*/*.html")
	if err != nil {
		panic(err)
	}

	return matches
}
