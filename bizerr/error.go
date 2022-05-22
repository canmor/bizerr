package bizerr

import "errors"

type Error struct {
	error
}

// New Construct Error, intended to return Error instead of pointer to Error inorder to
// define concrete error as package variable easier.
func New(key string) Error {
	return Error{errors.New(key)}
}

// WithParam Wrap error with params
func (e Error) WithParam(params ...string) error {
	return newParams(e, params...)
}
