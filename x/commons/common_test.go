package commons_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/x/commons"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		value []string
		exp   []string
	}{
		{
			value: []string{"1", "2", "3"},
			exp:   []string{"1", "2", "3"},
		},
		{
			value: []string{"1", "2", "3", "2", "3", "1"},
			exp:   []string{"1", "2", "3"},
		},
		{
			value: []string{"1", "2", "3", "1", "1", "1"},
			exp:   []string{"1", "2", "3"},
		},
	}

	for index, test := range tests {
		test := test
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			res := commons.Unique(test.value)
			require.Len(t, res, len(test.exp))

			for index, value := range res {
				require.Equal(t, value, test.exp[index])
			}
		})
	}
}

func TestIsURIValid(t *testing.T) {
	tests := []struct {
		uri      string
		expValid bool
	}{

		{
			uri:      "http://",
			expValid: false,
		},
		{
			uri:      "error.com",
			expValid: false,
		},
		{
			uri:      ".com",
			expValid: false,
		},
		{
			uri:      "ttps://",
			expValid: false,
		},
		{
			uri:      "ps://site.com",
			expValid: false,
		},
		{
			uri:      "https://",
			expValid: false,
		},
		{
			uri:      "https://example.com",
			expValid: true,
		},
		{
			uri:      "http://error.com",
			expValid: true,
		},
		{
			// This test refers to this issue: https://github.com/desmos-labs/desmos/v5/issues/233
			// It has been included to avoid regressions from being ever introduced about it
			uri:      "https://timgsa.baidu.com/timg?\\n\\nimage&quality=80&size=b9999_10000&sec=1594915557404&di=70d5872ec070ce3d22c7f2f11f10d7ff&imgtype=0&src=http%3A%2F%2Fa2.att.hudong.com%2F36%2F48%2F19300001357258133412489354717.jpg",
			expValid: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.uri, func(t *testing.T) {
			require.Equal(t, test.expValid, commons.IsURIValid(test.uri))
		})
	}
}
