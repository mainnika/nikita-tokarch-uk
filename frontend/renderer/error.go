package renderer

import (
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/content"
	routing "github.com/jackwhelpton/fasthttp-routing/v2"
)

// errorNotFound renders http error-404 template
func (r *Renderer) errorNotFound(c *routing.Context) (err error) {

	errorContent := content.Error{Message: "not found"}

	return c.Write(errorContent)
}
