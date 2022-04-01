package users

import "errors"

var (
	ErrUserExistWithUsername = errors.New("Username already exist in database")
)
