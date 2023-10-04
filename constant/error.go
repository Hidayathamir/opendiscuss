package constant

import "errors"

var (
	INVALID_REQUEST_BODY = errors.New("invalid request body")
	USER_NOT_FOUND       = errors.New("user not found")
)
