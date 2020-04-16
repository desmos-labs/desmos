package v040_test

import (
	"testing"

	v030magpie "github.com/desmos-labs/desmos/x/magpie/legacy/v0.3.0"
	v040magpie "github.com/desmos-labs/desmos/x/magpie/legacy/v0.4.0"
	"github.com/stretchr/testify/require"
)

func TestMigrate(t *testing.T) {
	v020GenesisState := v030magpie.NewGenesisState(
		v030magpie.DefaultSessionLength,
		[]v030magpie.Session{
			{
				SessionID:     1,
				Created:       50,
				Expiry:        58,
				Namespace:     "cosmos",
				ExternalOwner: "cosmos16udxavk6lapwzvpjy9e3f7dxdsdzhnf7aj0q3c",
				PubKey:        "cosmospub1addwnpepqg9hwc5cxnv56hgwruaa69a6mp2ua39hwv94n2qk93tc9652jpju50mhu8s",
				Signature:     "<Bech32>",
			},
			{
				SessionID:     2,
				Created:       15,
				Expiry:        18,
				Namespace:     "desmos",
				ExternalOwner: "desmos18aca5wx2wqadqqwe6mjutsr3ju9dzakhf6n6yf",
				PubKey:        "desmospub1addwnpepqgrk8mj7t4kssr3h0xnl77rwjqk8zdkwurula4ng9g9c07xzzr30wuf0ax2",
				Signature:     "<Bech32>",
			},
		},
	)

	migrated := v040magpie.Migrate(v020GenesisState)
	expected := v040magpie.NewGenesisState(
		v040magpie.DefaultSessionLength,
		[]v040magpie.Session{
			{
				SessionID:     1,
				Created:       50,
				Expiry:        58,
				Namespace:     "cosmos",
				ExternalOwner: "cosmos16udxavk6lapwzvpjy9e3f7dxdsdzhnf7aj0q3c",
				PubKey:        "cosmospub1addwnpepqg9hwc5cxnv56hgwruaa69a6mp2ua39hwv94n2qk93tc9652jpju50mhu8s",
				Signature:     "<Bech32>",
			},
			{
				SessionID:     2,
				Created:       15,
				Expiry:        18,
				Namespace:     "desmos",
				ExternalOwner: "desmos18aca5wx2wqadqqwe6mjutsr3ju9dzakhf6n6yf",
				PubKey:        "desmospub1addwnpepqgrk8mj7t4kssr3h0xnl77rwjqk8zdkwurula4ng9g9c07xzzr30wuf0ax2",
				Signature:     "<Bech32>",
			},
		},
	)

	// Check the session length
	require.Equal(t, expected.DefaultSessionLength, migrated.DefaultSessionLength)

	// Check for sessions
	require.Len(t, migrated.Sessions, len(expected.Sessions))
	for index, session := range migrated.Sessions {
		require.Equal(t, expected.Sessions[index], session)
	}
}
