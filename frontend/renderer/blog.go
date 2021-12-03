package renderer

import (
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/content"
	routing "github.com/jackwhelpton/fasthttp-routing/v2"
)

// blog handler renders blog data
func (r *Renderer) blog(c *routing.Context) (err error) {

	blogContent := content.Blog{
	}

	return c.Write(blogContent)
}
