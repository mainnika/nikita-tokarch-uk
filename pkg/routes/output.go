package routes

import (
	"io"

	routing "github.com/jackwhelpton/fasthttp-routing/v2"
	"github.com/valyala/fasthttp"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/templates"
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

	template := templates.MustGetTemplateOf(content)

	return template.Execute(w, content)
}

// useTemplateWriter is the routing middleware to set the default data writer
func (r *Routes) useTemplateWriter(c *routing.Context) (err error) {

	c.SetDataWriter(staticWriter)

	return
}
