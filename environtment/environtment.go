package environtment

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
)

const (
	_ENV_DB_USER             = "DB_USER"
	_ENV_DB_PASSWORD         = "DB_PASSWORD"
	_ENV_DB_HOST             = "DB_HOST"
	_ENV_DB_PORT             = "DB_PORT"
	_ENV_DB_NAME             = "DB_NAME"
	_ENV_KEY_JWT_SIGN_KEY    = "JWT_SIGN_KEY"
	_ENV_KEY_JWT_EXPIRE_HOUR = "JWT_EXPIRE_HOUR"
)

var (
	DB_USER         = "configure by env"
	DB_PASSWORD     = "configure by env"
	DB_HOST         = "configure by env"
	DB_PORT         = "configure by env"
	DB_NAME         = "configure by env"
	JWT_SIGN_KEY    = "configure by env"
	JWT_EXPIRE_HOUR = 0
)

func InitEnv() error {
	DB_USER = os.Getenv(_ENV_DB_USER)
	if DB_USER == "" {
		return errors.New("DB_USER not found in os environtment")
	}

	DB_PASSWORD = os.Getenv(_ENV_DB_PASSWORD)
	if DB_PASSWORD == "" {
		return errors.New("DB_PASSWORD not found in os environtment")
	}

	DB_HOST = os.Getenv(_ENV_DB_HOST)
	if DB_HOST == "" {
		return errors.New("DB_HOST not found in os environtment")
	}

	DB_PORT = os.Getenv(_ENV_DB_PORT)
	if DB_PORT == "" {
		return errors.New("DB_PORT not found in os environtment")
	}

	DB_NAME = os.Getenv(_ENV_DB_NAME)
	if DB_NAME == "" {
		return errors.New("DB_NAME not found in os environtment")
	}

	JWT_SIGN_KEY = os.Getenv(_ENV_KEY_JWT_SIGN_KEY)
	if JWT_SIGN_KEY == "" {
		return errors.New("JWT_SIGN_KEY not found in os environtment")
	}

	jwtExpireHourEnv := os.Getenv(_ENV_KEY_JWT_EXPIRE_HOUR)
	if jwtExpireHourEnv == "" {
		return errors.New("JWT_EXPIRE_HOUR not found in os environtment")
	}
	var err error
	JWT_EXPIRE_HOUR, err = strconv.Atoi(jwtExpireHourEnv)
	if err != nil {
		return errors.Wrap(err, "error convert string to int")
	}

	return nil
}
