package content

import (
	"fmt"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/ghost"
)

// Blog content data
type Blog struct {
	_ interface{} `template:"blog.go.tmpl"`
	ghost.Meta
	Pinned []ghost.Post
	Posts  []ghost.Post
}

// Title returns blog content title
func (i Blog) Title() string {
	return fmt.Sprintf("... %d of %d", i.Meta.Pagination.Page, i.Meta.Pagination.Pages)
}

// Description returns blog content description
func (i Blog) Description() string {
	return "TODO:"
}
