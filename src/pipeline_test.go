package gocd

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
			"../test-fixtures/pipelines/create_pipeline_success.json",
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
			"../test-fixtures/pipelines/create_pipeline_already_exists.json",
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
			"../test-fixtures/pipelines/create_pipeline_incorrect_material.json",
			4,
			DummyRequestBodyValidator,
			http.StatusUnprocessableEntity))

	defer server.Close()

	_, multiErr := client.CreatePipeline(PipelineConfig{})

	expect1 := "Validations failed for pipeline 'FromTemplate3'. Error(s): [Validation failed.]. Please correct and resubmit."
	assert.Equal(t, expect1, multiErr.Errors[0].Error())

	expect2 := "[Materials][destination] \"The destination directory must be unique across materials.\""
	assert.Equal(t, expect2, multiErr.Errors[2].Error())
}
