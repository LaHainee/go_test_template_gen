package model

import "errors"

var (
	ErrInvalidFile = errors.New("invalid file")
	ErrUnsupported = errors.New("unsupported")
	ErrNotFound    = errors.New("not found")
)
