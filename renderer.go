package echoviewrenderer

import (
	"io"
	"path/filepath"

	"github.com/labstack/echo"
)

// ViewRenderer renders views.
type ViewRenderer struct {
	views   map[string]view
	funcMap map[string]interface{}
}

// New returns a new ViewRenderer.
func New(pageTemplatesPath string, funcMap map[string]interface{}) (*ViewRenderer, error) {
	views := make(map[string]view)

	viewNames, err := getAllViewNames(pageTemplatesPath)
	if err != nil {
		return nil, err
	}

	for _, view := range viewNames {
		views[view] = newView(view, funcMap)
	}

	return &ViewRenderer{
		views:   views,
		funcMap: funcMap,
	}, nil
}

// Render renders a template document
func (tr *ViewRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return tr.views[name].templates.ExecuteTemplate(w, "base", data)
}

func getAllViewNames(pageTemplatesPath string) ([]string, error) {
	matches, err := filepath.Glob(pageTemplatesPath)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, match := range matches {
		base := filepath.Base(match)
		extensionLength := len(filepath.Ext(match))
		name := base[:len(base)-extensionLength]
		names = append(names, name)
	}

	return names, nil
}
