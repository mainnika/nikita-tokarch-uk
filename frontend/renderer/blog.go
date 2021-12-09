package renderer

import (
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/content"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/ghost"
	routing "github.com/jackwhelpton/fasthttp-routing/v2"
)

// blog handler renders blog data
func (r *Renderer) blog(c *routing.Context) (err error) {

	postsPerPage := r.ContentConfig.PostsPerPage
	currentPage := c.QueryArgs().GetUintOrZero("page")

	latestPosts, err := r.GhostClient.GetPosts(ghost.QueryLimit(postsPerPage), ghost.QueryPage(currentPage))
	if err != nil {
		return
	}

	blogContent := content.Blog{
		Meta:  latestPosts.Meta,
		Posts: latestPosts.Posts,
	}

	return c.Write(blogContent)
}
