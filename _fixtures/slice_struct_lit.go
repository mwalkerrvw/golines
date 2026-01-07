package fixtures

import "net/http"

func testSliceOfStructLiterals() {
	testCases := []struct {
		name       string
		method     string
		path       string
		query      string
		statusCode int
	}{

		// Short entries that fit on one line
		{"GetUser", "GET", "/api/users/123", "", http.StatusOK},
		{"DeleteUser", "DELETE", "/api/users/456", "", http.StatusNoContent},

		// Long entries that need to be split across multiple lines
		{"SearchUsersWithComplexFilterParameters", "GET", "/api/users/search", "name=john&status=active&role=admin", http.StatusOK},
		{"CreateOrganizationMembershipAssignment", "POST", "/api/organizations/org-123/memberships/assignments", "", http.StatusCreated},
		{"UpdateResourcePermissionsConfiguration", "PUT", "/api/resources/res-456/permissions/configuration", "validate=true", http.StatusOK},

		// More short entries
		{"ListItems", "GET", "/api/items", "limit=10", http.StatusOK},
	}

	_ = testCases
}
