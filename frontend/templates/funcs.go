package templates

import "html/template"

func UseFuncs() template.FuncMap {
	return template.FuncMap{
		"add": func(i int) int {
			return i + 1
		},
		"sub": func(i int) int {
			return i - 1
		},
		"getJSAppURL": func() string {
			return URLJSApp
		},
		"getIndexURL": func() string {
			return URLIndex
		},
		"getBlogURL": func() string {
			return URLBlog
		},
	}
}
