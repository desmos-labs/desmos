package keeper_test

import (
	"github.com/desmos-labs/desmos/x/posts/types/models/common"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestValidatePost() {
	id := types.PostID("dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1")
	id2 := types.PostID("e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163")
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	timeZone, err := time.LoadLocation("UTC")
	suite.NoError(err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)

	tests := []struct {
		name     string
		post     types.Post
		expError error
	}{
		{
			name: "Post message cannot be longer than 500 characters",
			post: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        strings.Repeat("a", 550),
				Created:        date,
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
			},
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"post with id dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1 has more than 500 characters"),
		},
		{
			name: "post optional data cannot contain more than 10 key-value",
			post: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "Message",
				Created:        date,
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []common.OptionalDataEntry{
					{"key1", "value"},
					{"key2", "value"},
					{"key3", "value"},
					{"key4", "value"},
					{"key5", "value"},
					{"key6", "value"},
					{"key7", "value"},
					{"key8", "value"},
					{"key9", "value"},
					{"key10", "value"},
					{"key11", "value"},
				},
				Creator: owner,
			},
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"post with id dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1 contains optional data with more than 10 key-value pairs"),
		},
		{
			name: "post optional data values cannot exceed 200 characters",
			post: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "Message",
				Created:        date,
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []common.OptionalDataEntry{
					{"key1", `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque euismod, mi at commodo 
							efficitur, quam sapien congue enim, ut porttitor lacus tellus vitae turpis. Vivamus aliquam 
							sem eget neque metus.`,
					},
				},
				Creator: owner,
			},
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"post with id dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1 has optional data with key key1 which value exceeds 200 characters."),
		},
		{
			name: "Valid post",
			post: types.Post{
				PostID:         id,
				Message:        "Message",
				Created:        date,
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
			},
			expError: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())
			err := keeper.ValidatePost(suite.ctx, suite.keeper, test.post)
			if test.expError != nil {
				suite.Equal(test.expError.Error(), err.Error())
			} else {
				suite.Equal(test.expError, err)
			}
		})
	}
}
