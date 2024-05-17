package appError

import "errors"

var (
	Internal = errors.New("internal app error")
)
