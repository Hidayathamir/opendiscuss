package utils

import (
	"time"

	"github.com/Hidayathamir/opendiscuss/environtment"
	"github.com/pkg/errors"
)

func SetTimeZone() error {
	var err error
	location := time.UTC

	if environtment.LOCATION != "" {
		location, err = time.LoadLocation(environtment.LOCATION)
		if err != nil {
			return errors.Wrap(err, "error load timezone")
		}
	}

	time.Local = location
	return nil
}
