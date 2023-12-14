package lang

import "errors"

var (
	ErrNameRequired            = errors.New("name is required")
	ErrUserIdRequired          = errors.New("UserId is required")
	ErrEmailRequired           = errors.New("Email is required")
	ErrPasswordRequired        = errors.New("Password is required")
	ErrConfirmPasswordRequired = errors.New("ConfirmPassword is required")
	ErrConfitmPasswordNotSame  = errors.New("Password confirmation does not match")
	ErrRoleRequired            = errors.New("Role is required")
)
