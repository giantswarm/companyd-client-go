package client

import (
	"github.com/giantswarm/api-schema"
)

// Creates a new company with the given ID (can be a UUID or slug version of the actual name).
func (c *Client) CreateCompany(companyID string, fields CompanyFields) error {

	request := struct {
		CompanyID string `json:"company_id"`
		CompanyFields
	}{
		CompanyID:     companyID,
		CompanyFields: fields,
	}

	resp, err := c.postJson(c.endpointUrl("/v1/company/"), request)
	if err != nil {
		return Mask(err)
	}

	if err := mapCommonApiSchemaErrors(resp); err != nil {
		return Mask(err)
	}

	if ok, err := apischema.IsStatusResourceCreated(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return nil
	}

	return Mask(ErrUnexpectedResponse)
}

// DeleteCompany deletes a new company with the given ID.
func (c *Client) DeleteCompany(companyID string) error {
	resp, err := c.post(c.endpointUrl("/v1/company/"+companyID+"/delete"), "", nil)
	if err != nil {
		return Mask(err)
	}

	if err := mapCommonApiSchemaErrors(resp); err != nil {
		return Mask(err)
	}

	if ok, err := apischema.IsStatusResourceDeleted(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return nil
	}

	return ErrUnexpectedResponse
}

func (c *Client) AddMembers(companyID string, members []string) error {
	resp, err := c.postJson(c.endpointUrl("/v1/company/"+companyID+"/members/add"), members)
	if err != nil {
		return Mask(err)
	}

	if err := mapCommonApiSchemaErrors(resp); err != nil {
		return Mask(err)
	}

	if ok, err := apischema.IsStatusResourceCreated(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return nil
	}

	return ErrUnexpectedResponse
}

func (c *Client) RemoveMembers(companyID string, members []string) error {
	resp, err := c.postJson(c.endpointUrl("/v1/company/"+companyID+"/members/remove"), members)
	if err != nil {
		return Mask(err)
	}

	if err := mapCommonApiSchemaErrors(resp); err != nil {
		return Mask(err)
	}

	if ok, err := apischema.IsStatusResourceDeleted(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return nil
	}

	return ErrUnexpectedResponse
}

type CompanyFields struct {
	DefaultCluster string `json:"default_cluster"`
}

type Company struct {
	CompanyID string   `json:"company_id"`
	Members   []string `json:"members"`

	CompanyFields
}

func (c *Client) GetCompany(companyID string, company *Company) error {
	resp, err := c.get(c.endpointUrl("/v1/company/" + companyID + "/"))
	if err != nil {
		return Mask(err)
	}

	if err := mapCommonApiSchemaErrors(resp); err != nil {
		return Mask(err)
	}

	if ok, err := apischema.IsStatusData(&resp.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(apischema.ParseData(&resp.Body, company))
	}
	return ErrUnexpectedResponse
}
