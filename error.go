package client

import (
	"github.com/juju/errgo/errors"
)

var (
	ErrUnexpectedResponse = errors.New("Unexpected response from user service")
	ErrCompanyNotFound    = errors.New("Company not found.")

	Mask = errors.MaskFunc()
)
