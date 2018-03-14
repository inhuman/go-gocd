package gocd

import (
	"github.com/hashicorp/go-multierror"
	"encoding/json"
	"strconv"
	"fmt"
)

type PipelineStatus struct {
	PausedCause string `json:"pausedCause"`
	PausedBy    string `json:"pausedBy"`
	Paused      bool   `json:"paused"`
	Schedulable bool   `json:"schedulable"`
	Locked      bool   `json:"locked"`
}


func (c *DefaultClient) GetPipelineStatus(pipelineName string) (*PipelineStatus, error) {
	var errors *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/pipelines/%s/status", pipelineName))).
		Set("Accept", "application/vnd.go.cd.v" + strconv.Itoa(ApiVersion) + "+json").
		End()

	if errs != nil {
		errors = multierror.Append(errors, errs...)
		return nil, errors.ErrorOrNil()
	}
	var PipelineStatus PipelineStatus

	jsonErr := json.Unmarshal([]byte(body), &PipelineStatus)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
		return nil, errors.ErrorOrNil()
	}

	return &PipelineStatus, errors.ErrorOrNil()
}






