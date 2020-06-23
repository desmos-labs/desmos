package v080

import (
	"strings"
	"time"

	v060profile "github.com/desmos-labs/desmos/x/profile/legacy/v0.6.0"
)

// Migrate accepts an exported v0.6.0 profile genesis state and migrates it
// to a v0.8.0 profile genesis state.
func Migrate(oldGenState v060profile.GenesisState, genesisTime time.Time) GenesisState {
	return GenesisState{
		Profiles: ConvertProfiles(oldGenState.Profiles, genesisTime),
	}
}

// ConvertProfiles take a list of v0.6.0 profiles and converts them to a list
// of v0.8.0 profiles. To do so it get rids of all the name and surname fields and
// moves the value of the moniker field to the dtag one. It also sets the creation
// date to the genesis time given.
func ConvertProfiles(oldProfiles []v060profile.Profile, genesisTime time.Time) []Profile {
	profiles := make([]Profile, len(oldProfiles))
	for index, profile := range oldProfiles {
		// Get the pictures
		var pictures *Pictures = nil
		if profile.Pictures != nil {
			pics := Pictures(*profile.Pictures)
			pictures = &pics
		}

		// Build the new profile
		profiles[index] = Profile{
			DTag:         GetProfileDTag(profile.Moniker),
			Moniker:      GetProfileMoniker(profile.Name, profile.Surname),
			Bio:          profile.Bio,
			Pictures:     pictures,
			Creator:      profile.Creator,
			CreationDate: genesisTime,
		}
	}

	return profiles
}

// GetProfileDTag returns the Dtag for the given profile. To do so, it takes the
// current profile moniker and remove all the whitespaces from it.
func GetProfileDTag(moniker string) string {
	return strings.ReplaceAll(moniker, " ", "")
}

// GetProfileMoniker returns the moniker for the given profile.
// To do so, it uses the name and surname currently set, joining them using
// a whitespace as the separator.
func GetProfileMoniker(name, surname *string) *string {
	var monikerParts []string
	if name != nil && len(*name) > 0 {
		monikerParts = append(monikerParts, *name)
	}
	if surname != nil && len(*surname) > 0 {
		monikerParts = append(monikerParts, *surname)
	}

	var moniker *string
	if len(monikerParts) > 0 {
		value := strings.Join(monikerParts, " ")
		moniker = &value
	}
	return moniker
}
