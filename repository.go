package gocd

import (
	"github.com/hashicorp/go-multierror"
	"fmt"
	"strconv"
	"encoding/json"
)

func (c *DefaultClient) GetAllRepositories(pipelineName string) (*PipelineStatus, error) {
	var multiError *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/pipelines/%s/status", pipelineName))).
		Set("Accept", "application/vnd.go.cd.v"+strconv.Itoa(ApiVersion)+"+json").
		End()

	if errs != nil {
		multiError = multierror.Append(multiError, errs...)
		return nil, multiError.ErrorOrNil()
	}
	var PipelineStatus PipelineStatus

	jsonErr := json.Unmarshal([]byte(body), &PipelineStatus)
	if jsonErr != nil {
		multiError = multierror.Append(multiError, jsonErr)
		return nil, multiError.ErrorOrNil()
	}

	return &PipelineStatus, multiError.ErrorOrNil()
}