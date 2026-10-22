package book

import "errors"

var (
	ErrBookNotFound  = errors.New("book not found")
	ErrCmdRepository = errors.New("unable to command repository")
)
