package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/posts/types"

	"github.com/desmos-labs/desmos/x/commons"

	"github.com/stretchr/testify/require"
)

func TestIsValidPostID(t *testing.T) {
	tests := []struct {
		postID   string
		expValid bool
	}{
		{
			postID:   "",
			expValid: false,
		},
		{
			postID:   "123",
			expValid: false,
		},
		{
			postID:   "d87f7e0c",
			expValid: false,
		},
		{
			postID:   "098f6bcd4621d373cade4e832627b4f6",
			expValid: false,
		},
		{
			postID:   "90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809",
			expValid: false,
		},
		{
			postID:   "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff",
			expValid: false,
		},
		{
			postID:   "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			expValid: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.postID, func(t *testing.T) {
			require.Equal(t, test.expValid, types.IsValidPostID(test.postID))
		})
	}
}

func TestIsValidSubspace(t *testing.T) {
	tests := []struct {
		subspace string
		expValid bool
	}{
		{
			subspace: "",
			expValid: false,
		},
		{
			subspace: "123",
			expValid: false,
		},
		{
			subspace: "d87f7e0c",
			expValid: false,
		},
		{
			subspace: "098f6bcd4621d373cade4e832627b4f6",
			expValid: false,
		},
		{
			subspace: "90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809",
			expValid: false,
		},
		{
			subspace: "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff",
			expValid: false,
		},
		{
			subspace: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			expValid: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.subspace, func(t *testing.T) {
			require.Equal(t, test.expValid, commons.IsValidSubspace(test.subspace))
		})
	}
}

func TestIsValidReactionCode(t *testing.T) {
	tests := []struct {
		code     string
		expValid bool
	}{
		{
			code:     "like",
			expValid: false,
		},
		{
			code:     ":like",
			expValid: false,
		},
		{
			code:     "like:",
			expValid: false,
		},
		{
			code:     ":like~:",
			expValid: false,
		},
		{
			code:     "::",
			expValid: false,
		},
		{
			code:     ": :",
			expValid: false,
		},
		{
			code:     ":\U0001F970:",
			expValid: false,
		},
		{
			code:     ":+1:",
			expValid: true,
		},
		{
			code:     ":-1:",
			expValid: true,
		},
		{
			code:     ":snow_man:",
			expValid: true,
		},
		{
			code:     ":snow-man:",
			expValid: true,
		},
	}

	for _, test := range tests {
		t.Run(test.code, func(t *testing.T) {
			require.Equal(t, test.expValid, types.IsValidReactionCode(test.code))
		})
	}
}
