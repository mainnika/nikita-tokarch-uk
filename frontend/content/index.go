package content

import (
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/ghost"
)

// Index content data
type Index struct {
	_ interface{} `template:"index.go.tmpl"`
	ghost.Meta
	Pinned []ghost.Post
	Posts  []ghost.Post
}
