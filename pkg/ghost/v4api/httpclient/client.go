package httpclient

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost/data"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost/params"
)

var _ ghost.Client = (*HTTPClient)(nil)

// Ghost content data URIs:
const (
	ghostAPIPrefix        = "/ghost/api/v3/"
	ghostAPIGetPosts      = ghostAPIPrefix + "content/posts/"
	ghostAPIGetPostBySlug = ghostAPIPrefix + "content/posts/slug/%s/"
	ghostAPIGetPageBySlug = ghostAPIPrefix + "content/pages/slug/%s/"
)

// HTTPClient implements the ghost http client
type HTTPClient struct {
	QueryTimeout time.Duration
	ContentKey   string
	Addr         string
	Secured      bool
	Headers      map[string]string

	client *fasthttp.HostClient

	setupClientOnce sync.Once
}

// setupClient creates the default http client
func (g *HTTPClient) setupClient() {

	g.client = &fasthttp.HostClient{
		Addr:  g.Addr,
		IsTLS: g.Secured,

		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
	}
}

// doQuery does the method and unmarshals the result into the easyjson Unmarshaler
func (g *HTTPClient) doQuery(path string, v easyjson.Unmarshaler, params params.Params) (err error) {

	g.setupClientOnce.Do(g.setupClient)

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(res)
		fasthttp.ReleaseRequest(req)
	}()

	g.setupRequest(path, req)
	g.applyParams(params, req)

	err = g.client.DoTimeout(req, res, g.QueryTimeout)
	if err != nil {
		return
	}
	if res.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("non OK status code: %d", res.StatusCode())
	}

	resBytes := res.Body()
	if resBytes == nil && v == nil {
		return fmt.Errorf("nothing to unmarshal")
	}
	if resBytes == nil {
		return
	}

	err = easyjson.Unmarshal(resBytes, v)

	return
}

// setupRequest does the necessary initial configuration to the http request
func (g *HTTPClient) setupRequest(path string, req *fasthttp.Request) {

	uri := req.URI()

	scheme := "http"
	if g.Secured {
		scheme = "https"
	}

	uri.SetHost(g.Addr)
	uri.SetPath(path)
	uri.SetScheme(scheme)

	uri.QueryArgs().Add("key", g.ContentKey)

	for hKey, hValue := range g.Headers {
		req.Header.Add(hKey, hValue)
	}
}

// applyParams function additionally configure the http request using params
func (g *HTTPClient) applyParams(p params.Params, req *fasthttp.Request) (err error) {

	uri := req.URI()

	limit := p.Limit
	if limit > 0 {
		uri.QueryArgs().Add("limit", strconv.Itoa(limit))
	}

	page := p.Page
	if page > 1 {
		uri.QueryArgs().Add("page", strconv.Itoa(page))
	}

	return
}

// GetPageBySlug returns the only one page using slug filter
func (g *HTTPClient) GetPageBySlug(slug string, queryModifiers ...params.Modifier) (pages *data.Pages, err error) {

	pages = &data.Pages{}
	defaultParams := params.Params{}
	method := fmt.Sprintf(ghostAPIGetPageBySlug, slug)

	err = g.doQuery(method, pages, defaultParams)
	if err != nil {
		pages = nil
	}

	return
}

// GetPosts returns posts
func (g *HTTPClient) GetPosts(queryModifiers ...params.Modifier) (posts *data.Posts, err error) {

	posts = &data.Posts{}
	defaultParams := params.Params{}
	combinedParams := params.Modifiers(queryModifiers).Apply(defaultParams)

	err = g.doQuery(ghostAPIGetPosts, posts, combinedParams)
	if err != nil {
		posts = nil
	}

	return
}

// GetPostBySlug returns the only one post using slug filter
func (g *HTTPClient) GetPostBySlug(slug string, queryModifiers ...params.Modifier) (posts *data.Posts, err error) {

	posts = &data.Posts{}
	defaultParams := params.Params{}
	combinedParams := params.Modifiers(queryModifiers).Apply(defaultParams)
	method := fmt.Sprintf(ghostAPIGetPostBySlug, slug)

	err = g.doQuery(method, posts, combinedParams)
	if err != nil {
		posts = nil
	}

	return
}
