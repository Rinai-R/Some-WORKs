package response

import (
	"errors"
)

var (
	ErrPasswordLength = errors.New("password error: length must be between 5 and 20 characters")
	ErrNameLength     = errors.New("name error: length must be between 5 and 20 characters")
	PasswordError     = errors.New("password error: password error")
)
