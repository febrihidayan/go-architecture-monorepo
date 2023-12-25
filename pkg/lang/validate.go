package lang

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNameRequired            = errors.New("name is required.")
	ErrUserIdRequired          = errors.New("UserId is required.")
	ErrEmailRequired           = errors.New("Email is required.")
	ErrPasswordRequired        = errors.New("Password is required.")
	ErrConfirmPasswordRequired = errors.New("ConfirmPassword is required.")
	ErrConfitmPasswordNotSame  = errors.New("Password confirmation does not match.")
	ErrRoleRequired            = errors.New("Role is required.")
	ErrEmailNotFound           = errors.New("The email has not been registered.")
	ErrEmailAlready            = errors.New("The email is already.")
	ErrPasswordIsIncorrent     = errors.New("The password you entered is incorrect.")
	ErrRoleAlready             = errors.New("The name role is already.")
	ErrPermissionAlready       = errors.New("The name permission is already.")
	ErrUnsupportFile           = errors.New("unsupport file type.")

	Locales = map[string]string{
		"filled": "The :attribute field is required.",
	}
)

func Trans(field string, attributes ...interface{}) error {
	var err string

	if str, ok := Locales[field]; ok {
		err = strings.ReplaceAll(str, ":attribute", fmt.Sprintf("%v", attributes))
	} else {
		err = "not found message."
	}

	return errors.New(err)
}
