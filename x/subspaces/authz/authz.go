package authz

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"

	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

// TODO: Revisit this once we have proper gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasCostPerIteration = uint64(10)

var (
	_ authz.Authorization = &GenericSubspaceAuthorization{}
)

// NewGenericSubspaceAuthorization creates a new GenericSubspaceAuthorization object.
func NewGenericSubspaceAuthorization(subspacesIDs []uint64, msgTypeURL string) *GenericSubspaceAuthorization {
	return &GenericSubspaceAuthorization{
		SubspacesIDs: subspacesIDs,
		Msg:          msgTypeURL,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a GenericSubspaceAuthorization) MsgTypeURL() string {
	return a.Msg
}

// Accept implements Authorization.Accept.
func (a GenericSubspaceAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	switch msg := msg.(type) {

	case types.SubspaceMsg:
		for _, subspaceID := range a.SubspacesIDs {
			ctx.GasMeter().ConsumeGas(gasCostPerIteration, "generic subspace authorization")
			if subspaceID == msg.GetSubspaceID() {
				return authz.AcceptResponse{Accept: true}, nil
			}
		}
		return authz.AcceptResponse{Accept: false}, nil

	default:
		return authz.AcceptResponse{}, errors.Wrap(sdkerrors.ErrInvalidRequest, "unsupported subspace msg type")
	}
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a GenericSubspaceAuthorization) ValidateBasic() error {
	if a.SubspacesIDs == nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "at least one subspace id is required")
	}

	for _, subspaceID := range a.SubspacesIDs {
		if subspaceID == 0 {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
		}
	}

	return nil
}
