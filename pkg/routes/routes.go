package routes

import (
	"sync"

	routing "github.com/jackwhelpton/fasthttp-routing/v2"
	"github.com/valyala/fasthttp"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/config"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/templates"
)

// Routes is the main handler that contains all routes handlers
type Routes struct {
	GhostClient   ghost.Client
	ContentConfig config.Content

	Base string

	router  *routing.Router
	handler fasthttp.RequestHandler

	initOnce sync.Once
}

// Handler invokes the lazy once-initializer and then does the request
func (r *Routes) Handler(ctx *fasthttp.RequestCtx) {
	r.initOnce.Do(r.init)
	r.handler(ctx)
}

// init has the renderer initialization
func (r *Routes) init() {

	router := routing.New()

	router.Use(r.useTemplateWriter)
	router.Use(r.useErrorHandler)
	router.NotFound(r.errorNotFound)

	root := router.Group(r.Base)
	root.Get(templates.URLRoot, r.rootRedirect)
	root.Get(templates.URLIndex, r.index)
	root.Get(templates.URLBlog, r.blog)

	r.router = router
	r.handler = router.HandleRequest
}
