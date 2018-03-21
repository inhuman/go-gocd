package go_gocd

import (
	"github.com/hashicorp/go-multierror"

	"encoding/json"
	"fmt"
	"os"
	"errors"
)

func (c *DefaultClient) CreatePackage(pkg Package) (*Package, *ApiResponse, *multierror.Error) {
	var multiError *multierror.Error

	response, body, errs := c.Request.
		Post(c.resolve("/go/api/admin/packages")).
	//Package endpoints works only with api v1 header
		Set("Accept", "application/vnd.go.cd.v1+json").
		Set("Content-Type", "application/json").
		SendStruct(pkg).
		End()

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string(body))
	}
	var apiResponse ApiResponse

	if errs != nil {
		multiError = multierror.Append(multiError, errs...)
		return nil, nil, multiError
	}

	if response.StatusCode != 200 {
		err := json.Unmarshal([]byte(body), &apiResponse)
		if err != nil {
			multiError = multierror.Append(multiError, err)
			return nil, nil, multiError
		}

		multiError = multierror.Append(multiError, errors.New(apiResponse.Message))

		// Check common pipeline errors
		if len(apiResponse.Data.Errors) > 0 {
			for fieldName, respErrArr := range apiResponse.Data.Errors {
				for _, respErr := range respErrArr {
					multiError = multierror.Append(
						multiError, errors.New("[Common]["+fieldName+"] "+string(respErr)))
				}
			}
		}
	}

	var Package Package

	jsonErr := json.Unmarshal([]byte(body), &Package)
	if jsonErr != nil {
		multiError = multierror.Append(multiError, jsonErr)
	}

	if Package.Name == "" {
		var resp ApiResponse
		jsonErr := json.Unmarshal([]byte(body), &resp)
		if jsonErr != nil {
			multiError = multierror.Append(multiError, jsonErr)
		}
		return nil, &resp, multiError

	} else {
		return &Package, nil, multiError
	}
}

func (c *DefaultClient) DeletePackage(id string) (*ApiResponse, *multierror.Error) {
	var multiError *multierror.Error

	_, body, errs := c.Request.
		Delete(c.resolve(fmt.Sprintf("/go/api/admin/packages/%s", id))).
	//Package endpoints works only with api v1 header
		Set("Accept", "application/vnd.go.cd.v1+json").
		End()

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string(body))
	}

	if errs != nil {
		multiError = multierror.Append(multiError, errs...)
		return nil, multiError
	}
	var ApiResponse ApiResponse

	jsonErr := json.Unmarshal([]byte(body), &ApiResponse)
	if jsonErr != nil {
		multiError = multierror.Append(multiError, jsonErr)
		return nil, multiError
	}

	return &ApiResponse, multiError
}
