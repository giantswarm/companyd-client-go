package client

import (
	errors "github.com/juju/errgo"

	"github.com/catalyst-zero/api-schema"
	"net/http"
)

var (
	ErrUnexpectedResponse   = errors.New("Unexpected response from companyd service")
	ErrCompanyNotFound      = errors.New("Company not found.")
	ErrCompanyAlreadyExists = errors.New("Company already exists.")

	ErrWrongInput = errors.New("Wrong input.")

	Mask = errors.MaskFunc(
		IsErrUnexpectedResponse, IsErrCompanyNotFound,
		IsErrCompanyAlreadyExists, IsErrWrongInput,
	)
)

func IsErrWrongInput(err error) bool {
	return errors.Cause(ErrWrongInput) == err
}

func IsErrCompanyNotFound(err error) bool {
	return errors.Cause(ErrCompanyNotFound) == err
}

func IsErrUnexpectedResponse(err error) bool {
	return errors.Cause(ErrUnexpectedResponse) == err
}

func IsErrCompanyAlreadyExists(err error) bool {
	return errors.Cause(ErrCompanyAlreadyExists) == err
}

func mapCommonApiSchemaErrors(resp *http.Response) error {
	if ok, err := apischema.IsStatusWrongInput(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrWrongInput)
	}

	if ok, err := apischema.IsStatusResourceNotFound(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrCompanyNotFound)
	}

	if ok, err := apischema.IsStatusResourceAlreadyExists(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrCompanyAlreadyExists)
	}

	return nil
}
