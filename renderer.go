package echoviewrenderer

import (
	"io"

	"github.com/labstack/echo"
)

// ViewRenderer renders views.
type ViewRenderer struct {
	views   map[string]view
	funcMap map[string]interface{}
}

// New returns a new ViewRenderer.
func New(viewNames []string, funcMap map[string]interface{}) *ViewRenderer {
	views := make(map[string]view)

	for _, view := range viewNames {
		views[view] = newView(view, funcMap)
	}

	return &ViewRenderer{
		views:   views,
		funcMap: funcMap,
	}
}

// Render renders a template document
func (tr *ViewRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return tr.views[name].templates.ExecuteTemplate(w, "base", data)
}
