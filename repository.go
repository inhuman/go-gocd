package gocd

import (
	"github.com/hashicorp/go-multierror"
	"fmt"
	"encoding/json"
)

func (c *DefaultClient) GetAllRepositories(pipelineName string) (*AllPackageRepositories, error) {
	var multiError *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve("/go/api/admin/repositories")).
	//Repositories endpoints works only with api v1 header
		Set("Accept", "application/vnd.go.cd.v1+json").
		End()

	if errs != nil {
		multiError = multierror.Append(multiError, errs...)
		return nil, multiError.ErrorOrNil()
	}
	var AllRepositories AllPackageRepositories

	jsonErr := json.Unmarshal([]byte(body), &AllRepositories)
	if jsonErr != nil {
		multiError = multierror.Append(multiError, jsonErr)
		return nil, multiError.ErrorOrNil()
	}

	return &AllRepositories, multiError.ErrorOrNil()
}

func (c *DefaultClient) GetRepository(id string) (*PackageRepository, error) {
	var multiError *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/admin/repositories/%s", id))).
	//Repositories endpoints works only with api v1 header
		Set("Accept", "application/vnd.go.cd.v1+json").
		End()

	if errs != nil {
		multiError = multierror.Append(multiError, errs...)
		return nil, multiError.ErrorOrNil()
	}
	var PackageRepository PackageRepository

	jsonErr := json.Unmarshal([]byte(body), &PackageRepository)
	if jsonErr != nil {
		multiError = multierror.Append(multiError, jsonErr)
		return nil, multiError.ErrorOrNil()
	}

	return &PackageRepository, multiError.ErrorOrNil()
}
