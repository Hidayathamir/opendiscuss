package environtment

import (
	"errors"
	"os"
	"strconv"
)

const (
	_ENV_KEY_DB_CONNECTION   = "DB_CONNECTION"
	_ENV_KEY_LOCATION        = "LOCATION"
	_ENV_KEY_JWT_SIGN_KEY    = "JWT_SIGN_KEY"
	_ENV_KEY_JWT_EXPIRE_HOUR = "JWT_EXPIRE_HOUR"
)

var (
	DB_CONNECTION   = "configure by env"
	LOCATION        = "configure by env"
	JWT_SIGN_KEY    = "configure by env"
	JWT_EXPIRE_HOUR = 0
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

	jwtExpireHourEnv := os.Getenv(_ENV_KEY_JWT_EXPIRE_HOUR)
	if jwtExpireHourEnv == "" {
		return errors.New("JWT_EXPIRE_HOUR not found in os environtment")
	}
	var err error
	JWT_EXPIRE_HOUR, err = strconv.Atoi(jwtExpireHourEnv)
	if err != nil {
		return err
	}

	return nil
}
