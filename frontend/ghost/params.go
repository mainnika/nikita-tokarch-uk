package ghost

import (
	"github.com/valyala/fasthttp"
)

// QueryParam is the generic query param applier
type QueryParam interface {
	Apply(headers *fasthttp.RequestHeader, args *fasthttp.Args)
}
