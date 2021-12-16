package ghost

//go:generate $GOPATH/bin/easyjson -pkg -no_std_marshalers

// Client is the ghost backend client
type Client interface {
	GetPosts(queryParams ...Modifier) (posts *Posts, err error)
	GetPostBySlug(slug string, queryParams ...Modifier) (posts *Posts, err error)
	GetPageBySlug(slug string, queryParams ...Modifier) (pages *Pages, err error)
}
