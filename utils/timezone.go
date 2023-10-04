package utils

import (
	"errors"
	"os"
	"time"

	"github.com/Hidayathamir/opendiscuss/constant"
)

func SetTimeZone() error {
	var err error
	location := time.UTC

	envLocation := os.Getenv(constant.LOCATION)
	if envLocation != "" {
		location, err = time.LoadLocation(envLocation)
		if err != nil {
			errorMessage := "error loading timezone:" + err.Error()
			return errors.New(errorMessage)
		}
	}

	time.Local = location
	return nil
}
