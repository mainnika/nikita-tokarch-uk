package templates

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"

	"github.com/sirupsen/logrus"
)

// Static go-templates source
var content embed.FS

// List of compiled go-templates
var Templates *template.Template

// Load embeded templates
func init() {

	Templates = template.New("")

	tmplNames, err := fs.Glob(content, "*.go.tmpl")
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)

	for _, name := range tmplNames {

		buf.Reset()

		tmplContent, err := content.Open(name)
		if err != nil {
			panic(err)
		}
		size, err := buf.ReadFrom(tmplContent)
		if err != nil {
			panic(err)
		}
		tmpl, err := Templates.New(name).Parse(buf.String())
		if err != nil {
			panic(err)
		}

		logrus.Debugf("Found template: %s, size:%d", tmpl.Name(), size)
	}

	logrus.Debugf("Templates loading complete%s", Templates.DefinedTemplates())
}
