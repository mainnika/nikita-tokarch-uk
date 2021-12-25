package content

import "code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost/data"

// Index content data
type Index struct {
	_ interface{} `template:"index.go.tmpl"`
	data.Meta
	Pinned []data.Post
	Posts  []data.Post
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
