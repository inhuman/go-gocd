package gocd

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestCreatePipelineAlreadyExists(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/admin/pipelines",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/create_pipeline_already_exists.json",
			4,
			DummyRequestBodyValidator,
			http.StatusUnprocessableEntity))

	defer server.Close()

	_, err := client.CreatePipeline(CreatePipelineData{})
	assert.Error(t, err, "Failed to add pipeline. The pipeline 'double_pipeline' already exists.")
}
