package bizerr

import (
	"errors"
	"fmt"
)

type withParams struct {
	error
	params []string
}

func newParams(err Error, params ...string) *withParams {
	return &withParams{error: err, params: params}
}

func (e *withParams) Error() string {
	return fmt.Sprintf("%v, params %v", e.error, e.params)
}

func (e *withParams) Unwrap() error {
	return e.error
}

// ExtractParams Unwrap withParams from error and return its withParams slice
// err passed in should wrapped with type `withParams`, otherwise it returns nil slice
func ExtractParams(err error) []string {
	var p *withParams
	if errors.As(err, &p) {
		return p.params
	}
	return nil
}