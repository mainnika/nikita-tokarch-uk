package ghost

//go:generate $GOPATH/bin/easyjson -pkg -no_std_marshalers

// Client is the ghost backend client
type Client interface {
	// GetPosts returns blog posts according to query params
	GetPosts(queryParams ...Modifier) (posts *Posts, err error)
	// GetPostBySlug returns a single post by its slug title and query params
	GetPostBySlug(slug string, queryParams ...Modifier) (posts *Posts, err error)
	// GetPageBySlug returns a single page by its slug title and query params
	GetPageBySlug(slug string, queryParams ...Modifier) (pages *Pages, err error)
}
