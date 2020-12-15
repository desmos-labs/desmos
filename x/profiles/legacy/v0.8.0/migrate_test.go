package v080_test

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v060 "github.com/desmos-labs/desmos/x/profiles/legacy/v0.6.0"
	v080 "github.com/desmos-labs/desmos/x/profiles/legacy/v0.8.0"
)

// newStrPtr returns a new string pointer
func newStrPtr(value string) *string {
	return &value
}

func TestMigrate080(t *testing.T) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	content, err := ioutil.ReadFile("v060state.json")
	require.NoError(t, err)

	var v060state v060.GenesisState
	err = json.Unmarshal(content, &v060state)
	require.NoError(t, err)

	genesisTime, err := time.Parse(time.RFC3339, "2020-01-01T15:00:00Z")
	require.NoError(t, err)

	v080state := v080.Migrate(v060state, genesisTime)

	// make sure that all profiles are migrated
	require.Equal(t, len(v080state.Profiles), len(v060state.Profiles))

	// make sure that params are properly set
	params := v080.Params{
		MonikerParams: v080.MonikerParams{
			MinMonikerLen: sdk.NewInt(2),
			MaxMonikerLen: sdk.NewInt(1000),
		},
		DtagParams: v080.DtagParams{
			RegEx:      `^[A-Za-z0-9_]+$`,
			MinDtagLen: sdk.NewInt(3),
			MaxDtagLen: sdk.NewInt(30),
		},
		MaxBioLen: sdk.NewInt(1000),
	}

	require.Equal(t, params, v080state.Params)
}

func TestConvertProfiles(t *testing.T) {
	user1, err := sdk.AccAddressFromBech32("desmos1rzgjtjudpn4ppj0g7lj7vvcd2cxsjcvvyyatkw")
	require.NoError(t, err)

	user2, err := sdk.AccAddressFromBech32("desmos18nx27z75wp8xk9j2t73hl6ggpnwpgqa8a423vg")
	require.NoError(t, err)

	genesisTime, err := time.Parse(time.RFC3339, "2020-01-01T15:00:00Z")
	require.NoError(t, err)

	profiles := v080.ConvertProfiles([]v060.Profile{
		{
			Moniker: "Leo DC",
			Name:    newStrPtr("Leonardo"),
			Surname: newStrPtr("Di Caprio"),
			Bio:     newStrPtr("Actor"),
			Pictures: &v060.Pictures{
				Profile: newStrPtr("https://text.com/first-profile-pic"),
				Cover:   newStrPtr("https://text.com/first-cover-pic"),
			},
			Creator: user1,
		},
		{
			Moniker: "Mon Bel",
			Name:    newStrPtr("Monica"),
			Surname: newStrPtr("Bellucci"),
			Bio:     newStrPtr("Actress"),
			Pictures: &v060.Pictures{
				Profile: newStrPtr("https://text.com/second-profile-pic"),
				Cover:   newStrPtr("https://text.com/second-cover-pic"),
			},
			Creator: user2,
		},
	}, genesisTime)

	expected := []v080.Profile{
		{
			DTag:    "Leo_DC",
			Moniker: newStrPtr("Leonardo Di Caprio"),
			Bio:     newStrPtr("Actor"),
			Pictures: &v080.Pictures{
				Profile: newStrPtr("https://text.com/first-profile-pic"),
				Cover:   newStrPtr("https://text.com/first-cover-pic"),
			},
			Creator:      user1,
			CreationDate: genesisTime,
		},
		{
			DTag:    "Mon_Bel",
			Moniker: newStrPtr("Monica Bellucci"),
			Bio:     newStrPtr("Actress"),
			Pictures: &v080.Pictures{
				Profile: newStrPtr("https://text.com/second-profile-pic"),
				Cover:   newStrPtr("https://text.com/second-cover-pic"),
			},
			Creator:      user2,
			CreationDate: genesisTime,
		},
	}

	require.Len(t, profiles, len(expected))
	for index, profile := range profiles {
		expectedProfile := expected[index]
		require.Equal(t, profile.DTag, expectedProfile.DTag)
		require.Equal(t, profile.Moniker, expectedProfile.Moniker)
		require.Equal(t, profile.Bio, expectedProfile.Bio)
		require.Equal(t, profile.Pictures, expectedProfile.Pictures)
		require.Equal(t, profile.Creator, expectedProfile.Creator)
		require.True(t, profile.CreationDate.Equal(expectedProfile.CreationDate))
	}
}

func TestGetProfileDtag(t *testing.T) {
	tests := []struct {
		moniker string
		expDTag string
	}{
		{moniker: "John Doe", expDTag: "John_Doe"},
		{moniker: "JDoe", expDTag: "JDoe"},
	}

	for index, test := range tests {
		test := test
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			require.Equal(t, test.expDTag, v080.GetProfileDTag(test.moniker))
		})
	}
}

func TestGetProfileMoniker(t *testing.T) {
	tests := []struct {
		name       string
		surname    string
		expMoniker string
	}{
		{name: "John", expMoniker: "John"},
		{surname: "Doe", expMoniker: "Doe"},
		{name: "John", surname: "Doe", expMoniker: "John Doe"},
		{name: "", surname: "", expMoniker: ""},
	}

	for index, test := range tests {
		test := test
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			value := v080.GetProfileMoniker(&test.name, &test.surname)

			if len(test.expMoniker) > 0 {
				require.Equal(t, test.expMoniker, *value)
			} else {
				require.Nil(t, value)
			}
		})
	}
}
