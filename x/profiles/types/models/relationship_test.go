package models_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/desmos-labs/desmos/x/profiles/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRelationshipStatus_String(t *testing.T) {
	relStatus := types.Sent

	require.Equal(t, "0", relStatus.String())
}

func TestRelationshipID_Valid(t *testing.T) {
	tests := []struct {
		name    string
		id      types.RelationshipID
		expBool bool
	}{
		{
			name:    "Valid id returns true",
			id:      types.RelationshipID("1234"),
			expBool: true,
		},
		{
			name:    "Invalid id returns false",
			id:      types.RelationshipID(""),
			expBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.id.Valid())
		})
	}
}

func TestRelationshipID_String(t *testing.T) {
	id := types.RelationshipID("12345")
	idString := id.String()
	require.Equal(t, "12345", idString)
}

func TestRelationshipID_Equals(t *testing.T) {
	tests := []struct {
		name    string
		id      types.RelationshipID
		otherID types.RelationshipID
		expBool bool
	}{
		{
			name:    "Equals IDs returns true",
			id:      types.RelationshipID("1234"),
			otherID: types.RelationshipID("1234"),
			expBool: true,
		},
		{
			name:    "Non equals IDs returns false",
			id:      types.RelationshipID("123"),
			otherID: types.RelationshipID("1234"),
			expBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.id.Equals(test.otherID))
		})
	}
}

func TestNewMonodirectionalRelationship(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	expectedMr := models.MonodirectionalRelationship{
		ID:       "13cf1724ba76d87b1ea55259aa46c44976d9296430ae779a20c9fc51bc28e2ef",
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

func TestMonodirectionalRelationship_RelationshipID(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	actualMr := models.NewMonodirectionalRelationship(sender, sender)

	require.Equal(t, types.RelationshipID("5f575d15aa5cec9c9df0f6fd19a356767b37fef7da8b5f99e2ff891be3e4f174"), actualMr.RelationshipID())
}

func TestMonodirectionalRelationship_String(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	actualMr := models.NewMonodirectionalRelationship(sender, sender).String()
	expectedMr := "Mono directional Relationship:\n[RelationshipID] 5f575d15aa5cec9c9df0f6fd19a356767b37fef7da8b5f99e2ff891be3e4f174 [Sender] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 -> [Receiver] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"

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
		ID:       "63b66fa545035229b8991bb6553268bfa59d29f4f5e5c345ac470f94613babbf",
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

	actualBr := models.NewBiDirectionalRelationship(sender, sender, status)

	require.Equal(t, sender, actualBr.Recipient())
}

func TestBidirectionalRelationship_RelationshipID(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)
	status := models.RelationshipStatus(0)

	actualBr := models.NewBiDirectionalRelationship(sender, sender, status)

	require.Equal(t, types.RelationshipID("77f4b74cb2b4fd89d8a0cd0f044bdbb5e9e300cced95d7c8b8d22c4ec13ac0f8"), actualBr.RelationshipID())
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
			expString: "Bidirectional Relationship:\n[RelationshipID] 77f4b74cb2b4fd89d8a0cd0f044bdbb5e9e300cced95d7c8b8d22c4ec13ac0f8 [Sender] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 <-> [Receiver] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47\nStatus: Relationship not yet accepted or denied",
		},
		{
			name:      "String representation with status accepted",
			br:        models.NewBiDirectionalRelationship(sender, sender, models.Accepted),
			expString: "Bidirectional Relationship:\n[RelationshipID] 77f4b74cb2b4fd89d8a0cd0f044bdbb5e9e300cced95d7c8b8d22c4ec13ac0f8 [Sender] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 <-> [Receiver] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47\nStatus: Relationship accepted",
		},
		{
			name:      "String representation with status denied",
			br:        models.NewBiDirectionalRelationship(sender, sender, models.Denied),
			expString: "Bidirectional Relationship:\n[RelationshipID] 77f4b74cb2b4fd89d8a0cd0f044bdbb5e9e300cced95d7c8b8d22c4ec13ac0f8 [Sender] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 <-> [Receiver] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47\nStatus: Relationship denied",
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
