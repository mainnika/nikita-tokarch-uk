package routes

import (
	"net/http"

	routing "github.com/jackwhelpton/fasthttp-routing/v2"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/content"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost/params"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/templates"
)

// rootRedirect redirects the root url to the index using http redirect
func (r *Routes) rootRedirect(c *routing.Context) (err error) {

	c.Redirect(templates.URLIndex, http.StatusFound)

	return
}

// index handler renders index data
func (r *Routes) index(c *routing.Context) (err error) {

	pinnedPageSlug := r.ContentConfig.Pinned
	postsPerPage := r.ContentConfig.PostsPerPage

	pinnedPages, err := r.GhostClient.GetPageBySlug(pinnedPageSlug)
	if err != nil {
		return
	}

	latestPosts, err := r.GhostClient.GetPosts(params.WithLimit(postsPerPage))
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
