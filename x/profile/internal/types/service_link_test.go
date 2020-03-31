package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestServiceLink_Validate(t *testing.T) {
	tests := []struct {
		name        string
		serviceLink types.ServiceLink
		expErr      error
	}{
		{
			name:        "Service link name empty returns error",
			serviceLink: types.NewServiceLink("", "credential", "proof"),
			expErr:      fmt.Errorf("name of the trusted service cannot be empty or blank"),
		},
		{
			name:        "Service link credential empty returns error",
			serviceLink: types.NewServiceLink("service", "", "proof"),
			expErr:      fmt.Errorf("credential of service service cannot be empty or blank"),
		},
		{
			name:        "Service link proof empty returns error",
			serviceLink: types.NewServiceLink("service", "credential", ""),
			expErr:      fmt.Errorf("service service proof cannot be empty or blank"),
		},
		{
			name:        "Valid service link returns no error",
			serviceLink: types.NewServiceLink("service", "credential", "proof"),
			expErr:      nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := test.serviceLink.Validate()
			require.Equal(t, test.expErr, actual)
		})
	}

}

func TestServiceLink_Equals(t *testing.T) {
	tests := []struct {
		name        string
		serviceLink types.ServiceLink
		otherSl     types.ServiceLink
		expBool     bool
	}{
		{
			name:        "Equals services link returns true",
			serviceLink: types.NewServiceLink("service", "credential", "proof"),
			otherSl:     types.NewServiceLink("service", "credential", "proof"),
			expBool:     true,
		},
		{
			name:        "Non equals services link returns false",
			serviceLink: types.NewServiceLink("service", "credential", "proof"),
			otherSl:     types.NewServiceLink("service", "credential", "pro"),
			expBool:     false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := test.serviceLink.Equals(test.otherSl)
			require.Equal(t, test.expBool, actual)
		})
	}
}
