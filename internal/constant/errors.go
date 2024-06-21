package constant

import (
	"errors"
)

var (
	// Technical Errors

	ErrUnauthorized    = errors.New("unauthorized")
	ErrNoAccess        = errors.New("no access")
	ErrNotFound        = errors.New("not found")
	ErrDuplicateRecord = errors.New("record is duplicate")
	ErrUnimplemented   = errors.New("unimplemented")
	ErrUserNotFound    = errors.New("user not found")

	// Business Errors

)
