package content

// Error content data
type Error struct {
	_ interface{} `template:"error.go.tmpl"`

	Message string
}
