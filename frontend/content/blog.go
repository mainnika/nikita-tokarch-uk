package content

import (
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/ghost"
)

// Blog content data
type Blog struct {
	_ interface{} `template:"blog.go.tmpl"`
	ghost.Meta
	Pinned []ghost.Post
	Posts  []ghost.Post
}
