package params

// Params are generics query argument
type Params struct {
	Limit int
	Page  int
}

// Modifier function takes params and makes some changes
type Modifier func(params Params) Params

// Modifiers is a list of modifier
type Modifiers []Modifier

// Apply function modifies params
func (ms Modifiers) Apply(params Params) Params {

	for _, m := range ms {
		params = m(params)
	}

	return params
}

// WithLimit modifier setups the limit
func WithLimit(limit int) Modifier {
	return func(params Params) Params {
		params.Limit = limit
		return params
	}
}

// WithPage modifier setups the page
func WithPage(page int) Modifier {
	return func(params Params) Params {
		params.Page = page
		return params
	}
}
