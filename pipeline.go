package go_gocd

import (
	"github.com/hashicorp/go-multierror"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/pkg/errors"
	"os"
)


func (c *DefaultClient) GetPipelineStatus(pipelineName string) (*PipelineStatus, *multierror.Error) {
	var multiError *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/pipelines/%s/status", pipelineName))).
		Set("Accept", "application/vnd.go.cd.v"+strconv.Itoa(ApiVersion)+"+json").
		End()

	if errs != nil {
		multiError = multierror.Append(multiError, errs...)
		return nil, multiError
	}
	var PipelineStatus PipelineStatus

	jsonErr := json.Unmarshal([]byte(body), &PipelineStatus)
	if jsonErr != nil {
		multiError = multierror.Append(multiError, jsonErr)
		return nil, multiError
	}

	return &PipelineStatus, multiError
}

func (c *DefaultClient) DeletePipeline(pipelineName string) (*ApiResponse, *multierror.Error) {
	var multiError *multierror.Error

	_, body, errs := c.Request.
		Delete(c.resolve(fmt.Sprintf("/go/api/admin/pipelines/%s", pipelineName))).
		Set("Accept", "application/vnd.go.cd.v"+strconv.Itoa(ApiVersion)+"+json").
		End()

	if errs != nil {
		multiError = multierror.Append(multiError, errs...)
		return nil, multiError
	}
	var apiResponse ApiResponse

	jsonErr := json.Unmarshal([]byte(body), &apiResponse)
	if jsonErr != nil {
		multiError = multierror.Append(multiError, jsonErr)
		return nil, multiError
	}

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string(body))
	}

	return &apiResponse, multiError
}

func (c *DefaultClient) CreatePipeline(pipelineData PipelineConfig) (*ApiResponse, *multierror.Error) {
	var multiError *multierror.Error

	response, body, errs := c.Request.
		Post(c.resolve("/go/api/admin/pipelines")).
		Set("Accept", "application/vnd.go.cd.v"+strconv.Itoa(ApiVersion)+"+json").
		SendStruct(pipelineData).
		End()

	multierror.Append(multiError, errs...)
	if errs != nil {
		return nil, multiError
	}

	var apiResponse ApiResponse

	if response.StatusCode != 200 {

		err := json.Unmarshal([]byte(body), &apiResponse)
		if err != nil {
			multiError = multierror.Append(multiError, err)
			return nil, multiError
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

		// Check material pipeline errors
		for _, mat := range apiResponse.Data.Materials {
			if len(mat.Errors) > 0 {
				for fieldName, respErrArr := range mat.Errors {
					for _, respErr := range respErrArr {
						multiError = multierror.Append(
							multiError, errors.New("[Materials]["+fieldName+"] "+string(respErr)))
					}
				}
			}
		}

		// Check environment variables pipeline errors
		for _, mat := range apiResponse.Data.EnvironmentVariables {
			if len(mat.Errors.ValueForDisplay) > 0 {
				for _, respErr := range mat.Errors.ValueForDisplay {
					multiError = multierror.Append(
						multiError, errors.New("[EnvironmentVariables] "+string(respErr)))

				}
			}
		}


		if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
			fmt.Println(string(body))
		}

		return nil, multiError
	}

	var resp ApiResponse

	jsonErr := json.Unmarshal([]byte(body), &resp)
	if jsonErr != nil {
		multiError = multierror.Append(multiError, jsonErr)
		return nil, multiError
	}

	return &resp, nil
}
