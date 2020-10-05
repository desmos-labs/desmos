package common_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/commons"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/posts/types/models/common"
)

func TestIsValidPostID(t *testing.T) {
	tests := []struct {
		postID   string
		expValid bool
	}{
		{"", false},
		{"123", false},
		{"d87f7e0c", false},
		{"098f6bcd4621d373cade4e832627b4f6", false},
		{"90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809", false},
		{"ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff", false},
		{"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.postID, func(t *testing.T) {
			require.Equal(t, test.expValid, common.IsValidPostID(test.postID))
		})
	}
}

func TestIsValidSubspace(t *testing.T) {
	tests := []struct {
		subspace string
		expValid bool
	}{
		{"", false},
		{"123", false},
		{"d87f7e0c", false},
		{"098f6bcd4621d373cade4e832627b4f6", false},
		{"90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809", false},
		{"ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff", false},
		{"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", true},
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
		{"like", false},
		{":like", false},
		{"like:", false},
		{":like~:", false},
		{"::", false},
		{": :", false},
		{":\U0001F970:", false},
		{":+1:", true},
		{":-1:", true},
		{":snow_man:", true},
		{":snow-man:", true},
	}

	for _, test := range tests {
		t.Run(test.code, func(t *testing.T) {
			require.Equal(t, test.expValid, common.IsValidReactionCode(test.code))
		})
	}
}
