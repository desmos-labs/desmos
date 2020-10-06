package common_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/posts/types/models/common"
	"github.com/stretchr/testify/assert"
)

func TestNewOptionalData(t *testing.T) {
	expOpd := common.OptionalData{
		Key:   "key",
		Value: "value",
	}

	opd := common.NewOptionalData("key", "value")

	assert.Equal(t, expOpd, opd)
}

func TestOptionalData_Equals(t *testing.T) {
	tests := []struct {
		name         string
		optionalData common.OptionalData
		otherOpData  common.OptionalData
		expBool      bool
	}{
		{
			name:         "Different optional data returns false",
			optionalData: common.NewOptionalData("key", "value"),
			otherOpData:  common.NewOptionalData("key", "val"),
			expBool:      false,
		},
		{
			name:         "Same optional data returns true",
			optionalData: common.NewOptionalData("key", "value"),
			otherOpData:  common.NewOptionalData("key", "value"),
			expBool:      true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expBool, test.optionalData.Equals(test.otherOpData))
		})
	}
}

func TestOptionalData_String(t *testing.T) {
	opt := common.NewOptionalData("optional", "data")
	assert.Equal(t, "[Key] [Value]\n[optional] [data]", opt.String())
}
