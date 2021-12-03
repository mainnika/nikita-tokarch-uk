package templates

import "html/template"

func UseFuncs() template.FuncMap {
	return template.FuncMap{
		"getIndexURL": func() string {
			return URLIndex
		},
	}
}
