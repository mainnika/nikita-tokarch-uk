package content

// Error content data
type Error struct {
	_ interface{} `template:"error.go.tmpl"`

	Message string
}

// Title returns error title
func (e Error) Title() string {
	return e.Message
}

// Description returns error description
func (e Error) Description() string {
	return e.Message
}
