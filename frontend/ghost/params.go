package ghost

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

// QueryParam is the generic query param applier
type QueryParam interface {
	Apply(headers *fasthttp.RequestHeader, args *fasthttp.Args)
}

// QueryLimit returns limit param query
func QueryLimit(limit int) queryLimit {
	return queryLimit{limit: limit}
}

// QueryPage returns page param query
func QueryPage(page int) queryPage {
	return queryPage{page: page}
}

// queryLimit implements the limit param query applier
type queryLimit struct{ limit int }

// Apply applies the limit argument to the query
func (ql queryLimit) Apply(headers *fasthttp.RequestHeader, args *fasthttp.Args) {

	if ql.limit == 0 {
		return
	}

	args.Add("limit", strconv.Itoa(ql.limit))
}

// queryPage implements the page param query applier
type queryPage struct{ page int }

// Apply applies the page argument to the query
func (qp queryPage) Apply(headers *fasthttp.RequestHeader, args *fasthttp.Args) {

	if qp.page < 2 {
		return
	}

	args.Add("page", strconv.Itoa(qp.page))
}
