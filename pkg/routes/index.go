package routes

import (
	"bytes"

	routing "github.com/jackwhelpton/fasthttp-routing/v2"
	"github.com/valyala/fasthttp"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/content"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost/params"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/templates"
)

// relativeRedirectBytes makes a relative redirect by using http Location header
func (r *Routes) relativeRedirectBytes(c *routing.Context, location []byte, statusCode int) (err error) {

	c.Response.Header.SetCanonical([]byte(fasthttp.HeaderLocation), location)
	c.Response.SetStatusCode(statusCode)

	return
}

// rootRedirect redirects the root url to the index using http redirect
func (r *Routes) rootRedirect(c *routing.Context) (err error) {
	return r.relativeRedirectBytes(c, []byte(templates.URLIndex), fasthttp.StatusFound)
}

// rootRedirect forcefully adds postfix to the url
func (r *Routes) usePostfixForce(c *routing.Context) (err error) {

	fullPath := c.Path()
	if len(fullPath) <= 1 {
		return c.Next()
	}

	dotIndex := bytes.LastIndexByte(fullPath, '.')
	if dotIndex >= 0 {
		return c.Next()
	}

	fullPath = append(fullPath, '.')
	fullPath = append(fullPath, []byte(templates.URLPostfix)...)

	return r.relativeRedirectBytes(c, fullPath, fasthttp.StatusFound)
}

// index handler renders index data
func (r *Routes) index(c *routing.Context) (err error) {

	pinnedPageSlug := r.ContentConfig.Pinned
	postsPerPage := r.ContentConfig.PostsPerPage

	pinnedPages, err := r.GhostClient.GetPageBySlug(pinnedPageSlug)
	if err != nil {
		return
	}

	latestPosts, err := r.GhostClient.GetPosts(params.WithLimit(postsPerPage))
	if err != nil {
		return
	}

	indexContent := content.Index{
		Pinned: pinnedPages.Pages,
		Meta:   latestPosts.Meta,
		Posts:  latestPosts.Posts,
	}

	return c.Write(indexContent)
}
