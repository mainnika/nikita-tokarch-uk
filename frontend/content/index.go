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

// Title returns index title
func (i Index) Title() string {

	if len(i.Pinned) > 0 {
		return i.Pinned[0].Title
	}
	if len(i.Posts) > 0 {
		return i.Posts[0].Title
	}

	return "UNKNOWN:"
}

// Description returns index description
func (i Index) Description() string {
	return "TODO:"
}
