package keeper_test

import (
	"fmt"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	address, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.Require().NoError(err)

	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		expResponse string
		expBroken   bool
	}{
		{
			name:        "Empty state does not break invariants",
			expResponse: "Every invariant condition is fulfilled correctly",
			expBroken:   false,
		},
		{
			name: "ValidProfilesInvariant broken",
			store: func(ctx sdk.Context) {
				profile, err := types.NewProfileFromAccount(
					"",
					authtypes.NewBaseAccountWithAddress(address),
					time.Now(),
				)
				suite.Require().NoError(err)
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid profiles",
				fmt.Sprintf("%s%s",
					"The following list contains invalid profiles:\n",
					"[DTag]: , [Creator]: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47\n",
				),
			),
			expBroken: true,
		},
		{
			name: "ValidUserBlocksInvariant broken",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				block := types.NewUserBlock("blocker", "blocked", "reason", "subspace")
				store.Set(
					types.UserBlockStoreKey(block.Blocker, block.Subspace, block.Blocked),
					suite.cdc.MustMarshalBinaryBare(&block),
				)
			},
			expBroken: true,
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid user blocks",
				fmt.Sprintf("%s%s",
					"The following list contains invalid user blocks:\n",
					"[Blocker]: blocker, [Blocked]: blocked, [Subspace]: subspace\n",
				),
			),
		},
		{
			name: "ValidRelationshipsInvariant broken",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				relationship := types.NewRelationship("creator", "recipient", "subspace")
				store.Set(
					types.RelationshipsStoreKey(relationship.Creator, relationship.Subspace, relationship.Recipient),
					suite.cdc.MustMarshalBinaryBare(&relationship),
				)
			},
			expBroken: true,
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid relationships",
				fmt.Sprintf("%s%s",
					"The following list contains invalid relationships:\n",
					"[Creator]: creator, [Recipient]: recipient, [Subspace]: subspace\n",
				),
			),
		},
		{
			name: "ValidDTagTransferRequests broken",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				request := types.NewDTagTransferRequest("dTag", "sender", "receiver")
				store.Set(
					types.DTagTransferRequestStoreKey(request.Sender, request.Receiver),
					suite.cdc.MustMarshalBinaryBare(&request),
				)
			},
			expBroken: true,
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid dtag transfer requests",
				fmt.Sprintf("%s%s",
					"The following list contains invalid DTag transfer requests:\n",
					"[Sender]: sender, [Receiver]: receiver\n",
				),
			),
		},
		{
			name: "ValidChainLinks broken",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				pubKey := `{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A6jN4EPjj8mHf722yjEaKaGdJpxnTR40pDvXlX1mni9C"}`

				var any codectypes.Any
				err = suite.cdc.UnmarshalJSON([]byte(pubKey), &any)
				suite.Require().NoError(err)

				var key cryptotypes.PubKey
				err = suite.cdc.UnpackAny(&any, &key)
				suite.Require().NoError(err)

				link := types.NewChainLink(
					"user",
					types.NewBech32Address("value", "prefix"),
					types.NewProof(key, "signature", "value"),
					types.NewChainConfig("chain_name"),
					time.Now(),
				)
				store.Set(
					types.ChainLinksStoreKey("user", "chain_name", "address"),
					suite.cdc.MustMarshalBinaryBare(&link),
				)
			},
			expBroken: true,
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid chain links",
				fmt.Sprintf("%s%s",
					"The following list contains invalid chain links:\n",
					"[User]: user, [Chain]: chain_name, [Address]: value\n",
				),
			),
		},
		{
			name: "ValidApplicationLinks broken",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				pubKey := `{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A6jN4EPjj8mHf722yjEaKaGdJpxnTR40pDvXlX1mni9C"}`

				var any codectypes.Any
				err = suite.cdc.UnmarshalJSON([]byte(pubKey), &any)
				suite.Require().NoError(err)

				var key cryptotypes.PubKey
				err = suite.cdc.UnpackAny(&any, &key)
				suite.Require().NoError(err)

				link := types.NewApplicationLink(
					"user",
					types.NewData("application", "username"),
					types.AppLinkStateVerificationStarted,
					types.NewOracleRequest(1, 1, types.NewOracleRequestCallData("", ""), "client_id"),
					nil,
					time.Now(),
				)
				store.Set(
					types.UserApplicationLinkKey("user", "application", "username"),
					suite.cdc.MustMarshalBinaryBare(&link),
				)
			},
			expBroken: true,
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid application links",
				fmt.Sprintf("%s%s",
					"The following list contains invalid application links:\n",
					"[User]: user, [Application]: application, [Username]: username\n",
				),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, broken := keeper.AllInvariants(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
			suite.Require().Equal(tc.expResponse, res)
		})
	}
}
