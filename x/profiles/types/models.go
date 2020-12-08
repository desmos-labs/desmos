package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewPictures is a constructor function for Pictures
func NewPictures(profile, cover string) Pictures {
	return Pictures{
		Profile: profile,
		Cover:   cover,
	}
}

// Validate check the validity of the Pictures
func (pic Pictures) Validate() error {
	if pic.Profile != "" {
		valid := commons.IsURIValid(pic.Profile)
		if !valid {
			return fmt.Errorf("invalid profile picture uri provided")
		}
	}

	if pic.Cover != "" {
		valid := commons.IsURIValid(pic.Cover)
		if !valid {
			return fmt.Errorf("invalid profile cover uri provided")
		}
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewProfile builds a new profile having the given dtag, creator and creation date
func NewProfile(dtag string, moniker, bio string, pictures Pictures, creationDate time.Time, creator string) Profile {
	return Profile{
		Dtag:         dtag,
		Moniker:      moniker,
		Bio:          bio,
		Pictures:     pictures,
		Creator:      creator,
		CreationDate: creationDate,
	}
}

// Update updates the fields of a given profile. An error is
// returned if the resulting profile contains invalid values.
func (profile Profile) Update(p2 Profile) (Profile, error) {
	if p2.Dtag == DoNotModify {
		p2.Dtag = profile.Dtag
	}

	if p2.Moniker == DoNotModify {
		p2.Moniker = profile.Moniker
	}

	if p2.Bio == DoNotModify {
		p2.Bio = profile.Bio
	}

	if p2.Pictures.Profile == DoNotModify {
		p2.Pictures.Profile = profile.Pictures.Profile
	}

	if p2.Pictures.Cover == DoNotModify {
		p2.Pictures.Cover = profile.Pictures.Cover
	}

	if p2.CreationDate.IsZero() {
		p2.CreationDate = profile.CreationDate
	}

	if p2.Creator == DoNotModify {
		p2.Creator = profile.Creator
	}

	newProfile := NewProfile(p2.Dtag, p2.Moniker, p2.Bio, p2.Pictures, p2.CreationDate, p2.Creator)
	err := newProfile.Validate()
	if err != nil {
		return Profile{}, err
	}

	return newProfile, nil
}

// Validate check the validity of the Profile
func (profile Profile) Validate() error {
	if strings.TrimSpace(profile.Dtag) == "" || profile.Dtag == DoNotModify {
		return fmt.Errorf("invalid profile DTag: %s", profile.Dtag)
	}

	if profile.Moniker == DoNotModify {
		return fmt.Errorf("invalid profile moniker: %s", profile.Moniker)
	}

	if profile.Bio == DoNotModify {
		return fmt.Errorf("invalid profile bio: %s", profile.Bio)
	}

	if profile.Pictures.Profile == DoNotModify {
		return fmt.Errorf("invalid profile picture: %s", profile.Pictures.Profile)
	}

	if profile.Pictures.Cover == DoNotModify {
		return fmt.Errorf("invalid profile cover: %s", profile.Pictures.Cover)
	}

	_, err := sdk.AccAddressFromBech32(profile.Creator)
	if err != nil {
		return fmt.Errorf("invalid creator address: %s", profile.Creator)
	}

	return profile.Pictures.Validate()
}

// ___________________________________________________________________________________________________________________

func NewDTagTransferRequest(dtagToTrade string, sender, receiver string) DTagTransferRequest {
	return DTagTransferRequest{
		DtagToTrade: dtagToTrade,
		Receiver:    receiver,
		Sender:      sender,
	}
}

// Validate checks the request validity
func (request DTagTransferRequest) Validate() error {
	_, err := sdk.AccAddressFromBech32(request.Sender)
	if err != nil {
		return fmt.Errorf("invalid sender address: %s", request.Sender)
	}

	_, err = sdk.AccAddressFromBech32(request.Receiver)
	if err != nil {
		return fmt.Errorf("invalid receiver address: %s", request.Receiver)
	}

	if request.Receiver == request.Sender {
		return fmt.Errorf("the sender and receiver must be different")
	}

	if strings.TrimSpace(request.DtagToTrade) == "" {
		return fmt.Errorf("invalid DTag to trade %s", request.DtagToTrade)
	}

	return nil
}
