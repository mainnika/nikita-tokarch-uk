package renderer

import (
	"io"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/templates"
	routing "github.com/jackwhelpton/fasthttp-routing/v2"
	"github.com/valyala/fasthttp"
)

var _ routing.DataWriter = (*TemplateWriter)(nil)

// staticWriter is thread-safe static instance of template writer
var staticWriter = &TemplateWriter{}

// TemplateWriter is the fasthttp data writer that loads and executes template using the content
type TemplateWriter struct{}

// SetHeader sets the content type to HTML since all templates are HTML
func (tw *TemplateWriter) SetHeader(rh *fasthttp.ResponseHeader) {
	rh.SetContentType(routing.MIME_HTML)
}

// Write executes the template and writes result to the response writer
func (tw *TemplateWriter) Write(w io.Writer, content interface{}) error {

	template := templates.GetTemplateOf(content)

	return template.Execute(w, content)
}

// useTemplateWriter is the routing middleware to set the default data writer
func (r *Renderer) useTemplateWriter(c *routing.Context) (err error) {

	c.SetDataWriter(staticWriter)

	return
}
