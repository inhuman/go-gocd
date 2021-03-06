package go_gocd

import (
	"testing"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"net/http"
	"fmt"
)

func TestCreatePackageSuccess(t *testing.T) {
	t.Parallel()

	client, server := newTestAPIClient("/go/api/admin/packages",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/package/create_package_success.json",
			1,
			DummyRequestBodyValidator,
			http.StatusOK))

	defer server.Close()

	pkg, resp, err := client.CreatePackage(Package{})

	var multiError *multierror.Error
	multiError = nil
	assert.Equal(t, multiError, err)

	var apiResp *ApiResponse
	apiResp = nil
	assert.Equal(t, apiResp, resp)

	assert.Equal(t, "package_name_4", pkg.Name)
}

func TestCreatePackageAlreadyExists(t *testing.T) {
	t.Parallel()

	client, server := newTestAPIClient("/go/api/admin/packages",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/package/create_package_already_exists.json",
			1,
			DummyRequestBodyValidator,
			http.StatusUnprocessableEntity))

	defer server.Close()

	pkg, _, err := client.CreatePackage(Package{})

	var pkgNil *Package
	pkgNil = nil
	assert.Equal(t, pkgNil, pkg)

	assert.Equal(t, "[] Validations failed for package 'package-id-sdf'. Error(s): [Validation failed.]. Please correct and resubmit.", err.Errors[0].Error())
	assert.Equal(t, "[data][id] Cannot save package or repo, found duplicate packages. [Repo Name: 'artifactory-rpm', Package Name: 'package_name_2'], [Repo Name: 'artifactory-rpm', Package Name: 'package_name_']", err.Errors[1].Error())
}

func TestCreatePackageWrongSpec(t *testing.T) {
	t.Parallel()

	client, server := newTestAPIClient("/go/api/admin/packages",
		serveFileAsJSONStatusCode(t,
			"POST",
			"test-fixtures/package/create_package_wrong_spec.json",
			1,
			DummyRequestBodyValidator,
			http.StatusUnprocessableEntity))

	defer server.Close()

	pkg, _, err := client.CreatePackage(Package{})

	var pkgNil *Package
	pkgNil = nil
	assert.Equal(t, pkgNil, pkg)


	assert.Equal(t, "[] Validations failed for package 'package-id-sdf'. Error(s): [Validation failed.]. Please correct and resubmit.", err.Errors[0].Error())
	assert.Equal(t, "[data][PACKAGE_SPEC] Package spec not specified", err.Errors[1].Error())
	assert.Equal(t, "[data][configuration] Unsupported key(s) found : PACKAGE_ENV. Allowed key(s) are : PACKAGE_SPEC", err.Errors[2].Error())
}

func TestDeletePackageSuccess(t *testing.T) {
	t.Parallel()

	client, server := newTestAPIClient(fmt.Sprintf("/go/api/admin/packages/%s", "package-id-sdf"),
		serveFileAsJSONStatusCode(t,
			"DELETE",
			"test-fixtures/package/delete_package_success.json",
			1,
			DummyRequestBodyValidator,
			http.StatusOK))

	defer server.Close()

	resp, err := client.DeletePackage("package-id-sdf")

	var multiError *multierror.Error
	multiError = nil
	assert.Equal(t, multiError, err)

	assert.Equal(t, "The package definition 'package-id-sdf' was deleted successfully.", resp.Message)
}

func TestDeletePackageNotExists(t *testing.T) {
	t.Parallel()

	client, server := newTestAPIClient(fmt.Sprintf("/go/api/admin/packages/%s", "package-id-sdf"),
		serveFileAsJSONStatusCode(t,
			"DELETE",
			"test-fixtures/package/delete_package_non_exists.json",
			1,
			DummyRequestBodyValidator,
			http.StatusOK))

	defer server.Close()

	resp, err := client.DeletePackage("package-id-sdf")

	var multiError *multierror.Error
	multiError = nil
	assert.Equal(t, multiError, err)

	assert.Equal(t, "Either the resource you requested was not found, or you are not authorized to perform this action.", resp.Message)
}

func TestDeletePackageAssociatedWithPipeline(t *testing.T) {
	t.Parallel()

	client, server := newTestAPIClient(fmt.Sprintf("/go/api/admin/packages/%s", "package-id-sdf"),
		serveFileAsJSONStatusCode(t,
			"DELETE",
			"test-fixtures/package/delete_package_associated_with_pipeline.json",
			1,
			DummyRequestBodyValidator,
			http.StatusOK))

	defer server.Close()

	resp, err := client.DeletePackage("package-id-sdf")

	var multiError *multierror.Error
	multiError = nil
	assert.Equal(t, multiError, err)

	assert.Equal(t, "Cannot delete the package definition 'package-id-sdf' as it is used by pipeline(s): '[fromtemplate]'", resp.Message)
}
