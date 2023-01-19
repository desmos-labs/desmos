package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
	"github.com/spf13/cobra"
)

// DONTCOVER

// GetCmdQueryAllowances returns the command to query the fee allowances of a specific user
func GetCmdQueryAllowances() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "allowances [subspace-id] [grantee-type] [[user-address or group-id]]",
		Short:   "Query allowances for the given user or user group",
		Example: fmt.Sprintf(`%s query subspaces allowances 1 user desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7`, version.AppName),
		Args:    cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var subspaceID uint64
			subspaceID, err = types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			var grantee types.Grantee
			switch args[1] {
			case "user":
				if len(args) < 3 {
					grantee = types.NewUserGrantee("")
					break
				}

				_, err := sdk.AccAddressFromBech32(args[2])
				if err != nil {
					return err
				}

				grantee = types.NewUserGrantee(args[2])

			case "group":
				if len(args) < 3 {
					grantee = types.NewGroupGrantee(0)
					break
				}

				groupID, err := types.ParseGroupID(args[2])
				if err != nil {
					return err
				}

				grantee = types.NewGroupGrantee(groupID)

			default:
				return fmt.Errorf("unsupported grantee type: %s", args[2])
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Allowances(
				context.Background(),
				types.NewQueryAllowancesRequest(subspaceID, grantee, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "allowances")

	return cmd
}
