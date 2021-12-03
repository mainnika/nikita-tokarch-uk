package renderer

import (
	"sync"

	routing "github.com/jackwhelpton/fasthttp-routing/v2"
	"github.com/valyala/fasthttp"
)

// Renderer is the main handler that contains all routes handlers
type Renderer struct {
	Base string

	router  *routing.Router
	handler fasthttp.RequestHandler

	initOnce sync.Once
}

// Handler invokes the lazy once-initializer and then does the request
func (r *Renderer) Handler(ctx *fasthttp.RequestCtx) {
	r.initOnce.Do(r.init)
	r.handler(ctx)
}

// init has the renderer initialization
func (r *Renderer) init() {

	router := routing.New()

	router.Use(r.useTemplateWriter)
	router.Use(r.useErrorHandler)
	router.NotFound(r.errorNotFound)

	r.router = router
	r.handler = router.HandleRequest
}
