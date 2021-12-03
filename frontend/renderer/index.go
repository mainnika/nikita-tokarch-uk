package renderer

import (
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/content"
	routing "github.com/jackwhelpton/fasthttp-routing/v2"
)

// index handler renders index data
func (r *Renderer) index(c *routing.Context) (err error) {

	indexContent := content.Index{
	}

	return c.Write(indexContent)
}
