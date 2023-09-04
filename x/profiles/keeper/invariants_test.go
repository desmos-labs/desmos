package keeper_test

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/desmos-labs/desmos/v6/testutil/profilestesting"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

var _ sdk.InvariantRegistry = &MockRegistry{}

type MockRegistry struct {
	RegisteredMap map[string]sdk.Invariant
}

func (m *MockRegistry) RegisterRoute(moduleName, route string, invar sdk.Invariant) {
	m.RegisteredMap[moduleName+route] = invar
}

func (suite *KeeperTestSuite) TestKeeper_RegisterInvariants() {
	mock := &MockRegistry{make(map[string]sdk.Invariant)}
	keeper.RegisterInvariants(mock, suite.k)

	suite.Require().Equal(
		runtime.FuncForPC(reflect.ValueOf(keeper.ValidProfilesInvariant(suite.k)).Pointer()).Name(),
		runtime.FuncForPC(reflect.ValueOf(mock.RegisteredMap[types.ModuleName+"valid-profiles"]).Pointer()).Name(),
	)
	suite.Require().Equal(
		runtime.FuncForPC(reflect.ValueOf(keeper.ValidDTagTransferRequests(suite.k)).Pointer()).Name(),
		runtime.FuncForPC(reflect.ValueOf(mock.RegisteredMap[types.ModuleName+"valid-dtag-transfer-requests"]).Pointer()).Name(),
	)
	suite.Require().Equal(
		runtime.FuncForPC(reflect.ValueOf(keeper.ValidChainLinks(suite.k)).Pointer()).Name(),
		runtime.FuncForPC(reflect.ValueOf(mock.RegisteredMap[types.ModuleName+"valid-chain-links"]).Pointer()).Name(),
	)
	suite.Require().Equal(
		runtime.FuncForPC(reflect.ValueOf(keeper.ValidApplicationLinks(suite.k)).Pointer()).Name(),
		runtime.FuncForPC(reflect.ValueOf(mock.RegisteredMap[types.ModuleName+"valid-application-links"]).Pointer()).Name(),
	)
}

func (suite *KeeperTestSuite) TestInvariants() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		expResponse string
		expBroken   bool
	}{
		{
			name:        "empty state does not break invariants",
			expResponse: "Every invariant condition is fulfilled correctly",
			expBroken:   false,
		},
		{
			name: "ValidProfilesInvariant broken",
			store: func(ctx sdk.Context) {
				profile, err := types.NewProfileFromAccount(
					"",
					profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					time.Now(),
				)
				suite.Require().NoError(err)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
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
			name: "ValidDTagTransferRequests broken",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				request := types.NewDTagTransferRequest("dTag", "sender", "receiver")
				store.Set(
					types.DTagTransferRequestStoreKey(request.Sender, request.Receiver),
					suite.cdc.MustMarshal(&request),
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
				err := suite.cdc.UnmarshalJSON([]byte(pubKey), &any)
				suite.Require().NoError(err)

				var key cryptotypes.PubKey
				err = suite.cdc.UnpackAny(&any, &key)
				suite.Require().NoError(err)

				link := types.NewChainLink(
					"user",
					types.NewBech32Address("value", "prefix"),
					types.NewProof(key, profilestesting.SingleSignatureFromHex("1234"), "value"),
					types.NewChainConfig("chain_name"),
					time.Now(),
				)
				store.Set(
					types.ChainLinksStoreKey("user", "chain_name", "address"),
					suite.cdc.MustMarshal(&link),
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
				err := suite.cdc.UnmarshalJSON([]byte(pubKey), &any)
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
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Now(),
				)
				store.Set(
					types.UserApplicationLinkKey("user", "application", "username"),
					suite.cdc.MustMarshal(&link),
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
