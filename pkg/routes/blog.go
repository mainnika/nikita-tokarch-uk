package routes

import (
	routing "github.com/jackwhelpton/fasthttp-routing/v2"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/content"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost/params"
)

// blog handler renders blog data
func (r *Routes) blog(c *routing.Context) (err error) {

	postsPerPage := r.ContentConfig.PostsPerPage
	currentPage := c.QueryArgs().GetUintOrZero("page")

	latestPosts, err := r.GhostClient.GetPosts(
		params.WithLimit(postsPerPage),
		params.WithPage(currentPage),
	)
	if err != nil {
		return
	}

	blogContent := content.Blog{
		Meta:  latestPosts.Meta,
		Posts: latestPosts.Posts,
	}

	return c.Write(blogContent)
}
