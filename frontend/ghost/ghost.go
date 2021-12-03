package ghost

// Client is the ghost backend client
type Client interface {
	GetPosts(params ...QueryParam) (posts *Posts, err error)
	GetPostBySlug(slug string, params ...QueryParam) (posts *Posts, err error)
	GetPageBySlug(slug string) (pages *Pages, err error)
}
