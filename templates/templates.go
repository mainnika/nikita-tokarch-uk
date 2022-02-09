package templates

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"reflect"

	"github.com/sirupsen/logrus"
)

// Static go-templates source
//go:embed blog.go.tmpl
//go:embed error.go.tmpl
//go:embed head.go.tmpl
//go:embed index.go.tmpl
//go:embed menu.go.tmpl
//go:embed post.go.tmpl
//go:embed analytics.go.tmpl
var content embed.FS

// Templates is a container of compiled templates
var Templates *template.Template = template.New("")

// Load embedded templates
func Load(funcs *Funcs) (err error) {

	Templates.Funcs(funcs.Use())

	tmplNames, err := fs.Glob(content, "*.go.tmpl")
	if err != nil {
		return fmt.Errorf("cannot match templates names using glob, %w", err)
	}

	buf := bytes.NewBuffer(nil)

	for _, name := range tmplNames {

		buf.Reset()

		tmplContent, err := content.Open(name)
		if err != nil {
			return fmt.Errorf("cannot open template content, name:%s, %w", name, err)
		}
		size, err := buf.ReadFrom(tmplContent)
		if err != nil {
			return fmt.Errorf("cannot read template content, name:%s, %w", name, err)
		}
		tmpl, err := Templates.New(name).Parse(buf.String())
		if err != nil {
			return fmt.Errorf("cannot parse template, name:%s, %w", name, err)
		}

		logrus.Debugf("Found template: %s, size:%d", tmpl.Name(), size)
	}

	logrus.Debugf("Templates loading complete%s", Templates.DefinedTemplates())

	return
}

// MustLookup wraps lookup function for the root template namespace
func MustLookup(name string) *template.Template {

	tmpl := Templates.Lookup(name)
	if tmpl == nil {
		panic(fmt.Errorf("cannot find template %s", name))
	}

	return tmpl
}

// MustGetTemplateOf returns template which is mapped to the content data
func MustGetTemplateOf(content interface{}) (template *template.Template) {

	el := reflect.TypeOf(content)
	numField := el.NumField()

	for i := 0; i < numField; i++ {
		field := el.Field(i)
		tag := field.Tag
		found, ok := tag.Lookup("template")
		if !ok {
			continue
		}

		return MustLookup(found)
	}

	panic(fmt.Errorf("content %v does not have a template tag", content))
}
