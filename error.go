package client

import (
	errors "github.com/juju/errgo"

	"github.com/giantswarm/api-schema"
	"net/http"
)

var (
	ErrUnexpectedResponse   = errors.New("Unexpected response from companyd service")
	ErrCompanyNotFound      = errors.New("Company not found.")
	ErrCompanyAlreadyExists = errors.New("Company already exists.")
	ErrMemberAlreadyExists  = errors.New("Member already exists")
	ErrMemberNotFound       = errors.New("Member not found")

	ErrWrongInput = errors.New("Wrong input.")

	Mask = errors.MaskFunc(
		IsErrUnexpectedResponse, IsErrCompanyNotFound,
		IsErrCompanyAlreadyExists, IsErrWrongInput,
		IsErrMemberAlreadyExists, IsErrMemberNotFound,
	)
)

const (
	// Make sure these strings are identical to companyd/middlewares/v1/v1.go#mapError
	reasonMemberAlreadyExists = "Member already exists"
	reasonMemberNotFound      = "Member not found."
)

func IsErrWrongInput(err error) bool {
	return errors.Cause(err) == ErrWrongInput
}

func IsErrCompanyNotFound(err error) bool {
	return errors.Cause(err) == ErrCompanyNotFound
}

func IsErrUnexpectedResponse(err error) bool {
	return errors.Cause(err) == ErrUnexpectedResponse
}

func IsErrCompanyAlreadyExists(err error) bool {
	return errors.Cause(err) == ErrCompanyAlreadyExists
}

func IsErrMemberAlreadyExists(err error) bool {
	return errors.Cause(err) == ErrMemberAlreadyExists
}

func IsErrMemberNotFound(err error) bool {
	return errors.Cause(err) == ErrMemberNotFound
}

func mapCommonApiSchemaErrors(resp *http.Response) error {
	if ok, err := apischema.IsStatusWrongInputWithReason(&resp.Body, reasonMemberAlreadyExists); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrMemberAlreadyExists)
	}

	if ok, err := apischema.IsStatusWrongInputWithReason(&resp.Body, reasonMemberNotFound); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrMemberNotFound)
	}

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
