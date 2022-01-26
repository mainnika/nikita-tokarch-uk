package templates

const (
	URLPostfix = "aspx"

	URLRoot  = "/"
	URLSlug  = "/<slug:[^/\\.]*>." + URLPostfix
	URLIndex = "/index." + URLPostfix
	URLBlog  = "/blog." + URLPostfix
	URLPost  = "/post." + URLPostfix
	URLJSApp = "/js-bin/app.js"
)
