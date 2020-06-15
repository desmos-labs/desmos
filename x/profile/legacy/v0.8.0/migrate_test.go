package v080_test

import (
	"strconv"
	"testing"

	v080 "github.com/desmos-labs/desmos/x/profile/legacy/v0.8.0"
	"github.com/stretchr/testify/require"
)

func TestGetProfileDtag(t *testing.T) {
	tests := []struct {
		moniker string
		expDTag string
	}{
		{moniker: "John Doe", expDTag: "JohnDoe"},
		{moniker: "JDoe", expDTag: "JDoe"},
	}

	for index, test := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			require.Equal(t, test.expDTag, v080.GetProfileDTag(test.moniker))
		})
	}
}

func TestGetProfileMoniker(t *testing.T) {
	tests := []struct {
		name       string
		surname    string
		expMoniker string
	}{
		{name: "John", expMoniker: "John"},
		{surname: "Doe", expMoniker: "Doe"},
		{name: "John", surname: "Doe", expMoniker: "John Doe"},
		{name: "", surname: "", expMoniker: ""},
	}

	for index, test := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			value := v080.GetProfileMoniker(&test.name, &test.surname)

			if len(test.expMoniker) > 0 {
				require.Equal(t, test.expMoniker, *value)
			} else {
				require.Nil(t, value)
			}
		})
	}
}
