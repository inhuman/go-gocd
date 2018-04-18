package go_gocd

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/hashicorp/go-multierror"
)

func TestCreatePipelineSuccess(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/admin/pipelines",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/pipelines/create_pipeline_success.json",
			4,
			DummyRequestBodyValidator,
			http.StatusOK))

	defer server.Close()

	_, err := client.CreatePipeline(PipelineConfig{})

	var multiError *multierror.Error
	multiError = nil

	assert.Equal(t, multiError, err)
}

func TestCreatePipelineAlreadyExists(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/admin/pipelines",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/pipelines/create_pipeline_already_exists.json",
			4,
			DummyRequestBodyValidator,
			http.StatusUnprocessableEntity))

	defer server.Close()

	_, err := client.CreatePipeline(PipelineConfig{})
	assert.Error(t, err, "Failed to add pipeline. The pipeline 'double_pipeline' already exists.")
}

func TestCreatePipelineIncorrectMaterial(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/admin/pipelines",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/pipelines/create_pipeline_incorrect_material.json",
			4,
			DummyRequestBodyValidator,
			http.StatusUnprocessableEntity))

	defer server.Close()

	_, multiErr := client.CreatePipeline(PipelineConfig{})

	expect1 := "[] Validations failed for pipeline 'FromTemplate3'. Error(s): [Validation failed.]. Please correct and resubmit."
	assert.Equal(t, expect1, multiErr.Errors[0].Error())

	expect2 := "[materials][destination] Invalid Destination Directory. Every material needs a different destination directory and the directories should not be nested."
	assert.Equal(t, expect2, multiErr.Errors[2].Error())
}

func TestPausePipeline(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/pipelines/pipeline1/pause",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/pipelines/pause_pipeline.json",
			1,
			DummyRequestBodyValidator,
			http.StatusOK))

	defer server.Close()

	resp, err := client.PausePipeline("pipeline1", "just pause")

	assert.Equal(t, "Pipeline 'pipeline1' paused successfully.", resp.Message)
	var multiError *multierror.Error
	multiError = nil

	assert.Equal(t, multiError, err)
}

func TestPausePipelineNotExists(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/pipelines/pipeline1/pause",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/pipelines/pause_pipeline_not_exists.json",
			1,
			DummyRequestBodyValidator,
			http.StatusNotFound))

	defer server.Close()

	_, err := client.PausePipeline("pipeline1", "just pause")

	assert.Equal(t, "pipeline 'pipeline1' not found.", err.Errors[0].Error())
}


func TestUnpausePipeline(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/pipelines/pipeline1/unpause",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/pipelines/unpause_pipeline.json",
			1,
			DummyRequestBodyValidator,
			http.StatusOK))

	defer server.Close()

	resp, err := client.UnpausePipeline("pipeline1")

	assert.Equal(t, "Pipeline 'pipeline1' unpaused successfully.", resp.Message)
	var multiError *multierror.Error
	multiError = nil

	assert.Equal(t, multiError, err)
}
