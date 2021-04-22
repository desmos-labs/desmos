package keeper_test

import (
	"strings"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func (suite *KeeperTestSuite) TestValidatePost() {
	tests := []struct {
		name     string
		post     types.Post
		expError error
	}{
		{
			name: "Message cannot be longer than 500 characters",
			post: types.Post{
				PostId:         "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				ParentId:       "e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				Message:        strings.Repeat("a", 550),
				Created:        time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"post with id dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1 has more than 500 characters"),
		},
		{
			name: "Optional data cannot contain more than 10 key-value",
			post: types.Post{
				PostId:         "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				ParentId:       "e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				Message:        "Message",
				Created:        time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []types.OptionalDataEntry{
					types.NewOptionalDataEntry("key1", "value"),
					types.NewOptionalDataEntry("key2", "value"),
					types.NewOptionalDataEntry("key3", "value"),
					types.NewOptionalDataEntry("key4", "value"),
					types.NewOptionalDataEntry("key5", "value"),
					types.NewOptionalDataEntry("key6", "value"),
					types.NewOptionalDataEntry("key7", "value"),
					types.NewOptionalDataEntry("key8", "value"),
					types.NewOptionalDataEntry("key9", "value"),
					types.NewOptionalDataEntry("key10", "value"),
					types.NewOptionalDataEntry("key11", "value"),
				},
				Creator: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"post with id dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1 contains optional data with more than 10 key-value pairs"),
		},
		{
			name: "Optional data values cannot exceed 200 characters",
			post: types.Post{
				PostId:         "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				ParentId:       "e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				Message:        "Message",
				Created:        time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []types.OptionalDataEntry{
					types.NewOptionalDataEntry("key1",
						`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque euismod, mi at commodo 
							efficitur, quam sapien congue enim, ut porttitor lacus tellus vitae turpis. Vivamus aliquam 
							sem eget neque metus.`,
					),
				},
				Creator: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"post with id dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1 has optional data with key key1 which value exceeds 200 characters."),
		},
		{
			name: "Valid post",
			post: types.Post{
				PostId:         "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				Message:        "Message",
				Created:        time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.k.SetParams(suite.ctx, types.DefaultParams())
			err := suite.k.ValidatePost(suite.ctx, test.post)

			if test.expError != nil {
				suite.Require().Error(err)
				suite.Require().Equal(test.expError.Error(), err.Error())
			} else {
				suite.Require().Equal(test.expError, err)
			}
		})
	}
}
