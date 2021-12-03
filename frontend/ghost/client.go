package ghost

import (
	"fmt"
	"sync"
	"time"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)
// Ghost content data URIs:
const (
	ghostAPIPrefix        = "/ghost/api/v3/"
	ghostAPIGetPageBySlug = ghostAPIPrefix + "content/pages/slug/%s/"
)

// HTTPClient implements the ghost http client
type HTTPClient struct {
	QueryTimeout time.Duration
	ContentKey   string
	Addr         string
	Secured      bool

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
func (g *HTTPClient) doQuery(method string, v easyjson.Unmarshaler, params ...QueryParam) (err error) {

	g.setupClientOnce.Do(g.setupClient)

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(res)
		fasthttp.ReleaseRequest(req)
	}()

	uri := req.URI()
	uri.SetHost(g.Addr)
	uri.SetPath(method)
	uri.QueryArgs().Add("key", g.ContentKey)
	if g.client.IsTLS {
		uri.SetScheme("https")
	}

	for _, param := range params {
		param.Apply(&req.Header, uri.QueryArgs())
	}

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

// GetPageBySlug returns the only one page using slug filter
func (g *HTTPClient) GetPageBySlug(slug string) (pages *Pages, err error) {

	pages = &Pages{}
	method := fmt.Sprintf(ghostAPIGetPageBySlug, slug)

	err = g.doQuery(method, pages)
	if err != nil {
		pages = nil
	}

	return
}