package models_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewMonodirectionalRelationship(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	expectedMr := models.MonodirectionalRelationship{
		Sender:   sender,
		Receiver: receiver,
	}

	actualMr := models.NewMonodirectionalRelationship(sender, receiver)
	require.Equal(t, expectedMr, actualMr)
}

func TestMonodirectionalRelationship_Creator(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	actualMr := models.NewMonodirectionalRelationship(sender, sender)

	require.Equal(t, sender, actualMr.Creator())
}

func TestMonodirectionalRelationship_Recipient(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	actualMr := models.NewMonodirectionalRelationship(sender, sender)

	require.Equal(t, sender, actualMr.Recipient())
}

func TestMonodirectionalRelationship_String(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	actualMr := models.NewMonodirectionalRelationship(sender, sender).String()
	expectedMr := "Mono directional Relationship:\n[Sender] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 -> [Receiver] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"

	require.Equal(t, expectedMr, actualMr)
}

func TestMonodirectionalRelationship_Equals(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name    string
		mr      models.MonodirectionalRelationship
		otherMr models.MonodirectionalRelationship
		expBool bool
	}{
		{
			name:    "Same relationships returns true",
			mr:      models.NewMonodirectionalRelationship(sender, receiver),
			otherMr: models.NewMonodirectionalRelationship(sender, receiver),
			expBool: true,
		},
		{
			name:    "Different relationships returns false",
			mr:      models.NewMonodirectionalRelationship(sender, receiver),
			otherMr: models.NewMonodirectionalRelationship(sender, sender),
			expBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.mr.Equals(test.otherMr))
		})
	}
}

func TestMonodirectionalRelationship_Validate(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name   string
		mr     models.MonodirectionalRelationship
		expErr error
	}{
		{
			name:   "empty sender returns error",
			mr:     models.NewMonodirectionalRelationship(sdk.AccAddress{}, receiver),
			expErr: fmt.Errorf("relationship sender cannot be empty"),
		},
		{
			name:   "empty receiver returns error",
			mr:     models.NewMonodirectionalRelationship(sender, sdk.AccAddress{}),
			expErr: fmt.Errorf("relationship receiver cannot be empty"),
		},
		{
			name:   "equals sender and receiver returns error",
			mr:     models.NewMonodirectionalRelationship(sender, sender),
			expErr: fmt.Errorf("you can't create a relationship with yourself"),
		},
		{
			name:   "valid relationship returns no error",
			mr:     models.NewMonodirectionalRelationship(sender, receiver),
			expErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.mr.Validate())
		})
	}
}

func TestMonoDirectionalRelationships_AppendIfMissing(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name    string
		mrs     models.MonoDirectionalRelationships
		mr      models.MonodirectionalRelationship
		expBool bool
		expMrs  models.MonoDirectionalRelationships
	}{
		{
			name:    "Already present mono relationship isn't inserted",
			mrs:     models.MonoDirectionalRelationships{models.NewMonodirectionalRelationship(sender, receiver)},
			mr:      models.NewMonodirectionalRelationship(sender, receiver),
			expBool: false,
			expMrs:  nil,
		},
		{
			name:    "New mono relationship is inserted correctly",
			mrs:     models.MonoDirectionalRelationships{models.NewMonodirectionalRelationship(sender, receiver)},
			mr:      models.NewMonodirectionalRelationship(sender, sender),
			expBool: true,
			expMrs: models.MonoDirectionalRelationships{
				models.NewMonodirectionalRelationship(sender, receiver),
				models.NewMonodirectionalRelationship(sender, sender),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mrs, appended := test.mrs.AppendIfMissing(test.mr)
			require.Equal(t, test.expMrs, mrs)
			require.Equal(t, test.expBool, appended)
		})
	}
}

func TestNewBiDirectionalRelationship(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	status := models.RelationshipStatus(0)

	expectedMr := models.BidirectionalRelationship{
		Sender:   sender,
		Receiver: receiver,
		Status:   status,
	}

	actualMr := models.NewBiDirectionalRelationship(sender, receiver, status)
	require.Equal(t, expectedMr, actualMr)
}

func TestBiDirectionalRelationship_Creator(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	status := models.RelationshipStatus(0)

	actualMr := models.NewBiDirectionalRelationship(sender, sender, status)
	require.Equal(t, sender, actualMr.Creator())
}

