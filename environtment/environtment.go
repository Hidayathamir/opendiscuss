package environtment

import (
	"errors"
	"os"
)

const (
	_ENV_KEY_DB_CONNECTION = "DB_CONNECTION"
	_ENV_KEY_LOCATION      = "LOCATION"
	_ENV_KEY_JWT_SIGN_KEY  = "JWT_SIGN_KEY"
)

var (
	DB_CONNECTION = "configure by env"
	LOCATION      = "configure by env"
	JWT_SIGN_KEY  = "configure by env"
)

func InitEnv() error {
	DB_CONNECTION = os.Getenv(_ENV_KEY_DB_CONNECTION)
	if DB_CONNECTION == "" {
		return errors.New("DB_CONNECTION not found in os environtment")
	}

	LOCATION = os.Getenv(_ENV_KEY_LOCATION)

	JWT_SIGN_KEY = os.Getenv(_ENV_KEY_JWT_SIGN_KEY)
	if JWT_SIGN_KEY == "" {
		return errors.New("JWT_SIGN_KEY not found in os environtment")
	}

	return nil
}
