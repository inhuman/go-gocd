package go_gocd

import (
	"github.com/hashicorp/go-multierror"
	"encoding/json"
	"strconv"
	"fmt"
	"os"
	jerrparser "github.com/inhuman/go-json-errors-parser"
)

func (c *DefaultClient) GetPipelineStatus(pipelineName string) (*PipelineStatus, *multierror.Error) {
	var multiError *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/pipelines/%s/status", pipelineName))).
		Set("Accept", "application/vnd.go.cd.v"+strconv.Itoa(ApiVersion)+"+json").
		End()

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string(body))
	}

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

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string(body))
	}

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
	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string(body))
	}

	multierror.Append(multiError, errs...)
	if errs != nil {
		return nil, multiError
	}

	if response.StatusCode != 200 {

		parsedErrors := jerrparser.ParseErrors(string(body))
		if parsedErrors.IsErrors() {
			multiError = multierror.Append(multiError, parsedErrors.GetErrors()...)
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

func (c *DefaultClient) PausePipeline(pipelineName, pauseCause string) (*ApiResponse, *multierror.Error) {
	var multiError *multierror.Error

	_, body, errs := c.Request.
		Post(c.resolve(fmt.Sprintf("/go/api/pipelines/%s/pause", pipelineName))).
		Set("Accept", "application/vnd.go.cd.v1+json").
		SendStruct(struct {
		Message string
	}{
		pauseCause,
	}).
		End()

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string(body))
	}

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

func (c *DefaultClient) UnpausePipeline(pipelineName string) (*ApiResponse, *multierror.Error) {
	var multiError *multierror.Error

	_, body, errs := c.Request.
		Post(c.resolve(fmt.Sprintf("/go/api/pipelines/%s/unpause", pipelineName))).
		Set("Accept", "application/vnd.go.cd.v1+json").
		Set("X-GoCD-Confirm", "true").
		End()

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string(body))
	}

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