func TestBiDirectionalRelationship_Recipient(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	status := models.RelationshipStatus(0)

	actualMr := models.NewBiDirectionalRelationship(sender, sender, status)

	require.Equal(t, sender, actualMr.Recipient())
}

func TestBiDirectionalRelationship_String(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	tests := []struct {
		name      string
		br        models.BidirectionalRelationship
		expString string
	}{
		{
			name:      "String representation with status sent",
			br:        models.NewBiDirectionalRelationship(sender, sender, models.Sent),
			expString: "Bidirectional Relationship:\n[Sender] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 <-> [Receiver] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47\nStatus: Relationship not yet accepted or denied",
		},
		{
			name:      "String representation with status accepted",
			br:        models.NewBiDirectionalRelationship(sender, sender, models.Accepted),
			expString: "Bidirectional Relationship:\n[Sender] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 <-> [Receiver] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47\nStatus: Relationship accepted",
		},
		{
			name:      "String representation with status denied",
			br:        models.NewBiDirectionalRelationship(sender, sender, models.Denied),
			expString: "Bidirectional Relationship:\n[Sender] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 <-> [Receiver] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47\nStatus: Relationship denied",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expString, test.br.String())
		})
	}
}

func TestBiDirectionalRelationship_Equals(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	status := models.RelationshipStatus(0)

	tests := []struct {
		name    string
		mr      models.BidirectionalRelationship
		otherMr models.BidirectionalRelationship
		expBool bool
	}{
		{
			name:    "Same relationships returns true",
			mr:      models.NewBiDirectionalRelationship(sender, receiver, status),
			otherMr: models.NewBiDirectionalRelationship(sender, receiver, status),
			expBool: true,
		},
		{
			name:    "Different relationships returns false",
			mr:      models.NewBiDirectionalRelationship(sender, sender, status),
			otherMr: models.NewBiDirectionalRelationship(sender, receiver, status),
			expBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.mr.Equals(test.otherMr))
		})
	}
}

func TestBidirectionalRelationship_Validate(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	status := models.RelationshipStatus(0)

	tests := []struct {
		name   string
		mr     models.BidirectionalRelationship
		expErr error
	}{
		{
			name:   "empty sender returns error",
			mr:     models.NewBiDirectionalRelationship(sdk.AccAddress{}, receiver, status),
			expErr: fmt.Errorf("relationship sender cannot be empty"),
		},
		{
			name:   "empty receiver returns error",
			mr:     models.NewBiDirectionalRelationship(sender, sdk.AccAddress{}, status),
			expErr: fmt.Errorf("relationship receiver cannot be empty"),
		},
		{
			name:   "equals sender and receiver returns error",
			mr:     models.NewBiDirectionalRelationship(sender, sender, status),
			expErr: fmt.Errorf("you can't create a relationship with yourself"),
		},
		{
			name:   "valid relationship returns no error",
			mr:     models.NewBiDirectionalRelationship(sender, receiver, status),
			expErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.mr.Validate())
		})
	}
}

func TestBiDirectionalRelationship_AppendIfMissing(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	status := models.RelationshipStatus(0)

	tests := []struct {
		name    string
		mrs     models.BidirectionalRelationships
		mr      models.BidirectionalRelationship
		expBool bool
		expMrs  models.BidirectionalRelationships
	}{
		{
			name:    "Already present mono relationship isn't inserted",
			mrs:     models.BidirectionalRelationships{models.NewBiDirectionalRelationship(sender, receiver, status)},
			mr:      models.NewBiDirectionalRelationship(sender, receiver, status),
			expBool: false,
			expMrs:  nil,
		},
		{
			name:    "New mono relationship is inserted correctly",
			mrs:     models.BidirectionalRelationships{models.NewBiDirectionalRelationship(sender, receiver, status)},
			mr:      models.NewBiDirectionalRelationship(sender, sender, status),
			expBool: true,
			expMrs: models.BidirectionalRelationships{
				models.NewBiDirectionalRelationship(sender, receiver, status),
				models.NewBiDirectionalRelationship(sender, sender, status),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mrs, appended := test.mrs.AppendIfMissing(test.mr)
			require.Equal(t, test.expMrs, mrs)
			require.Equal(t, test.expBool, appended)
		})
	}
}
