package ghost

//go:generate $GOPATH/bin/easyjson -pkg -no_std_marshalers

// Client is the ghost backend client
type Client interface {
	GetPosts(params ...QueryParam) (posts *Posts, err error)
	GetPostBySlug(slug string, params ...QueryParam) (posts *Posts, err error)
	GetPageBySlug(slug string) (pages *Pages, err error)
}
