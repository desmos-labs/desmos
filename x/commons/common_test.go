package commons_test

import (
	"strconv"
	"testing"

	"github.com/desmos-labs/desmos/x/commons"
	"github.com/stretchr/testify/require"
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

func newStrPtr(value string) *string {
	return &value
}

func TestStringPtrsEqual(t *testing.T) {
	tests := []struct {
		first     *string
		second    *string
		expEquals bool
	}{
		{
			first:     newStrPtr("first"),
			second:    newStrPtr("second"),
			expEquals: false,
		},
		{
			first:     nil,
			second:    newStrPtr("second"),
			expEquals: false,
		},
		{
			first:     newStrPtr("first"),
			second:    nil,
			expEquals: false,
		},
		{
			first:     newStrPtr("first"),
			second:    newStrPtr("first"),
			expEquals: true,
		},
		{
			first:     nil,
			second:    nil,
			expEquals: true,
		},
	}

	for index, test := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			require.Equal(t, test.expEquals, commons.StringPtrsEqual(test.first, test.second))
		})
	}
}
