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

type CreatePipelineData struct {
	Group    string   `json:"group"`
	Pipeline Pipeline `json:"pipeline"`
}

type CreatePipelineResponse struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Doc struct {
			Href string `json:"href"`
		} `json:"doc"`
		Find struct {
			Href string `json:"href"`
		} `json:"find"`
	} `json:"_links"`
	LabelTemplate string      `json:"label_template"`
	LockBehavior  string      `json:"lock_behavior"`
	Name          string      `json:"name"`
	Template      interface{} `json:"template"`
	Origin struct {
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Doc struct {
				Href string `json:"href"`
			} `json:"doc"`
		} `json:"_links"`
		Type string `json:"type"`
	} `json:"origin"`
	Parameters           []interface{} `json:"parameters"`
	EnvironmentVariables []interface{} `json:"environment_variables"`
	Materials            []Material    `json:"materials"`
	Stages               []Stage       `json:"stages"`
	TrackingTool         interface{}   `json:"tracking_tool"`
	Timer                interface{}   `json:"timer"`
}

type Pipeline struct {
	LabelTemplate string      `json:"label_template"`
	LockBehavior  string      `json:"lock_behavior"`
	Name          string      `json:"name"`
	Template      interface{} `json:"template"`
	Materials     []Material  `json:"materials"`
	Stages        []Stage     `json:"stages"`
}

type Stage struct {
	Name                  string `json:"name"`
	FetchMaterials        bool   `json:"fetch_materials"`
	CleanWorkingDirectory bool   `json:"clean_working_directory"`
	NeverCleanupArtifacts bool   `json:"never_cleanup_artifacts"`
	Approval struct {
		Type string `json:"type"`
		Authorization struct {
			Roles []interface{} `json:"roles"`
			Users []interface{} `json:"users"`
		} `json:"authorization"`
	} `json:"approval"`
	EnvironmentVariables []interface{} `json:"environment_variables"`
	Jobs []struct {
		Name                 string        `json:"name"`
		RunInstanceCount     interface{}   `json:"run_instance_count"`
		Timeout              int           `json:"timeout"`
		EnvironmentVariables []interface{} `json:"environment_variables"`
		Resources            []interface{} `json:"resources"`
		Tasks []struct {
			Type string `json:"type"`
			Attributes struct {
				RunIf            []string    `json:"run_if"`
				Command          string      `json:"command"`
				WorkingDirectory interface{} `json:"working_directory"`
			} `json:"attributes"`
		} `json:"tasks"`
	} `json:"jobs"`
}

type Material struct {
	Type string `json:"type"`
	Attributes struct {
		URL             string      `json:"url"`
		Destination     string      `json:"destination"`
		Filter          interface{} `json:"filter"`
		InvertFilter    bool        `json:"invert_filter"`
		Name            interface{} `json:"name"`
		AutoUpdate      bool        `json:"auto_update"`
		Branch          string      `json:"branch"`
		SubmoduleFolder interface{} `json:"submodule_folder"`
		ShallowClone    bool        `json:"shallow_clone"`
	} `json:"attributes"`
}

func (c *DefaultClient) GetPipelineStatus(pipelineName string) (*PipelineStatus, error) {
	var errors *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/pipelines/%s/status", pipelineName))).
		Set("Accept", "application/vnd.go.cd.v"+strconv.Itoa(ApiVersion)+"+json").
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

func (c *DefaultClient) CreatePipeline(pipelineData CreatePipelineData) (*CreatePipelineResponse, error) {
	var errors *multierror.Error

	_, body, errs := c.Request.
		Post(c.resolve("/go/api/admin/pipelines")).
		Set("Accept", "application/vnd.go.cd.v"+strconv.Itoa(ApiVersion)+"+json").
		SendStruct(pipelineData).
		End()
	multierror.Append(errors, errs...)
	if errs != nil {
		return nil, errors.ErrorOrNil()
	}

	var resp CreatePipelineResponse

	jsonErr := json.Unmarshal([]byte(body), &resp)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
		return nil, errors.ErrorOrNil()
	}

	return &resp, nil
}
