package v231_test

import (
	v100 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v100"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/app"
	v200 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v200"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func TestStoreMigration(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	profilesKey := sdk.NewKVStoreKey("profiles")
	ctx := testutil.DefaultContext(profilesKey, sdk.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(profilesKey)

	testCases := []struct {
		name     string
		key      []byte
		oldValue []byte
		newValue []byte
	}{
		{
			name: "Application links migrated correctly",
			key: types.UserApplicationLinkKey(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				"twitter",
				"user",
			),
			oldValue: v200.MustMarshalApplicationLink(cdc, v200.NewApplicationLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				v200.NewData("twitter", "user"),
				v200.AppLinkStateVerificationStarted,
				v200.NewOracleRequest(
					1,
					1,
					v200.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				v200.NewSuccessResult("76616c7565", "signature"), // The value should be HEX
				time.Date(2022, 1, 0, 00, 00, 00, 000, time.UTC),
			)),
			newValue: types.MustMarshalApplicationLink(cdc, types.NewApplicationLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewData("twitter", "user"),
				types.AppLinkStateVerificationStarted,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewSuccessResult("76616c7565", "signature"), // The value should be HEX
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 3, 1, 00, 00, 00, 000, time.UTC),
			)),
		},
		{
			name: "Profiles params are migrated correctly",
			key: v100.Params{
				Nickname: v100.NicknameParams{},
				DTag:     v100.DTagParams{},
				Bio:      v100.BioParams{},
				Oracle:   v100.OracleParams{},
			},
		},
	}
}
