package ghost

import (
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost/data"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost/params"
)

// Client is the ghost backend client
type Client interface {
	// GetPosts returns blog posts according to query params
	GetPosts(queryParams ...params.Modifier) (posts *data.Posts, err error)
	// GetPostBySlug returns a single post by its slug title and query params
	GetPostBySlug(slug string, queryParams ...params.Modifier) (posts *data.Posts, err error)
	// GetPageBySlug returns a single page by its slug title and query params
	GetPageBySlug(slug string, queryParams ...params.Modifier) (pages *data.Pages, err error)
}
