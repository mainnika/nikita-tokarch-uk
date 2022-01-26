package routes

import (
	"net/http"

	routing "github.com/jackwhelpton/fasthttp-routing/v2"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/content"
)

// slug renders page by its slug
func (r *Routes) slug(c *routing.Context) (err error) {

	pageSlug := c.Param("slug")
	if pageSlug == "" {
		return routing.NewHTTPError(http.StatusNotFound)
	}

	page, err := r.GhostClient.GetPageBySlug(pageSlug)
	if err != nil {
		return
	}

	pageContent := content.Blog{
		Meta:  page.Meta,
		Posts: page.Pages,
	}

	return c.Write(pageContent)
}
