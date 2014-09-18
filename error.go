package client

import (
	"github.com/juju/errgo/errors"
)

var (
	ErrUnexpectedResponse   = errors.New("Unexpected response from companyd service")
	ErrCompanyNotFound      = errors.New("Company not found.")
	ErrCompanyAlreadyExists = errors.New("Company already exists.")

	ErrWrongInput = errors.New("Wrong input.")

	Mask = errors.MaskFunc()
)
