package client

import (
	"github.com/catalyst-zero/api-schema"
)

// FindCompaniesByUser returns all companies that the given userID is a member of.
func (c *Client) FindCompaniesByUser(userID string) ([]string, error) {
	zeroVal := []string{}

	resp, err := c.get(c.endpointUrl("/v1/user/" + userID))
	if err != nil {
		return zeroVal, Mask(err)
	}

	if err := mapCommonApiSchemaErrors(resp); err != nil {
		return zeroVal, Mask(err)
	}

	result := make([]string, 0)
	if ok, err := apischema.IsStatusData(&resp.Body); err != nil {
		return zeroVal, Mask(err)
	} else if ok {
		if err := apischema.ParseData(&resp.Body, &result); err != nil {
			return zeroVal, Mask(err)
		}
		return result, nil
	}
	return zeroVal, ErrUnexpectedResponse
}
