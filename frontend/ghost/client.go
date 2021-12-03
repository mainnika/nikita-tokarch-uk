package ghost

import (
	"time"

	"github.com/valyala/fasthttp"
)
// Ghost content data URIs:
const (
	ghostAPIPrefix = "/ghost/api/v3/"
)

// HTTPClient implements the ghost http client
type HTTPClient struct {
	QueryTimeout time.Duration
	ContentKey   string
	Addr         string
	Secured      bool

	client *fasthttp.HostClient
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
