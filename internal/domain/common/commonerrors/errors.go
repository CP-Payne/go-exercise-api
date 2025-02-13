package commonerrors

import "errors"

var (
	ErrInvalidUpdatedAt = errors.New("date connot be in the past")
)
