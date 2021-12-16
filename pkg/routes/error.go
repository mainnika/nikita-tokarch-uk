package routes

import (
	"fmt"
	"net/http"

	routing "github.com/jackwhelpton/fasthttp-routing/v2"
	"github.com/sirupsen/logrus"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/content"
)

// errorNotFound renders http error-404 template
func (r *Routes) errorNotFound(c *routing.Context) (err error) {

	errorContent := content.Error{Message: "not found"}

	return c.Write(errorContent)
}

// useErrorHandler is the middleware that catch handlers errors and render error template
func (r *Routes) useErrorHandler(c *routing.Context) (err error) {

	worker := func() (err error) {

		defer func() {
			r := recover()
			if r == nil {
				return
			}

			err = routing.NewHTTPError(http.StatusInternalServerError,
				fmt.Sprintf("panic:\n%v", r))
		}()

		err = c.Next()

		return
	}

	err = worker()
	if err == nil {
		return
	}

	c.Abort()

	logrus.Warnf("Cannot process request, %v", err)

	errorContent := content.Error{
		Message: err.Error(),
	}

	return c.Write(errorContent)
}
