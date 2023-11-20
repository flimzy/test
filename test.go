package test

import (
	"github.com/pkg/errors"
)

func Test() error {
	return errors.New("an error")
}
