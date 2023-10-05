package utils

import (
	"errors"
	"time"

	"github.com/Hidayathamir/opendiscuss/environtment"
)

func SetTimeZone() error {
	var err error
	location := time.UTC

	if environtment.LOCATION != "" {
		location, err = time.LoadLocation(environtment.LOCATION)
		if err != nil {
			errorMessage := "error loading timezone:" + err.Error()
			return errors.New(errorMessage)
		}
	}

	time.Local = location
	return nil
}
