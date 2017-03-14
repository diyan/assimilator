package renderer

import (
	"io"

	rice "github.com/GeertJohan/go.rice"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

func init() {
	RegisterTags()
	RegisterFilters()
}

func New(box *rice.Box) *Renderer {
	templateLoader := riceTemplateLoader{
		templateBox: box,
	}
	return &Renderer{
		templateSet: pongo2.NewSet("assimilator", templateLoader),
	}
}

type riceTemplateLoader struct {
	templateBox *rice.Box
}

// Abs calculates the path to a given template. Whenever a path must be resolved
// due to an import from another template, the base equals the parent template's path.
func (loader riceTemplateLoader) Abs(base, name string) string {
	// TODO make sure we can ignore `base` argument
	return name
}

// Get returns an io.Reader where the template's content can be read from.
func (loader riceTemplateLoader) Get(path string) (io.Reader, error) {
	return loader.templateBox.Open(path)
}

type Renderer struct {
	templateSet *pongo2.TemplateSet
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Re-read template from file on each request during development
	// TODO use cache/no-cache via settings
	tpl, err := r.templateSet.FromFile(name)
	if err != nil {
		return err
	}
	return tpl.ExecuteWriter(getPongoContext(data), w)
}

func getPongoContext(data interface{}) pongo2.Context {
	if context, ok := data.(map[string]interface{}); ok {
		return context
	}
	return nil
}
