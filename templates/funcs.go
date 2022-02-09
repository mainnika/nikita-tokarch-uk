package templates

import (
	"html/template"
	"net/url"
	"sync"
)

type Funcs struct {
	Version string

	YandexKey string

	compiledJSAppURL string

	initOnce sync.Once
}

func (f *Funcs) init() {

	jsAppURL, err := url.Parse(URLJSApp)
	if err != nil {
		panic(err)
	}

	{
		q := jsAppURL.Query()
		q.Add("version", f.Version)

		jsAppURL.RawQuery = q.Encode()
	}

	f.compiledJSAppURL = jsAppURL.String()
}

func (f *Funcs) add(i int) int {
	return i + 1
}

func (f *Funcs) sub(i int) int {
	return i - 1
}

func (f *Funcs) getJSAppURL() string {

	f.initOnce.Do(f.init)

	return f.compiledJSAppURL
}

func (f *Funcs) getIndexURL() string {
	return URLIndex
}

func (f *Funcs) getBlogURL() string {
	return URLBlog
}

func (f *Funcs) getYaKey() string {
	return f.YandexKey
}

// Use returns a func map with template helpers functions
func (f *Funcs) Use() template.FuncMap {
	return template.FuncMap{
		"add":         f.add,
		"sub":         f.sub,
		"getJSAppURL": f.getJSAppURL,
		"getIndexURL": f.getIndexURL,
		"getBlogURL":  f.getBlogURL,
		"getYaKey":    f.getYaKey,
	}
}
