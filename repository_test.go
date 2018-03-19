package gocd

import (
	"github.com/hashicorp/go-multierror"
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestGetAllRepositories(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/admin/repositories",
		serveFileAsJSONStatusCode(t,
			"GET",
			"test-fixtures/repository/get_all_repositories_success.json",
			1,
			DummyRequestBodyValidator,
			http.StatusOK))

	defer server.Close()

	_, err := client.GetAllRepositories()

	var multiError *multierror.Error
	multiError = nil

	assert.Equal(t, multiError, err)
}

func TestRepository(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/admin/repositories/b83e0ac7-889c-4d28-9f9e-c1fdfae3749f",
		serveFileAsJSONStatusCode(t,
			"GET",
			"test-fixtures/repository/get_repository.json",
			1,
			DummyRequestBodyValidator,
			http.StatusOK))

	defer server.Close()

	_, err := client.GetRepository("b83e0ac7-889c-4d28-9f9e-c1fdfae3749f")

	var multiError *multierror.Error
	multiError = nil

	assert.Equal(t, multiError, err)
}
