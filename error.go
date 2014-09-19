package client

import (
	"github.com/juju/errgo/errors"

	"github.com/catalyst-zero/api-schema"
	"net/http"
)

var (
	ErrUnexpectedResponse   = errors.New("Unexpected response from companyd service")
	ErrCompanyNotFound      = errors.New("Company not found.")
	ErrCompanyAlreadyExists = errors.New("Company already exists.")

	ErrWrongInput = errors.New("Wrong input.")

	Mask = errors.MaskFunc()
)

func mapCommonApiSchemaErrors(resp *http.Response) error {
	if ok, err := apischema.IsStatusWrongInput(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrWrongInput)
	}

	if ok, err := apischema.IsStatusResourceAlreadyExists(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrCompanyAlreadyExists)
	}

	return nil
}
