package renderer

import (
	"net/http"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/content"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/templates"
	routing "github.com/jackwhelpton/fasthttp-routing/v2"
)

// rootRedirect redirects the root url to the index using http redirect
func (r *Renderer) rootRedirect(c *routing.Context) (err error) {

	c.Redirect(templates.URLIndex, http.StatusFound)

	return
}

// index handler renders index data
func (r *Renderer) index(c *routing.Context) (err error) {

	indexContent := content.Index{
	}

	return c.Write(indexContent)
}
