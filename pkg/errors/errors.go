package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

func Wrapf(err error, formatString string, args ...interface{}) error {
	return errors.Wrap(err, fmt.Sprintf(formatString, args...))
}

func Wrap(err error, formatString string) error {
	return errors.Wrap(err, fmt.Sprintf(formatString, nil))
}
