package gocd

import (
	"testing"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"net/http"
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

	pkg, resp, err := client.CreatePackage(Package{})

	var multiError *multierror.Error
	multiError = nil
	assert.Equal(t, multiError, err)

	var pkgNil *Package
	pkgNil = nil
	assert.Equal(t, pkgNil, pkg)

	assert.Equal(t, "Validations failed for package 'package-id-sdf'. Error(s): [Validation failed.]. Please correct and resubmit.", resp.Message)
	assert.Equal(t, "\"Cannot save package or repo, found duplicate packages. [Repo Name: 'artifactory-rpm', Package Name: 'package_name_2'], [Repo Name: 'artifactory-rpm', Package Name: 'package_name_']\"", string(resp.Data.Errors["id"][0]))

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

	pkg, resp, err := client.CreatePackage(Package{})

	var multiError *multierror.Error
	multiError = nil
	assert.Equal(t, multiError, err)

	var pkgNil *Package
	pkgNil = nil
	assert.Equal(t, pkgNil, pkg)

	assert.Equal(t, "Validations failed for package 'package-id-sdf'. Error(s): [Validation failed.]. Please correct and resubmit.", resp.Message)
	assert.Equal(t, "\"Unsupported key(s) found : PACKAGE_ENV. Allowed key(s) are : PACKAGE_SPEC\"", string(resp.Data.Errors["configuration"][0]))
	assert.Equal(t, "\"Package spec not specified\"", string(resp.Data.Errors["PACKAGE_SPEC"][0]))

}
