package renderer

import (
	"net/http"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/content"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/ghost"
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

	pinnedPageSlug := r.ContentConfig.Pinned
	postsPerPage := r.ContentConfig.PostsPerPage

	pinnedPages, err := r.GhostClient.GetPageBySlug(pinnedPageSlug)
	if err != nil {
		return
	}

	latestPosts, err := r.GhostClient.GetPosts(ghost.QueryLimit(postsPerPage))
	if err != nil {
		return
	}

	indexContent := content.Index{
		Pinned: pinnedPages.Pages,
		Meta:   latestPosts.Meta,
		Posts:  latestPosts.Posts,
	}

	return c.Write(indexContent)
}
